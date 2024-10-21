package test

import (
	"fmt"
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"github.com/rizky201008/subscription-tracker-api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

var SettingRepository = repositories.NewSettingRepository()

func insertFakeSettingData() {
	data := domain.Setting{
		UserID:   "ABCD",
		Currency: "USD",
	}

	err := Db.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

func truncateSettingTable() {
	query := Db.Exec("TRUNCATE TABLE settings")
	if query.Error != nil {
		panic(query.Error)
	}
}

func TestGetSetting(t *testing.T) {
	data, err := SettingRepository.GetSetting("ABCD", Db)
	assert.Nil(t, err)
	assert.Equal(t, data.UserID, "ABCD")
	fmt.Println(data)
}

func TestSaveSettingSuccess(t *testing.T) {
	data := domain.Setting{
		UserID:   "BBBB",
		Currency: "USD",
	}

	setting, err := SettingRepository.SaveSetting(data, Db)
	assert.Nil(t, err)
	assert.Equal(t, setting.UserID, "BBBB")
}

func TestSaveSettingError(t *testing.T) {
	data := domain.Setting{
		UserID:   "ABCD", // duplicate ID
		Currency: "USD",
	}

	setting, err := SettingRepository.SaveSetting(data, Db)
	assert.NotNil(t, err)
	assert.Equal(t, setting.UserID, "ABCD")
	fmt.Println(err)
}

func TestUpdateSettingSuccess(t *testing.T) {
	var setting *domain.Setting
	result, err := SettingRepository.GetSetting("ABCD", Db)
	setting = &result
	assert.Nil(t, err)
	setting.Currency = "IDR"
	updated, err := SettingRepository.UpdateSetting(*setting, Db)
	assert.Equal(t, updated.Currency, "IDR")
}

func TestUpdateSettingError(t *testing.T) {
	var setting *domain.Setting
	result, err := SettingRepository.GetSetting("ABCD", Db)
	setting = &result
	assert.Nil(t, err)
	setting.Currency = "IDRRRRR" // exceed maximum length
	_, err = SettingRepository.UpdateSetting(*setting, Db)
	assert.NotNil(t, err)
}
