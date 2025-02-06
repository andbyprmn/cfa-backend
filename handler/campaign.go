package handler

import (
	"cfa-backend/auth"
	"cfa-backend/campaign"
	"cfa-backend/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	authService     auth.Service
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
