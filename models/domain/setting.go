package domain

type Setting struct {
	UserID   string `gorm:"column:user_id;type:text"`
	Currency string `gorm:"column:currency;type:varchar(5);default:USD;"`
}
