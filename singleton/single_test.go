package singleton

import (
	"fmt"
	"testing"
)

func Test_Slacker(t *testing.T){
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1 == s2 )
}


func Test_Slovently(t *testing.T){
	s1 := GetInstance2()
	s2 := GetInstance2()
	fmt.Println(s1 == s2)
	s3 := NewConfig()
	s4 := NewConfig()
	fmt.Println(s3 == s4 )
	fmt.Println(s1 == s3 )
}