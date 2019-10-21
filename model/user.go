package model

type User struct {
	Model

	Name              string `sql:"index" gorm:"size:32;unique" json:"name"`
	IfeiKey           string `gorm:"size:32" json:"ifeikey"`
	Password          string `gorm:"not null"`
	TotalArticleCount int    `json:"total_article_count"`

	Articles []Article `gorm:"foreignkey:UserID"`
}
