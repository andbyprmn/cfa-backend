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
// @Summary      Get detail of transaction
// @Description  Get detail of transaction by campaign id
// @Tags         transactions
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

	transactionDetail, err := h.transactionService.GetTransactionByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get campaign transaction!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsDetailFormatter := transaction.FormatCampaignTransactions(transactionDetail)
	response := helper.APIResponse("List of campaign transactions!", http.StatusOK, "success", transactionsDetailFormatter)
	c.JSON(http.StatusOK, response)
}
