package mysqlCon

import "time"
import "iFei/model"

const (
	MysqlUri = "root:toor@(127.0.0.1:3306)/ifei?charset=utf8mb4&parseTime=True&loc=Local"
)

type MysqlConfig struct {
	MysqlUri string
	//MaxConnNUm int
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	Models          []interface{}
}

func LoadMysqlConfig() (IFMysqlCon *MysqlConfig) {

	IFMysqlCon = &MysqlConfig{
		MysqlUri:        MysqlUri,
		MaxIdleConns:    15,
		MaxOpenConns:    25,
		ConnMaxLifetime: time.Second * 4,
		Models: []interface{}{
			&model.Model{},
			&model.User{},
		},
	}
	return
}
