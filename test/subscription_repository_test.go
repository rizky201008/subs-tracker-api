package test

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"github.com/rizky201008/subscription-tracker-api/repositories"
	"time"
)

var SubscriptionRepository = repositories.NewSubscriptionRepository()

func insertFakeSubscriptionData() {
	data := domain.Subscription{
		Amount:       100000,
		Cycle:        domain.Monthly,
		ColorHex:     "000000",
		DueDate:      time.Date(2024, time.October, 8, 0, 0, 0, 0, time.UTC),
		PlatformName: "Spotify",
		Reminder:     false,
		UserId:       "ABC",
	}

	err := Db.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

func truncateSubscriptionsTable() {
	query := Db.Exec("TRUNCATE TABLE settings")
	if query.Error != nil {
		panic(query.Error)
	}
}
