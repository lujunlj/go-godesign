package main

import "godesgin/factory"

func main(){
	sdf := factory.SenderFactory{}
	sdf.Produce("sms").Send()
	sdf.Produce("email").Send()
	sdf.Produce("express").Send()

	msdf := factory.MutiSenderFactory{}
	msdf.ProduceEmailSender().Send()
	msdf.ProduceExpressSender().Send()
	msdf.ProduceSmsSender().Send()

	factory.ProduceSmsSender().Send()
	factory.ProduceExpressSender().Send()
	factory.ProduceEmailSender().Send()

}

