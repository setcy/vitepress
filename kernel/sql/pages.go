package sql

import "time"

type Page struct {
	ID          int `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string
	Hash        string
	Title       string
	Description string
	IsPublished bool
	Content     string
	Render      string
	Toc         string
}

func QueryContentByPath(path string) (content string, err error) {
	var page Page

	err = db.Debug().Select("render").Where("path = ?", path).First(&page).Error
	return page.Render, err
}
