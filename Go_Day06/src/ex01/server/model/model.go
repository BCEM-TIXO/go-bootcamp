package model

import (
	"gorm.io/gorm"
	"html/template"
)

type ArticleRecord struct {
	gorm.Model
	Title string `gorm:"not null"`
}

type Paging struct {
	Page     int
	PrevPage int
	NextPage int
	HasPrev  bool
	HasNext  bool
}

type HomePage struct {
	Title          string
	Paging         Paging
	ArticlesTitles []ArticleTitle
}

type ArticleTitle struct {
	Title string
}
type Article struct {
	Title   ArticleTitle
	Content template.HTML
}
