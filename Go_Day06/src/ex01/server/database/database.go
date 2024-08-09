package database

import (
	"ex01/server/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDB(user, password, dbname, host string, port int) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.ArticleRecord{})

	return &Database{db: db}
}

func (d *Database) CreateArticleRecord(record model.ArticleRecord) error {
	result := d.db.Create(&record)
	return result.Error
}

func (d *Database) GetPage(page, pageSize int) []model.ArticleTitle {
	var articles_r []model.ArticleRecord
	d.db.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&model.ArticleRecord{}).Select("title").Find(&articles_r)
	articles := make([]model.ArticleTitle, len(articles_r))
	for i, a := range articles_r {
		articles[i].Title = a.Title
	}
	return articles
}

func (d *Database) PageCount(pageSize int) int {
	var count int64
	d.db.Find(&model.ArticleRecord{}).Count(&count)
	return (int(count) + (pageSize - 1)) / pageSize
}
