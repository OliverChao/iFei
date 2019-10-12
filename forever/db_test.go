package forever

import (
	"iFei/config/mysqlCon"
	"iFei/model"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {

	MysqlUri := "root:toor@(127.0.0.1:3306)/ifei?charset=utf8mb4&parseTime=True&loc=Local"
	IFMysqlCon := &mysqlCon.MysqlConfig{
		MysqlUri:        MysqlUri,
		MaxIdleConns:    15,
		MaxOpenConns:    25,
		ConnMaxLifetime: time.Second * 4,
		Models: []interface{}{
			&model.Model{},
			&model.User{},
		},
	}
	Connect(IFMysqlCon)

	DisconnectDB()
}
