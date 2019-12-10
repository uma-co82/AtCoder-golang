package main

import (
	"fmt"
	"os"

	"../../internal/cmd/ringofavoritenumbers"
	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
	numbers, err := pkg.ArgumentValidate(args, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answer, err := ringofavoritenumbers.Of(numbers[0], numbers[1])
	fmt.Printf("A: %d", int(answer))
}
