package service

import (
	"CRUD/pkg/model"
	"CRUD/pkg/repository"
)

type Authorisation interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, int64, error)
}

type News interface {
	FindNewsById(id int) (model.News, error)
	GetAllNews() ([]model.News, error)
	PostNews(news model.News) (int, error)
	RemoveNews(id int) error
	UpdateNews(id int, news string) (model.News, error)
}

type Subscribers interface {
	SubscribeToNews(userId, newsId int) (model.News, error)
	UnsubscribeFromNews(userId any, newsId int) error
	GetAllSubscribersByNewsID(id int) ([]string, error)
}

type Service struct {
	Authorisation
	News
	Subscribers
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		News:          NewNewsService(repos.News),
		Subscribers:   NewSubscriberService(repos.Subscribers, repos.News),
	}
}
