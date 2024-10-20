package test

import "testing"

func TestMain(m *testing.M) {
	truncateSettingTable()
	truncateSubscriptionsTable()
	insertFakeSettingData()
	insertFakeSubscriptionData()
	m.Run()
}
