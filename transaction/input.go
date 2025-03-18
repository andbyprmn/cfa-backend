package transaction

import "cfa-backend/user"

type GetCampaignIDTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
