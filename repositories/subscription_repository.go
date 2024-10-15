package repositories

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetAll(db *gorm.DB) (domain.Subscription, error)
	GetByID(id string, db *gorm.DB) (domain.Subscription, error)
	Save(data domain.Subscription, db *gorm.DB) (domain.Subscription, error)
	Update(data domain.Subscription, db *gorm.DB) (domain.Subscription, error)
}

type SubscriptionRepositoryImpl struct{}

func NewSubscriptionRepository() SubscriptionRepository {
	return &SubscriptionRepositoryImpl{}
}

func (SubscriptionRepositoryImpl) GetAll(db *gorm.DB) (domain.Subscription, error) {
	var subscription domain.Subscription
	err := db.Find(&subscription).Error
	return subscription, err
}

func (SubscriptionRepositoryImpl) GetByID(id string, db *gorm.DB) (domain.Subscription, error) {
	var subscription domain.Subscription
	err := db.First(id).First(&subscription).Error
	return subscription, err
}

func (SubscriptionRepositoryImpl) Save(data domain.Subscription, db *gorm.DB) (domain.Subscription, error) {
	err := db.Save(&data).Error
	return data, err
}

func (SubscriptionRepositoryImpl) Update(data domain.Subscription, db *gorm.DB) (domain.Subscription, error) {
	err := db.Where("id = ?", data.ID).Save(&data).Error
	return data, err
}
