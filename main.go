package main

import "fmt"

// Не раз решал подобное на литкоде так что быстро справился с задачей
func Computers(n int) string {
	var r string
	switch {
	case n%10 == 1 && n%100 != 11:
		r = "компьютер"
	case n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20):
		r = "компьютера"
	default:
		r = "компьютеров"
	}
	return fmt.Sprintf("%d %s", n, r)
}

func main() {
	fmt.Println(Computers(25))
	fmt.Println(Computers(41))
	fmt.Println(Computers(1048))
}
