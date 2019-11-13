// +build ignore

package main

func foo(a int, b *int, c struct{ e, f int }, d *struct{ e, f int }) (g struct{ h, i float32 }, j int) {
	g.h = float32(a + *b)
	g.i = float32(c.e * d.f)

	return
}

func bar(a int, b *int) (struct{ h, i float32 }, int) {
	return struct{ h, i float32 }{1, 2}, 3
}

type A struct {
	a, b int
}

type B = *A

func (a A) baz() {

}
func (a *A) baz2() {

}

func (b B) foo() {}

func test() {
	a := A{}

	a.baz()
	a.baz2()

	(&a).baz()

	b := &A{}

	b.foo()
}
