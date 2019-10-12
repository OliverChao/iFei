package forever

func MysqlUnRegister() {
	DisconnectDB()

}

func RedisUnRegister() {
	DisConnectRedis()

}
