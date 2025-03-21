package handler

import (
	"cfa-backend/campaign"
	"cfa-backend/helper"
	"cfa-backend/user"
	"fmt"
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

// GetCampaigns godoc
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

// GetCampaign godoc
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

// CreateCampaign godoc
// @Summary      Create campaign
// @Description  Create new campaign
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        body  body  campaign.CreateCampaignInput  true  "Campaign create data"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaigns [post]
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

	createdCampaignFormatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Campaign has been successfuly created!", http.StatusOK, "success", createdCampaignFormatter)
	c.JSON(http.StatusOK, response)
}

// UpdateCampaign godoc
// @Summary      Update Campaign
// @Description  Update campaign by campaign id
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        id query int false "Campaign ID"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaigns/:id [put]
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputURI campaign.GetCampaignDetailInput
	var input campaign.CreateCampaignInput

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	err := c.ShouldBindUri(&inputURI)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to update campaign!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedCampaign, err := h.campaignService.UpdateCampaign(inputURI, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to update campaign!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCampaignFormatter := campaign.FormatCampaign(updatedCampaign)
	response := helper.APIResponse("Campaign has been successfuly updated!", http.StatusOK, "success", updatedCampaignFormatter)
	c.JSON(http.StatusOK, response)
}

// UploadImageCampaign godoc
// @Summary      Upload campaign image
// @Description  Upload a new campaign image for the user campaign
// @Tags         Campaigns
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "Campaign image file"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /campaign-images [post]
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload campaign image!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image!", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	input.User = currentUser
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image!", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image!", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign image successfuly uploaded!", http.StatusOK, "error", data)

	c.JSON(http.StatusOK, response)
}
