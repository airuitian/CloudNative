package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 0
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	tp := reflect.TypeOf(f)
	switch tp.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Println("unkown", tp)
	}
}

type Customer struct {
	ID   string
	Name string `json:"name"`
	Age  int
}

func (e *Customer) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two"}
	b := map[int]string{1: "one", 3: "two"}
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 3, 4}
	t.Log(reflect.DeepEqual(s1, s2))
	t.Log(reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log(reflect.DeepEqual(c1, c2))

}

func TestInvokeByName(t *testing.T) {
	e := &Customer{"1", "Mike", 40}
	t.Logf("(%[1]v,%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	t.Logf(reflect.ValueOf(*e).FieldByName("Name").String())

	if nameField, ok := reflect.TypeOf(e).Elem().FieldByName("Name"); !ok {
		t.Error("failed to get field")
	} else {
		t.Log(nameField.Name)
		t.Log(nameField.Tag.Get("json"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log(e)

}

func TestFillBySettings(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	for k, v := range settings {
		reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v))
	}
	return nil
}
