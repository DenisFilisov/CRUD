package repository

import (
	"CRUD/pkg/model"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SubscribersPostgres struct {
	db *sqlx.DB
}

func (s *SubscribersPostgres) GetAllSubscribersByNewsID(id int) ([]string, error) {
	query := fmt.Sprintf("SELECT u.name FROM %s as f INNER JOIN %s as u ON u.id=f.user_id WHERE f.news_id=$1", followersTable, usersTable)
	var followers []string
	row, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var f string
		if err := row.Scan(&f); err != nil {
			return nil, err
		}
		followers = append(followers, f)
	}

	return followers, nil
}

func (s *SubscribersPostgres) UnsubscribeFromNews(userId any, newsId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 and news_id=$2", followersTable)
	_, err := s.db.Query(query, userId, newsId)
	return err
}

func (s *SubscribersPostgres) IsSubscribed(userId any, newsId int) error {
	query := fmt.Sprintf("SELECT b.id FROM %s as a INNER JOIN %s as b ON a.id=b.news_id WHERE user_id=$1 and news_id=$2", newsTable, followersTable)
	row := s.db.QueryRow(query, userId, newsId)
	var id int

	if err := row.Scan(&id); err != nil {
		return errors.New("User not subscribed to this news")
	}

	return nil
}

func NewSubscribersPostgres(db *sqlx.DB) *SubscribersPostgres {
	return &SubscribersPostgres{db: db}
}

func (s *SubscribersPostgres) SubscribeToNews(userId, newsId int) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 and news_id =$2", followersTable)
	var news model.Followers
	err := s.db.Get(&news, query, userId, newsId)
	if err == nil {
		return errors.New("This user allready subscribed")
	}

	insert := fmt.Sprintf("INSERT INTO %s (user_id, news_id) values ($1, $2)", followersTable)
	err1 := s.db.QueryRow(insert, userId, newsId)
	return err1.Err()
}
