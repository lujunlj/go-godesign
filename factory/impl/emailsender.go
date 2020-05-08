package impl

import "fmt"

//邮件发送者
type EmailSender struct {}

func (es EmailSender) Send(){
	fmt.Println("sended email !")
}
