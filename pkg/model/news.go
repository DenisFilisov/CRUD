package model

type News struct {
	Id          int64  `json:"id" db:"id" example:"1"`
	Description string `json:"description" example:"custom description"`
}
