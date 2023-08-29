package model

type Followers struct {
	Id     int
	UserId int `json:"userId" db:"user_id"`
	NewsId int `json:"newsId" binding:"required" db:"news_id"`
}
