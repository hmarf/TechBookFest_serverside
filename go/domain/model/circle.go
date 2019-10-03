package model

import (
	"database/sql"
)

type Circle struct {
	CircleURL   string         `json:"CircleURL"`
	Circle      string         `json:"Circle"`
	CircleImage string         `json:"CircleImage"`
	Arr         string         `json:"arr"`
	Genere      string         `json:"Genere"`
	Keyword     string         `json:"Keyword"`
	Title       string         `json:"Title"`
	Content     sql.NullString `json:"Content"`
}
