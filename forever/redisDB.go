package forever

import (
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"iFei/config/redisCon"
)

var client *redis.Client

func ConnectionRedis(con *redisCon.RedisConfig) {
	//client := redis.NewClient{con.RedisOption}
	client = redis.NewClient(con.RedisOption)
	logrus.Info("[redis]connect successfully...")
}
func DisConnectionRedis() {
	if e := client.Close(); e != nil {
		logrus.Error("[redis]close failed...")
	} else {
		logrus.Info("[redis]close successfully")
	}
}
