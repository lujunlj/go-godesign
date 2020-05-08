package factory

import "godesgin/factory/impl"

type MutiSenderFactory struct {}

func (msf MutiSenderFactory) ProduceSmsSender() ISender{
	return impl.SmsSender{}
}

func (msf MutiSenderFactory) ProduceExpressSender() ISender {
	return impl.ExpressSender{}
}

func (msf MutiSenderFactory) ProduceEmailSender() ISender {
	return impl.EmailSender{}
}


