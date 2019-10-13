package model

type User struct {
	Model

	Name              string `gorm:"size:32" json:"name"`
	IfeiKey           string `gorm:"size:32" json:"ifeikey"`
	TotalArticleCount int    `json:"totalArticleCount"`

	Articles []Article `gorm:"foreignkey:UserID"`
}
