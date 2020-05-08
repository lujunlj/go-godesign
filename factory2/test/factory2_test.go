package test

import (
	"godesgin/factory2"
	"godesgin/factory2/impl"
	"testing"
)

func TestFactory(t *testing.T) {
	var meiDfactory factory2.IFacotry = impl.MeiDfactory{}
	meiDfactory.ProduceTV().ShowTV()
	meiDfactory.ProduceRefrigerator().Ice()

	var tclFactory factory2.IFacotry = impl.TclFactory{}
	tclFactory.ProduceRefrigerator().Ice()
	tclFactory.ProduceTV().ShowTV()
}
