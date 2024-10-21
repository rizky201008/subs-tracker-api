package web

type SettingRequest struct {
	UserID   string `json:"userId"`
	Currency string `json:"currency"`
}
