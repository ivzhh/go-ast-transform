// +build ignore

package main

func foo(a int, b *int, c struct{ e, f int }, d *struct{ e, f int }) (g struct{ h, i float32 }) {
	g.h = float32(a + *b)
	g.i = float32(c.e * c.f)

	return
}
