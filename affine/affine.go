package affine

import (
	"errors"
	"math"
)

// Encrypt aplica el cifrado Afín al texto dado.
func Encrypt(text string, a, b int) (string, error) {
	m := 27 // Tamaño del alfabeto español (incluye la 'ñ')
	if gcd(a, m) != 1 {
		return "", errors.New("a debe ser coprimo con el tamaño del alfabeto (27). Valores válidos para 'a': {1, 2, 4, 5, 7, 8, 10, 11, 13, 14, 16, 17, 19, 20, 22, 23, 25}")
	}

	var encrypted string
	for _, char := range text {
		switch {
		case char >= 'a' && char <= 'z':
			x := positionInAlphabet(char)
			cipher := (a*x + b) % m
			encrypted += letterFromPosition(cipher)
		case char == 'ñ':
			x := 14 // Posición de la 'ñ'
			cipher := (a*x + b) % m
			encrypted += letterFromPosition(cipher)
		default:
			encrypted += string(char) // No cifra caracteres no alfabéticos
		}
	}

	return encrypted, nil
}

// Decrypt aplica el descifrado Afín al texto dado.
func Decrypt(text string, a, b int) (string, error) {
	m := 27 // Tamaño del alfabeto español (incluye la 'ñ')
	invA, err := ModInverse(a, m)
	if err != nil {
		return "", err
	}

	var decrypted string
	for _, char := range text {
		switch {
		case char >= 'a' && char <= 'z':
			y := positionInAlphabet(char)
			plain := (invA * (y - b + m)) % m
			decrypted += letterFromPosition(plain)
		case char == 'ñ':
			y := 14 // Posición de la 'ñ'
			plain := (invA * (y - b + m)) % m
			decrypted += letterFromPosition(plain)
		default:
			decrypted += string(char) // No descifra caracteres no alfabéticos
		}
	}

	return decrypted, nil
}

// positionInAlphabet devuelve la posición de una letra en el alfabeto español.
func positionInAlphabet(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	if char == 'ñ' {
		return 14
	}
	return -1
}

// letterFromPosition devuelve la letra correspondiente a una posición en el alfabeto español.
func letterFromPosition(pos int) string {
	if pos < 14 {
		return string('a' + rune(pos))
	}
	if pos == 14 {
		return "ñ"
	}
	return string('a' + rune(pos-1))
}

// gcd calcula el Máximo Común Divisor (MCD) de dos números.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return int(math.Abs(float64(a)))
}

// ModInverse calcula el inverso modular de a mod m.
func ModInverse(a, m int) (int, error) {
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x, nil
		}
	}
	return 0, errors.New("no existe inverso modular para el valor dado")
}
