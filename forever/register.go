package forever

import (
	"github.com/sirupsen/logrus"
	"iFei/config/baseCon"
	"iFei/config/mysqlCon"
)

func BaseConRegister() {
	IFconBase := baseCon.LoadBaseConfig()
	logrus.SetLevel(IFconBase.LogLevel)
}

func MysqlRegister() {
	IFMysqlCon := mysqlCon.LoadMysqlConfig()
	Connect(IFMysqlCon)
	//db, err := gorm.Open("mysql", IFMysqlCon.MysqlUri)
	//if err != nil{
	//	logrus.Errorf("connect mysql error...")
	//	return
	//}

}
