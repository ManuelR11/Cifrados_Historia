package rot13

// Encrypt aplica el cifrado ROT13 al texto dado.
func Encrypt(text string) string {
	var encrypted string

	for _, char := range text {
		switch {
		case char >= 'a' && char <= 'z':
			// Rota la letra 13 posiciones en minúsculas
			encrypted += string('a' + (char-'a'+13)%26)
		case char >= 'A' && char <= 'Z':
			// Rota la letra 13 posiciones en mayúsculas
			encrypted += string('A' + (char-'A'+13)%26)
		default:
			// No modifica caracteres no alfabéticos
			encrypted += string(char)
		}
	}

	return encrypted
}

// Decrypt para ROT13 es el mismo que Encrypt, ya que el cifrado es simétrico.
func Decrypt(text string) string {
	return Encrypt(text)
}
