package repository

import (
	"CRUD/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user model.User) (int, error)
	FindUserByUserNameAndPswd(username, password string) (model.User, error)
}

type News interface {
	GetAllNews() ([]model.News, error)
	FindNewsById(id int) (model.News, error)
	PostNews(news model.News) (int, error)
	RemoveNews(id int) error
	UpdateNews(id int, news string) (model.News, error)
}

type Subscribers interface {
	SubscribeToNews(userId, newsId int) error
	IsSubscribed(userId any, newsId int) error
	UnsubscribeFromNews(userId any, newsId int) error
	GetAllSubscribersByNewsID(id int) ([]string, error)
}

type Repository struct {
	Authorisation
	News
	Subscribers
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		News:          NewNewsPostgres(db),
		Subscribers:   NewSubscribersPostgres(db),
	}
}
