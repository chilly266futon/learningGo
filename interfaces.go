package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch va := a.(type) {
	case int:
		return va * b
	case string:
		return strings.Repeat(va, b)
	case fmt.Stringer:
		return strings.Repeat(va.String(), b)
	default:
		return nil
	}
}
