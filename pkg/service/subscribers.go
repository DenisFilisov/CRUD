package service

import (
	"CRUD/pkg/model"
	"CRUD/pkg/repository"
	"errors"
	"fmt"
)

type subscriberService struct {
	repoS repository.Subscribers
	repoN repository.News
	auth  repository.Authorisation
}

func (f *subscriberService) GetAllSubscribersByNewsID(id int) ([]string, error) {
	return f.repoS.GetAllSubscribersByNewsID(id)
}

func NewSubscriberService(repoS repository.Subscribers, repoN repository.News) *subscriberService {
	return &subscriberService{repoS: repoS, repoN: repoN}
}

func (f *subscriberService) UnsubscribeFromNews(userId any, newsId int) error {
	err := f.repoS.IsSubscribed(userId, newsId)
	if err != nil {
		return err
	}
	return f.repoS.UnsubscribeFromNews(userId, newsId)
}

func (f *subscriberService) SubscribeToNews(userId, newsId int) (model.News, error) {
	news, err := f.repoN.FindNewsById(newsId)

	if err != nil {
		return model.News{}, errors.New(fmt.Sprintf("Can't find news with id=%d", newsId))
	}

	err = f.repoS.SubscribeToNews(userId, newsId)
	if err != nil {
		return news, err
	}

	return news, err
}
