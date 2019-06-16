package polymorphism_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.println(\"Hello World\")"
}

type GoProgrammer struct {
}

func (j *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestPolymorphism(t *testing.T) {
	var p Programmer
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
