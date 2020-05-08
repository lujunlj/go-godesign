package singleton

type config2 struct {}

var Cfg2 *config2
var Cfg3 = new(config2)

func init(){
	Cfg2 = new(config2)
}

func GetInstance2() *config2 {
	return Cfg2
}
func NewConfig() *config2{
	return Cfg3
}
