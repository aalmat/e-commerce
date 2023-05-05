package models

type Search struct {
	Keyword string `json:"keyword" validate:"required"`
}
