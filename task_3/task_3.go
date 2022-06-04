package task3

import (
	"math/rand"
)

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ .!?"

func crypted_alphabet_generator(key int64) map[string][3]int {
	rand.Seed(key)
	var CRYPTED_ALPHABET = make(map[string][3]int)
	var random_numbers []int = rand.Perm(len(ALPHABET) * 3)
	var i int = 0
	for symbol := range ALPHABET {
		CRYPTED_ALPHABET[string(ALPHABET[symbol])] = [3]int{random_numbers[i], random_numbers[i+1], random_numbers[i+2]}
		i += 3
	}

	return CRYPTED_ALPHABET
}

func Cryptor(text string, key int64) []int {
	var CRYPTED_ALPHABET = crypted_alphabet_generator(key)
	var crypted_text []int = make([]int, len(text))
	for symbol := range text {
		crypted_text[symbol] = CRYPTED_ALPHABET[string(text[symbol])][rand.Intn(3)]
	}

	return crypted_text
}

func Decryptor(crypted_text []int, key int64) string {
	var CRYPTED_ALPHABET = crypted_alphabet_generator(key)
	var decrypted_text string = ""

	for i := range crypted_text {
		for symbol := range CRYPTED_ALPHABET {
			for e := range CRYPTED_ALPHABET[symbol] {
				if crypted_text[i] == CRYPTED_ALPHABET[symbol][e] {
					decrypted_text += symbol
				}
			}
		}
	}

	return decrypted_text
}
