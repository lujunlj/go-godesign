package impl

import "fmt"


//快递发送者
type ExpressSender struct {}

func (exs ExpressSender) Send(){
	fmt.Println("sended express !")
}
