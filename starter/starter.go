package starter

import "fmt"

func SayHello(str string) string {
	return fmt.Sprintf("%s%s%s", "Hello ", str, ". Welcome!")
}

func OddOrEven(i int) interface{} {
	evenVal, oddVal := 0, 1
	return If(i%2 == 0).getValue(
		evenVal,
		oddVal,
	)
}

type If bool

func (c If) getValue(a interface{}, b interface{}) interface{} {
	if c {
		return a
	} else {
		return b
	}
}
