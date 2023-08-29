package service

import (
	"CRUD/pkg/model"
	"CRUD/pkg/repository"
	"errors"
	"fmt"
)

type NewsService struct {
	repo repository.News
}

func (n *NewsService) UpdateNews(id int, news string) (model.News, error) {
	_, err := n.repo.FindNewsById(id)
	if err != nil {
		return model.News{}, err
	}
	return n.repo.UpdateNews(id, news)
}

func (n *NewsService) RemoveNews(id int) error {
	_, err := n.repo.FindNewsById(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't find news with id=%d", id))
	}
	return n.repo.RemoveNews(id)
}

func (n *NewsService) PostNews(news model.News) (int, error) {
	return n.repo.PostNews(news)
}

func (n *NewsService) FindNewsById(id int) (model.News, error) {
	return n.repo.FindNewsById(id)
}

func (n *NewsService) GetAllNews() ([]model.News, error) {
	return n.repo.GetAllNews()
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}
