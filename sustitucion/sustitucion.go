package sustitucion

import "errors"

// Alfabeto fijo para el cifrado de sustitución (incluye la ñ)
const substitutionAlphabet = "qwertyuiopasdfghjklzñxcvbnm"

// Encrypt aplica el cifrado de sustitución usando el alfabeto predefinido.
func Encrypt(text string) (string, error) {
	if len(substitutionAlphabet) != 27 {
		return "", errors.New("el alfabeto de sustitución debe contener exactamente 27 caracteres (incluyendo la 'ñ')")
	}

	var encrypted string

	for _, char := range text {
		switch {
		case isLowercaseLetter(char):
			index := positionInAlphabet(char)
			encrypted += string(substitutionAlphabet[index])
		case isUppercaseLetter(char):
			index := positionInAlphabet(char + 'a' - 'A')                // Convierte a minúscula para obtener la posición
			encrypted += string(substitutionAlphabet[index] - 'a' + 'A') // Convierte el resultado a mayúscula
		default:
			encrypted += string(char) // No cifra caracteres no alfabéticos
		}
	}

	return encrypted, nil
}

// Decrypt aplica el descifrado de sustitución usando el alfabeto predefinido.
func Decrypt(text string) (string, error) {
	if len(substitutionAlphabet) != 27 {
		return "", errors.New("el alfabeto de sustitución debe contener exactamente 27 caracteres (incluyendo la 'ñ')")
	}

	var decrypted string
	reverseMap := make(map[rune]rune)

	// Crear un mapa inverso del alfabeto
	for i, char := range substitutionAlphabet {
		reverseMap[char] = rune('a' + i)
	}

	for _, char := range text {
		switch {
		case isLowercaseLetter(char):
			decrypted += string(reverseMap[char])
		case isUppercaseLetter(char):
			lowerChar := char + 'a' - 'A'                          // Convierte a minúscula
			decrypted += string(reverseMap[lowerChar] - 'a' + 'A') // Convierte el resultado a mayúscula
		default:
			decrypted += string(char) // No descifra caracteres no alfabéticos
		}
	}

	return decrypted, nil
}

// Función auxiliar para determinar si una letra es minúscula (incluyendo la ñ)
func isLowercaseLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || char == 'ñ'
}

// Función auxiliar para determinar si una letra es mayúscula (incluyendo la Ñ)
func isUppercaseLetter(char rune) bool {
	return (char >= 'A' && char <= 'Z') || char == 'Ñ'
}

// Función para obtener la posición de una letra en el alfabeto (incluye la ñ)
func positionInAlphabet(char rune) int {
	if char >= 'a' && char <= 'z' {
		if char >= 'ñ' {
			return int(char-'a') + 1 // La ñ ocupa la posición 14
		}
		return int(char - 'a')
	}
	if char == 'ñ' {
		return 14 // La ñ ocupa la posición 14
	}
	return -1
}
