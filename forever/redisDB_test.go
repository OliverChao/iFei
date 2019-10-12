package forever

import (
	"iFei/config/redisCon"
	"testing"
)

func TestConnectionRedis(t *testing.T) {
	redisConfig := redisCon.LoadRedisConfig()

	ConnectionRedis(redisConfig)
	client.FlushAll()
	client.HSet("user:1:info", "age", 20)
	get := client.HGet("user:1:info", "age")
	s, _ := get.Result()
	if s != "20" {
		t.Fail()
	}
	DisConnectionRedis()
}
