package transaction

import (
	"cfa-backend/campaign"
	"cfa-backend/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// CampaignImages []campaign.CampaignImage
}
