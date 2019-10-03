package infra

import (
	"database/sql"
	"go/domain/model"
	"go/domain/repository"
)

type circleInfraStruct struct {
	db *sql.DB
}

func NewCircleDB(db *sql.DB) repository.CircleRepository {
	return &circleInfraStruct{db: db}
}

func (c *circleInfraStruct) GetCircleData(keyword string) (circles []model.Circle, err error) {
	sql := "SELECT * FROM circle where content like '%" + keyword + "%'"
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		circle := model.Circle{}
		if err := rows.Scan(&circle.CircleURL, &circle.Circle, &circle.CircleImage,
			&circle.Arr, &circle.Genere, &circle.Keyword, &circle.Title, &circle.Content); err != nil {
			return nil, err
		}
		circles = append(circles, circle)
	}
	return
}

func (c *circleInfraStruct) GetAllCircleData() (circles []model.Circle, err error) {
	sql := "SELECT * FROM circle"
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		circle := model.Circle{}
		if err := rows.Scan(&circle.CircleURL, &circle.Circle, &circle.CircleImage,
			&circle.Arr, &circle.Genere, &circle.Keyword, &circle.Title, &circle.Content); err != nil {
			return nil, err
		}
		circles = append(circles, circle)
	}
	return
}
