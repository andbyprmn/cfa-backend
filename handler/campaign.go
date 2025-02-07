package handler

import (
	"cfa-backend/campaign"
	"cfa-backend/helper"
	"cfa-backend/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

// UploadAvatar godoc
// @Summary      Get list of campaign
// @Description  Get list of campaign by user id
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        user_id query int false "User ID"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaigns [get]
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Error to get campaigns!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignsFormatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("List of campaigns!", http.StatusOK, "success", campaignsFormatter)
	c.JSON(http.StatusOK, response)
}

// UploadAvatar godoc
// @Summary      Get detail of campaign
// @Description  Get detail of campaign by campaign id
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        id query int false "Campaign ID"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaigns/:id [get]
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail campaign!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignByID(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail campaign!", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignsDetailFormatter := campaign.FormatCampaignDetail(campaignDetail)
	response := helper.APIResponse("Detail of campaign!", http.StatusOK, "success", campaignsDetailFormatter)
	c.JSON(http.StatusOK, response)
}

// UploadAvatar godoc
// @Summary      Create campaign
// @Description  Create new campaign
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        body  body  campaign.CreateCampaignInput  true  "Campaign create data"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaigns/:id [get]
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to create campaign!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createCampaignFormatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Campaign has been successfuly created!", http.StatusOK, "success", createCampaignFormatter)
	c.JSON(http.StatusOK, response)
}
