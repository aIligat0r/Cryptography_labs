package task1

func CryptorDecryptor(text string, key int, direction int) string {
	shift, offset := rune(key), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1:
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}
		runes[index] = char
	}

	return string(runes)
}
