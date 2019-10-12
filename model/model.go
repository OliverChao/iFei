package model

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	TestData string `gorm:"column:test_data;type:varchar(35)"`
	Name     string `gorm:"column:name;index:model_name"`
}

type User struct {
	gorm.Model

	Name string `gorm:"size:32" json:"name"`
	//Nickname          string `gorm:"size:32" json:"nickname"`
	//AvatarURL         string `gorm:"size:255" json:"avatarURL"`
	IfeiKey string `gorm:"size:32" json:"ifeikey"`
	Locale  string `gorm:"size:32" json:"locale"`
	//TotalArticleCount int    `json:"totalArticleCount"`
	//GithubId          string `gorm:"255" json:"githubId"`
}
