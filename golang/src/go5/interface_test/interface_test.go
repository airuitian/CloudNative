package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
}
