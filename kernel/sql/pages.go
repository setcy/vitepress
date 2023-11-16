package sql

import (
	"encoding/json"
	"gorm.io/datatypes"
	"time"
)

type Page struct {
	ID          int `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string
	Hash        string
	Title       string
	Description string
	Content     string
	Render      string
	Toc         datatypes.JSON
}

type Toc struct {
	Title    string `json:"title"`
	Anchor   string `json:"anchor"`
	Children []*Toc `json:"children,omitempty"`
}

func QueryContentByPath(path string) (content string, toc []*Toc, err error) {
	var page Page

	err = db.Select("render", "toc").Where("path = ?", path).First(&page).Error
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(page.Toc, &toc)
	if err != nil {
		return "", nil, err
	}

	return page.Render, toc, nil
}

func QueryPageMeta() (page []*Page, err error) {

	err = db.Debug().Select([]string{"path", "title"}).Find(&page).Error
	if err != nil {
		return nil, err
	}

	return page, err
}
