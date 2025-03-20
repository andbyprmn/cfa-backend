package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	transactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int                              `json:"id"`
	Amount    int                              `json:"amount"`
	Status    string                           `json:"status"`
	CreatedAt time.Time                        `json:"created_at"`
	Campaign  UserCampaignTransactionFormatter `json:"campaign"`
}

type UserCampaignTransactionFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	userCampaignTransactionFormatter := UserCampaignTransactionFormatter{}
	userCampaignTransactionFormatter.Name = transaction.Campaign.Name

	if len(transaction.Campaign.CampaignImages) > 0 {
		userCampaignTransactionFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = userCampaignTransactionFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	transactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)

	}

	return transactionsFormatter
}
