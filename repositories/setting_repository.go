package repositories

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"gorm.io/gorm"
)

type SettingRepository interface {
	GetSetting(id string, db *gorm.DB) (domain.Setting, error)
	SaveSetting(setting domain.Setting, db *gorm.DB) (domain.Setting, error)
	UpdateSetting(setting domain.Setting, db *gorm.DB) (domain.Setting, error)
}

type SettingRepositoryImpl struct{}

func NewSettingRepository() SettingRepository {
	return SettingRepositoryImpl{}
}

func (SettingRepositoryImpl) UpdateSetting(setting domain.Setting, db *gorm.DB) (domain.Setting, error) {
	err := db.First(&setting, setting.ID).Updates(&setting).Error
	return setting, err
}

func (SettingRepositoryImpl) GetSetting(id string, db *gorm.DB) (domain.Setting, error) {
	var setting domain.Setting
	err := db.First(&setting, "user_id = ?", id).Error
	return setting, err
}

func (SettingRepositoryImpl) SaveSetting(setting domain.Setting, db *gorm.DB) (domain.Setting, error) {
	err := db.Create(&setting).Error
	return setting, err
}
