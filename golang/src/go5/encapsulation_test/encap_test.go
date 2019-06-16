package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("%s-%s-%d", e.Id, e.Name, e.Age)
}

//func (e *Employee) String() string{
//	return fmt.Sprintf("%s-%s-%d",e.Id,e.Name,e.Age)
//}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e2 := Employee{Name: "Mike", Age: 30}
	e3 := new(Employee)
	e3.Id = "2"
	e3.Age = 22
	e3.Name = "Rose"
	t.Log(e)
	t.Log(e2)
	t.Log(e3)
	t.Logf("%T", e)
	t.Logf("%T", e2)
	t.Logf("%T", e3)

}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
