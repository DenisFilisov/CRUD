package service

import (
	"CRUD/pkg/model"
	"CRUD/pkg/repository"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Authorisation interface {
	CreateUser(user model.User) (int, error)
	GenerateTokens(oldToken string, user model.User) (string, string, error)
	ParseToken(token string) (int, int64, error)
	RefreshToken(refreshToken string) (string, string, error)
	FindUserByUsernameAndPswd(username, password string) (model.User, error)
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

type Kafka interface {
	SentMessage(message *logrus.Entry)
}

type Service struct {
	Authorisation
	News
	Subscribers
	Kafka
}

func NewService(repos *repository.Repository, producer *kafka.Producer, producerChan chan kafka.Event) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		News:          NewNewsService(repos.News),
		Subscribers:   NewSubscriberService(repos.Subscribers, repos.News),
		Kafka:         NewKafkaSender(producerChan, producer),
	}
}
