package model

import "github.com/jinzhu/gorm"

type Model struct{
	gorm.Model
	TestData string `gorm:"column:test_data;type:varchar(35)"`
	Name string `gorm:"column:name;index:model_name"`

}