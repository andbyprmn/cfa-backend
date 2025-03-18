package transaction

import "gorm.io/gorm"

type Repository interface {
	GetTransactionByCampaignID(ID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransactionByCampaignID(ID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", ID).Order("id DESC").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
