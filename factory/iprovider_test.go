package factory

import "testing"

func TestIProvider(t *testing.T){
	smsft := SMSFactory{}
	smsft.Produce().Send()
	emailft := EmailFactory{}
	emailft.Produce().Send()
	expressft := ExpressFactory{}
	expressft.Produce().Send()
}

func TestAdd(t *testing.T) {
	t.Log(t.Name(),"Begin...")
	a := 2
	b := 3
	t.Logf("a + b is %d",a+b)
	t.Log(t.Name(),"End...")
}