package task7

import (
	task1 "lab_1/task_1"
)

func ROT13(text string) string {
	var result string
	result = task1.CryptorDecryptor(text, 13, +1)

	return result
}
