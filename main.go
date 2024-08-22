package main

import (
	"fmt"
)

func view(field [9]string) (string, bool) {
	switch {
	case field[0] == "X" && field[1] == "X" && field[2] == "X":
		return "X Выйграли!", true
	case field[3] == "X" && field[4] == "X" && field[5] == "X":
		return "X Выйграли!", true
	case field[6] == "X" && field[7] == "X" && field[8] == "X":
		return "X Выйграли!", true
	case field[0] == "X" && field[3] == "X" && field[6] == "X":
		return "X Выйграли!", true
	case field[1] == "X" && field[4] == "X" && field[7] == "X":
		return "X Выйграли!", true
	case field[2] == "X" && field[5] == "X" && field[8] == "X":
		return "X Выйграли!", true
	case field[0] == "X" && field[4] == "X" && field[8] == "X":
		return "X Выйграли!", true
	case field[2] == "X" && field[4] == "X" && field[6] == "X":
		return "X Выйграли!", true
	case field[0] == "0" && field[1] == "0" && field[2] == "0":
		return "0 Выйграли!", true
	case field[3] == "0" && field[4] == "0" && field[5] == "0":
		return "0 Выйграли!", true
	case field[6] == "0" && field[7] == "0" && field[8] == "0":
		return "0 Выйграли!", true
	case field[0] == "0" && field[3] == "0" && field[6] == "0":
		return "0 Выйграли!", true
	case field[1] == "0" && field[4] == "0" && field[7] == "0":
		return "0 Выйграли!", true
	case field[2] == "0" && field[5] == "0" && field[8] == "0":
		return "0 Выйграли!", true
	case field[0] == "0" && field[4] == "0" && field[8] == "0":
		return "0 Выйграли!", true
	case field[2] == "0" && field[4] == "0" && field[6] == "0":
		return "0 Выйграли!", true
	default:
		return "Слудующий ход", false
	}
}

func main() {
	var coord int
	field := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Ход X, выберите номер клетки:")
			fmt.Scan(&coord)
			field[coord-1] = "X"
		} else {
			fmt.Println("Ход 0, выберите номер клетки:")
			fmt.Scan(&coord)
			field[coord-1] = "0"
		}
		fmt.Printf("%s\n%s\n%s\n", field[0:3], field[3:6], field[6:])
		a, b := view(field)
		if b == true {
			fmt.Println(a)
			break
		}
		if i == 9 && b == false {
			fmt.Println("Ничья!")
		}

	}

}
