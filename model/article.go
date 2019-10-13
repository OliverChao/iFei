package model

import "time"

type Article struct {
	Model
	PushedAt     time.Time `json:"pushedAt"`
	Title        string    `gorm:"size:255" json:"title"`
	Content      string    `gorm:"type:text" json:"content"`
	Tags         string    `gorm:"type:text" json:"tags"`
	Commentable  bool      `json:"commentable"`
	ViewCount    int       `json:"viewCount"`
	CommentCount int       `json:"commentCount"`

	Path   string `sql:"index" gorm:"size:255" json:"path"`
	Stared int    `sql:"index" json:"star"`
	Topped bool   `json:"topped"`

	UserID uint64 `json:"user_id"`

	Comments []Comment `gorm:"foreignkey:ArticleID" json:"comments"`
	//Tag      []Tag `gorm:"foreignkey:"`
}

type Comment struct {
	Model

	ArticleID uint64    `json:"articleID"`
	Content   string    `gorm:"type:text" json:"content"`
	ReplyID   uint64    `json:"reply_id"`
	PushedAt  time.Time `json:"pushedAt"`

	UserID uint64 `json:"UserID"`
	User   User   `gorm:"foreignkey:UserID"`
}

type Tag struct {
	Model

	Title        string `gorm:"size:255" json:"title"`
	ArticleCount int    `json:"article_count"`
	//Articles     []Article `gorm:"foreignkey:ArticleID"`
}
