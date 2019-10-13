package forever

import "testing"

func TestBaseConRegister(t *testing.T) {
	//tests := []struct {
	//	name string
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})
	//}
}

func TestMysqlRegister(t *testing.T) {

	MysqlRegister()

	MysqlUnRegister()

}
func TestMysqlFunction(t *testing.T) {

	MysqlRegister()
	QueryDemo()
	MysqlUnRegister()
}

func TestRedisRegister(t *testing.T) {
	//tests := []struct {
	//	name string
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})
	//}
}

func TestMysqlDropAll(t *testing.T) {
	MysqlDropAll()
	MysqlUnRegister()
}
