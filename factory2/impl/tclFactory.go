package impl

import "godesgin/factory2"

type TclFactory struct {}

func (tclFactory TclFactory) ProduceTV() factory2.Itv {
	return factory2.TCLTV{
		Name:  "t30",
		Color: "red",
	}
}

func (tclFactory TclFactory) ProduceRefrigerator() factory2.Irefrigerator{
	return factory2.TCLfrigerator{}
}

