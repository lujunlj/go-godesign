package factory

import "godesgin/factory/impl"

type StaticSenderFactory struct{}

//var sdf StaticSenderFactory
//
//func init(){
//	sdf = StaticSenderFactory{}
//}

func ProduceEmailSender() ISender{
	return impl.EmailSender{}
}

func ProduceSmsSender() ISender{
	return impl.SmsSender{}
}
func ProduceExpressSender() ISender{
	return impl.ExpressSender{}
}

