package main

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

//func area(f figures) (func(float64) float64, bool) {
//	switch f {
//	case square:
//		return func(x float64) float64 { return x * x }, true
//	case circle:
//		return func(x float64) float64 { return x * x * 3.14 }, true
//	case triangle:
//		return func(x float64) float64 { return x * x * 0.5 }, true
//	default:
//		return nil, false
//	}
//}
//
//func main() {
//	myFigure := square
//	ar, ok := area(myFigure)
//	if !ok {
//		fmt.Println("Ошибка")
//		return
//	}
//	x := 2.5
//	myArea := ar(x)
//	fmt.Println(myArea)
//}
