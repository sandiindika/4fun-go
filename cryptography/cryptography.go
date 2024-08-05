package cryptography

func Encrypt(text string, s int8) string {
	result := ""

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(s))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+rune(s))%26 + 'a')
		} else {
			result += string(char)
		}
	}

	return result
}

func Decrypt(encryptedText string, s int8) string {
	return Encrypt(encryptedText, -s)
}
