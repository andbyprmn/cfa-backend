package transaction

import (
	"cfa-backend/campaign"
	"errors"
)

type Service interface {
	GetTransactionByID(input GetCampaignIDTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository: repository, campaignRepository: campaignRepository}
}

func (s *service) GetTransactionByID(input GetCampaignIDTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("You do not have authorization to get list of campaign transactions!")
	}

	transactions, err := s.repository.GetTransactionByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, err
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetTransactionByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, err
}
