package main

// +build ignore
// https://github.com/teh-cmc/go-internals/blob/master/chapter2_interfaces/README.md#function-and-method-calls

//go:noinline
func Add(a, b int32) int32 { return a + b }

type Adder struct{ id int32 }

//go:noinline
func (adder *Adder) AddPtr(a, b int32) int32 { return a + b }

//go:noinline
func (adder Adder) AddVal(a, b int32) int32 { return a + b }

func main() {
	Add(10, 32) // direct call of top-level function

	adder := Adder{id: 6754}
	adder.AddPtr(10, 32) // direct call of method with pointer receiver
	adder.AddVal(10, 32) // direct call of method with value receiver

	(&adder).AddVal(10, 32) // implicit dereferencing
}
