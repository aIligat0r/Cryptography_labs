package task2

import (
	"fmt"
	"math/rand"
	"strings"
)

const ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func crypt_alphabet_generator(key int64) []string {
	var crypted_alpabet [26]string
	for symbol := range ALPHABET {
		crypted_alpabet[symbol] = string(ALPHABET[symbol])
	}
	rand.Seed(key)
	rand.Shuffle(len(crypted_alpabet), func(i, j int) { crypted_alpabet[i], crypted_alpabet[j] = crypted_alpabet[j], crypted_alpabet[i] })
	return crypted_alpabet[:]
}

func Cryptor(source_text string, key int64) string {

	var crypted_text string
	var crypted_alpabet []string = crypt_alphabet_generator(key)
	fmt.Printf("Crypted alphabet: %s (Key: %d)\n", strings.Join(crypted_alpabet[:], ""), key)

	for symbol_i := range source_text {
		upper := false
		if source_text[symbol_i] < 90 {
			upper = true
		}

		if strings.Contains(strings.ToLower(ALPHABET), strings.ToLower(string(source_text[symbol_i]))) == false {
			crypted_text += string(source_text[symbol_i])
		} else {
			crypted_symbol_i := strings.Index(strings.ToLower(ALPHABET), strings.ToLower(string(source_text[symbol_i])))
			if upper == true {
				crypted_text += strings.ToUpper(crypted_alpabet[crypted_symbol_i])
			} else {
				crypted_text += strings.ToLower(crypted_alpabet[crypted_symbol_i])
			}

		}
	}
	return crypted_text
}

func Decryptor(crypted_text string, key int64) string {

	var decrypted_text string
	var crypted_alpabet []string = crypt_alphabet_generator(key)

	for symbol_i := range crypted_text {
		upper := false
		if ALPHABET[symbol_i] < 90 {
			upper = true
		}

		if strings.Contains(strings.ToLower(strings.Join(crypted_alpabet[:], "")), strings.ToLower(string(ALPHABET[symbol_i]))) == false {
			decrypted_text += string(crypted_text[symbol_i])
		} else {
			decrypted_symbol_i := strings.Index(strings.ToLower(strings.Join(crypted_alpabet[:], "")), strings.ToLower(string(ALPHABET[symbol_i])))
			if upper == true {
				decrypted_text += strings.ToUpper(string(crypted_alpabet[decrypted_symbol_i]))
			} else {
				decrypted_text += strings.ToLower(string(crypted_alpabet[decrypted_symbol_i]))
			}
		}

	}
	return decrypted_text
}
