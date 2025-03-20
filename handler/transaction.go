package handler

import (
	"cfa-backend/helper"
	"cfa-backend/transaction"
	"cfa-backend/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

// GetTransaction godoc
// @Summary      Get list of campaign transactions
// @Description  Get list of transactions by campaign id
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        id query int false "Campaign ID"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaign/:id/transactions [get]
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignIDTransactionInput

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get campaign transactions!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get campaign transactions!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormatter := transaction.FormatCampaignTransactions(transactions)
	response := helper.APIResponse("List of campaign transactions!", http.StatusOK, "success", transactionsFormatter)
	c.JSON(http.StatusOK, response)
}

// GetTransaction godoc
// @Summary      Get list of user transactions
// @Description  Get list of transactions by user id
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        id query int false "User ID"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /user/:id/transactions [get]
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	transactions, err := h.transactionService.GetTransactionByUserID(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get user transactions!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormatter := transaction.FormatUserTransactions(transactions)
	response := helper.APIResponse("List of user transactions!", http.StatusOK, "success", transactionsFormatter)
	c.JSON(http.StatusOK, response)

}
