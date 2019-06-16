package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSth(p interface{}) {
	//if i,ok := p.(int);ok{
	//	fmt.Println(i)
	//	return
	//}
	//
	//if s,ok := p.(string);ok {
	//	fmt.Println(s)
	//	return
	//}
	//fmt.Println("unkown")

	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unkown")

	}
}

func TestEmptyInterface(t *testing.T) {
	DoSth(10)
	DoSth("10")
}
