package models

type Session struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	SKey   string `json:"skey"`
}
