package impl

import "fmt"

type SmsSender struct {}

func(sms SmsSender) Send(){
	fmt.Println("sended sms !")
}
