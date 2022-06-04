package main

import (
	"fmt"
	task1 "lab_1/task_1"
	task2 "lab_1/task_2"
	task3 "lab_1/task_3"
	task4 "lab_1/task_4"
	task5 "lab_1/task_5"
	task6 "lab_1/task_6"
	task7 "lab_1/task_7"
)

func main() {

	var task int
	var source_text string = "ITISTESTSTRING"

	println("Task 1\nTask 2\nTask 3\nTask 4\nTask 5\nTask 6\nTask 7\n")

	fmt.Print("Select task:\n> ")
	fmt.Scan(&task)
	println("Text:", source_text)

	switch task {
	case 1:
		crypt := task1.CryptorDecryptor(source_text, 5, -1)
		fmt.Println("Crypted text:", crypt)
		decrypt := task1.CryptorDecryptor(crypt, 5, 1)
		fmt.Println("Decrypted text:", decrypt)
	case 2:
		crypt := task2.Cryptor(source_text, 5)
		fmt.Println("Crypted text:", crypt)
		decrypt := task2.Decryptor(crypt, 5)
		fmt.Println("Decrypted text:", decrypt)
	case 3:
		crypt := task3.Cryptor(source_text, 5)
		fmt.Println("Crypted text:", crypt)
		decrypt := task3.Decryptor(crypt, 5)
		println("Decrypted text:", decrypt)
	case 4:
		table := task4.Table_generator("KEY", 5, 5, 0)
		task4.Show_table(table)
		crypt := task4.Decryptor_Encryptor(source_text, table, "KEY", true)
		fmt.Println("Crypted text:", crypt)
		decrypt := task4.Decryptor_Encryptor(crypt, table, "KEY", false)
		fmt.Println("Decrypted text:", decrypt)
	case 5:
		crypt := task5.Cryptor(source_text, "KEY")
		fmt.Println("Crypted text:", crypt)
		decrypt := task5.Decryptor(crypt, "KEY")
		fmt.Println("Decrypted text:", decrypt)
	case 6:
		task6.CaesarCryptBrutter("Pbatenghyngvbaf! Vg'f n Pnrfne pvcure!")
	case 7:
		crypt := task7.ROT13(source_text)
		fmt.Println("Crypted text:", crypt)
		decrypt := task7.ROT13(crypt)
		fmt.Println("Decrypted text:", decrypt)
	}
}
