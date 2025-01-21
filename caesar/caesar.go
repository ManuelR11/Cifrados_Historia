package caesar

// Encrypt aplica el cifrado César a un texto dado con un desplazamiento específico.
func Encrypt(text string, shift int) string {
	shift = shift % 26
	var encrypted string

	for _, char := range text {
		switch {
		case char >= 'a' && char <= 'z':
			encrypted += string('a' + (char-'a'+rune(shift))%26)
		case char >= 'A' && char <= 'Z':
			encrypted += string('A' + (char-'A'+rune(shift))%26)
		default:
			encrypted += string(char)
		}
	}

	return encrypted
}

// Decrypt aplica el descifrado César a un texto dado con un desplazamiento específico.
func Decrypt(text string, shift int) string {
	return Encrypt(text, -shift) // El descifrado es un cifrado con desplazamiento negativo
}
