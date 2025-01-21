package atbash

// Encrypt aplica el cifrado Atbash al texto dado.
func Encrypt(text string) string {
	var encrypted string

	for _, char := range text {
		switch {
		case char >= 'a' && char <= 'z':
			// Calcula la letra opuesta en el alfabeto
			encrypted += string('z' - (char - 'a'))
		case char >= 'A' && char <= 'Z':
			// Calcula la letra opuesta en el alfabeto para mayúsculas
			encrypted += string('Z' - (char - 'A'))
		default:
			// No modifica caracteres no alfabéticos
			encrypted += string(char)
		}
	}

	return encrypted
}

// Decrypt para Atbash es el mismo que Encrypt, ya que el cifrado es simétrico.
func Decrypt(text string) string {
	return Encrypt(text)
}
