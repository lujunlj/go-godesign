package factory

import "godesgin/factory/impl"

type IProvider interface {
	Produce() ISender
}

type SMSFactory struct {}

type EmailFactory struct {}

type ExpressFactory struct {}

func (smsf SMSFactory) Produce() ISender{
	return impl.SmsSender{}
}

func (email EmailFactory) Produce() ISender {
	return impl.EmailSender{}
}

func (express ExpressFactory) Produce() ISender {
	return impl.ExpressSender{}
}

