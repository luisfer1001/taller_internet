package tools

import "errors"

func Suma(a, b int) int {
	var c int

	c = a + b

	return c
}

func Resta(a int, b int) int {
	var c int

	c = a - b

	return c
}

func SumaResta(a, b int) (int, int) {
	var rs int
	var rr int

	rs = Suma(a, b)
	rr = Resta(a, b)

	return rs, rr
}

func Division(a, b int) (int, error) {
	var c int

	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}

	c = a / b

	return c, nil
}
