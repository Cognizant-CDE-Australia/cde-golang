package model

import "time"

type Book struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Created time.Time `json:"time"`
}
