package forever

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"iFei/config/mysqlCon"
)

var db *gorm.DB
var err error

func Connect(con *mysqlCon.MysqlConfig) {
	db, err = gorm.Open("mysql", con.MysqlUri)
	if err != nil {
		logrus.Fatalf("[mysql] connect mysql error" + err.Error())
		//return
	} else {
		logrus.Info("[mysql] connect successfully")
	}

	if err = db.AutoMigrate(con.Models...).Error; err != nil {
		logrus.Fatal("[mysql] auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(con.MaxIdleConns)
	db.DB().SetMaxOpenConns(con.MaxOpenConns)
	db.DB().SetConnMaxLifetime(con.ConnMaxLifetime)
}

func DisconnectDB() {
	if err := db.Close(); err != nil {
		logrus.Errorf("[mysql] Disconnect failed: " + err.Error())
	} else {
		logrus.Info("[mysql] Disconnect successful")
	}
}
