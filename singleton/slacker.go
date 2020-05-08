package singleton

import "sync"

type config struct {}

var Cfg *config
var oSingle sync.Once


//懒汉式
func GetInstance() *config {
	oSingle.Do(
		func(){
			if Cfg == nil {
				Cfg = new(config)
			}
		})
	return Cfg
}


