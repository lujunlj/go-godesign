package factory

import (
	"godesgin/factory/impl"
	"strings"
)

type SenderFactory struct{}

func (sdf SenderFactory) Produce(tp string) ISender {
	if tp == "" {
		return nil
	} else if strings.EqualFold(tp, "sms") {
		return impl.SmsSender{}
	} else if strings.EqualFold(tp, "email") {
		return impl.EmailSender{}
	} else if strings.EqualFold(tp, "express") {
		return impl.ExpressSender{}
	} else {
		return nil
	}
}
