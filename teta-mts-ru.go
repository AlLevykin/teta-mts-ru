package main

import (
	"fmt"
	"os"
	"teta-mts-ru/pkg/utils"
)

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Fscan(os.Stdin, &s)
	fmt.Println(utils.Counter(s))
}
