package factory2

import "fmt"

type Irefrigerator interface {
	Ice()
}

type TCLfrigerator struct {}

type MeiDfrigerator struct {}

func (tclF TCLfrigerator) Ice(){
	fmt.Println("TCL 冰箱")
}

func (meiDfrigerator MeiDfrigerator) Ice(){
	fmt.Println("MEID 冰箱")
}