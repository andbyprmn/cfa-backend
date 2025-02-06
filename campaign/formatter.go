package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                            `json:"id"`
	UserID           int                            `json:"user_id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	ImageURL         string                         `json:"image_url"`
	GoalAmount       int                            `json:"goal_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	Description      string                         `json:"description"`
	Slug             string                         `json:"slug"`
	Perks            []string                       `json:"perks"`
	User             CampaignDetailUserFormatter    `json:"user"`
	Images           []CampaignDetailImageFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignDetailImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Description:      campaign.Description,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	//Set Objek data
	formatter.Perks = perks

	user := campaign.User
	campaignDetailUserFormatter := CampaignDetailUserFormatter{}
	campaignDetailUserFormatter.Name = user.Name
	campaignDetailUserFormatter.ImageURL = user.AvatarFileName

	//Set Object user
	formatter.User = campaignDetailUserFormatter

	campaignDetailImagesFormatter := []CampaignDetailImageFormatter{}
	for _, campaignImage := range campaign.CampaignImages {
		imageFormatter := CampaignDetailImageFormatter{}
		isPrimary := false

		imageFormatter.ImageURL = campaignImage.FileName
		if campaignImage.IsPrimary == 1 {
			isPrimary = true
		}

		imageFormatter.IsPrimary = isPrimary
		campaignDetailImagesFormatter = append(campaignDetailImagesFormatter, imageFormatter)
	}

	//Set Object images
	formatter.Images = campaignDetailImagesFormatter

	return formatter
}
