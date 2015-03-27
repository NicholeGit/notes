package singleton

// 结构体的名字,而且采用小写,私有化
type singleton struct {
}

// 对象,而且采用小写,私有化
var instance *singleton

// 只能通过函数取得
// 保证唯一
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
