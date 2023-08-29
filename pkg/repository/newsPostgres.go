package repository

import (
	"CRUD/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func (n *NewsPostgres) UpdateNews(id int, news string) (model.News, error) {
	query := fmt.Sprintf("UPDATE %s set description=$1 where id=$2", newsTable)
	n.db.QueryRow(query, news, id)

	return n.FindNewsById(id)
}

func (n *NewsPostgres) RemoveNews(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", newsTable)
	_, err := n.db.Query(query, id)
	return err
}

func (n *NewsPostgres) PostNews(news model.News) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (description) values ($1) RETURNING id", newsTable)
	row := n.db.QueryRow(query, news.Description)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (n *NewsPostgres) GetAllNews() ([]model.News, error) {
	query := fmt.Sprintf("SELECT * FROM %s", newsTable)
	rows, err := n.db.Query(query)
	if err != nil {
		return nil, err
	}

	var news []model.News
	for rows.Next() {
		var n model.News
		err := rows.Scan(&n.Id, &n.Description)
		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}

	return news, nil
}

func (n *NewsPostgres) FindNewsById(newsId int) (model.News, error) {
	var news model.News
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", newsTable)
	if err := n.db.Get(&news, query, newsId); err != nil {
		return model.News{}, err
	}
	return news, nil
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}
