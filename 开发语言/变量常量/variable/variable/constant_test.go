package variable

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstant2(t *testing.T) {
	//a := 1 //0001
	//a := 2 //0010
	//a := 3 //0011
	a := 7 //0111
	t.Log(a & Readable == Readable, a & Writable == Writable, a & Executable == Executable)
}