package ejercicios

import "strconv"

func ConvertiraInt(ejercicio01 string) (int, string) {
	var numero int
	numero, _ = strconv.Atoi(ejercicio01)

	if numero >= 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "No es mayor a 100"
	}

}
