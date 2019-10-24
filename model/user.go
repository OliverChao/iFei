package model

type User struct {
	Model

	Name              string `sql:"index" gorm:"size:32;unique" json:"name"`
	IfeiKey           string `gorm:"size:32" json:"ifeikey"`
	Password          string `gorm:"not null"`
	TotalArticleCount int    `json:"total_article_count"`

	Articles []Article `gorm:"foreignkey:UserID"`
}

func (u *User) RetData() map[string]interface{} {
	m := make(map[string]interface{})
	m["name"] = u.Name
	m["upload"] = u.TotalArticleCount
	m["_id"] = u.Model.ID
	return m
}

//func (u *User)GetArticles()[]Article{
//	// todo : add error control
//	db := forever.GetGlobalGormDB()
//	db.Model(&u).Related(&u.Articles)
//	return u.Articles
//}
