package web

import "time"

type SubscriptionRequest struct {
	PlatformName string    `json:"platform_name"`
	Amount       float64   `json:"amount"`
	DueDate      time.Time `json:"due_date"`
	Reminder     bool      `json:"reminder"`
	Cycle        string    `json:"cycle"`
	ColorHex     string    `json:"color_hex"`
}
