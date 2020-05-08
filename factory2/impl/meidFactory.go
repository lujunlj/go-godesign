package impl

import "godesgin/factory2"

type MeiDfactory struct {}

func (meidf MeiDfactory) ProduceTV() factory2.Itv{
	return factory2.MeiDTV{}
}

func (meiDf MeiDfactory) ProduceRefrigerator() factory2.Irefrigerator{
	return factory2.MeiDfrigerator{}
}