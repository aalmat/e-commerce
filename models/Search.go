package models

type Search struct {
	Keyword string `json:"keyword" binding: required`
}
