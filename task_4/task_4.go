package task4

import (
	"strings"
)

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func matrix_generator(x int, y int, element_value int) (matrix [][]int) {
	matrix = make([][]int, x)
	for i := range matrix {
		matrix[i] = make([]int, y)
		for e := 0; e < y; e++ {
			matrix[i][e] = element_value
		}
	}
	return matrix
}

func contains(slice []int, element int) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func Show_table(table [][]int) {
	a := 0
	println()
	for i := range table {
		for e := range table[i] {
			print("|", string(table[i][e]))
			a++
			if a == 5 {
				print("|")
				a = 0
				println()
			}
		}
	}
	println()
}

func Table_generator(key string, matrix_X int, matrix_Y int, matrix_Elements int) [][]int {
	matrix := matrix_generator(matrix_X, matrix_Y, matrix_Elements)
	alphabet := []rune(strings.ToUpper(key) + strings.ToUpper(ALPHABET))
	// fmt.Println(matrix)
	var symbols_log []int = make([]int, 0)

	var i, e int = 0, 0
	for s := 0; s < len(alphabet); s++ {
		// fmt.Println("s = ", s, "\t", "e = ", e, "letter = ", alphabet[s])
		if alphabet[s] == 'J' {
			alphabet[s] = 'I'
		}
		if contains(symbols_log, int(alphabet[s])) == false && alphabet[s] >= 65 && alphabet[s] <= 90 {
			if e >= matrix_Y {
				e = 0
				i++
			}
			matrix[i][e] = int(alphabet[s])
			e++
		}
		symbols_log = append(symbols_log, int(alphabet[s]))
	}

	return matrix
}

func get_position(symbol int, table [][]int) []int {
	var position = make([]int, 0)

	for column := range table {
		for row := range table[column] {
			if symbol == table[column][row] {
				position = append(position, column)
				position = append(position, row)
			}
		}
	}
	return position
}

func first_condition(symbol_posion []int, table [][]int, encode bool) string {
	var crypted_symbol string = ""
	var operator int
	if encode == true {
		operator = +1
	} else {
		operator = -1
	}

	if symbol_posion[1] == len(table[0])-1 || symbol_posion[1] == 0 {
		if encode == true {
			crypted_symbol = string(table[symbol_posion[0]][0])
		} else {
			crypted_symbol = string(table[symbol_posion[0]][len(table)-2])
		}
	} else {
		crypted_symbol = string(table[symbol_posion[0]][symbol_posion[1]+operator])
	}
	return crypted_symbol
}

func second_condition(symbol_posion []int, table [][]int, encode bool) string {
	var crypted_symbol string = ""
	var operator int
	if encode == true {
		operator = +1
	} else {
		operator = -1
	}

	if symbol_posion[0] == len(table[0])-1 || symbol_posion[0] == 0 {
		if encode == true {
			crypted_symbol = string(table[0][symbol_posion[1]])
		} else {
			crypted_symbol = string(table[len(table)-2][symbol_posion[1]])
		}
	} else {
		crypted_symbol = string(table[symbol_posion[0]+operator][symbol_posion[1]])
	}
	return crypted_symbol
}

func third_condition(f_symbol_p []int, s_symbol_p []int, table [][]int) string {
	var crypted_symbol string = ""
	crypted_symbol = string(table[f_symbol_p[0]][s_symbol_p[1]])
	return crypted_symbol
}

func Decryptor_Encryptor(text string, table [][]int, key string, encode bool) string {
	var decrypt_text string = ""
	var source_text string = strings.ReplaceAll(strings.ToUpper(text), "J", "I")

	if len(source_text)%2 != 0 {
		source_text += "X"
	}

	for s := 0; s < len(source_text); s += 2 {

		x := int(source_text[s])
		y := 0
		if len(source_text) != s {
			y = int(source_text[s+1])
		}

		x_position := get_position(x, table)
		y_position := get_position(y, table)

		if x_position[0] == y_position[0] {
			decrypt_text += first_condition(x_position, table, encode)
			decrypt_text += first_condition(y_position, table, encode)
		} else if x_position[1] == y_position[1] {
			decrypt_text += second_condition(x_position, table, encode)
			decrypt_text += second_condition(y_position, table, encode)
		} else {
			decrypt_text += third_condition(x_position, y_position, table)
			decrypt_text += third_condition(y_position, x_position, table)
		}
	}

	return decrypt_text
}
