package factory2

import "fmt"

type Itv interface {
	ShowTV()
}

type TCLTV struct {
	Name  string
	Color string
}

func (tclTv TCLTV) ShowTV() {
	fmt.Printf("This is tclTV ! Name is : %s , Color is : %s", tclTv.Name, tclTv.Color)
}

type MeiDTV struct {
	Name  string
	Color string
}

func (meiDTV MeiDTV) ShowTV() {
	fmt.Printf("MeiDTV is %v", MeiDTV{
		Name:  "111",
		Color: "222",
	})
}
