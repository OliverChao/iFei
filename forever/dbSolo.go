package forever

import (
	"errors"
	"github.com/sirupsen/logrus"
	"iFei/model"
)

func AddArticle(article *model.Article) {
	//todo : 从 userid 中查出user, 更改user信息
	db.Create(article)
}

func VerifyUser(name, passwd string) (*model.User, error) {
	user := model.User{
		Name:     name,
		Password: passwd,
	}
	//if err := db.Model(&user).Find(&user).Error; err!=nil{
	//	logrus.Error("VerifyUser failed")
	//	return nil, err
	//}else{
	//	logrus.Info("VerifyUser successfully")
	//}
	//return &user,nil
	db.Model(&user).Find(&user)
	if user.Password != passwd {
		logrus.Error("VerifyUser failed")
		return nil, errors.New("VerifyUser failed")
	} else {
		logrus.Info("VerifyUser successfully")
	}
	return &user, nil

}
