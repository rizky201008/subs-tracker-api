package domain

import (
	"gorm.io/gorm"
)

type BillingCycle string

const (
	Weekly  BillingCycle = "W"
	Monthly BillingCycle = "M"
	Yearly  BillingCycle = "Y"
)

func StringToBilling(s string) BillingCycle {
	switch s {
	case "W":
		return Weekly
	case "M":
		return Monthly
	case "Y":
		return Yearly
	}
	return Monthly
}

func (b BillingCycle) Value() string {
	switch b {
	case Weekly:
		return "W"
	case Monthly:
		return "M"
	case Yearly:
		return "Y"
	}
	return ""
}

func (b BillingCycle) IsValid() bool {
	switch b {
	case Weekly, Monthly, Yearly:
		return true
	}
	return false
}

type Subscription struct {
	gorm.Model
	PlatformName string       `gorm:"column:platform_name;type:varchar(255)"`
	UserId       string       `gorm:"type:text;column:user_id;not null"`
	Amount       float64      `gorm:"column:amount;not null"`
	DueDate      string       `gorm:"column:due_date;type:date;not null"`
	Reminder     bool         `gorm:"column:reminder;type:boolean;default:false"`
	Cycle        BillingCycle `gorm:"column:cycle;type:enum('W','M','Y')"`
	ColorHex     string       `gorm:"column:color_hex;type:varchar(6)"`
}
