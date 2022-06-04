package task6

import task1 "lab_1/task_1"

func CaesarCryptBrutter(crypted_text string) {
	for i := 1; i < 26; i++ {
		brute_result := task1.CryptorDecryptor(crypted_text, i, +1)
		println(i, brute_result)
	}
}
