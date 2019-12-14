package main

import (
	"fmt"
	"os"

	"github.com/uma-co82/AtCoder-golang/pkg"
	"github.com/uma-co82/AtCoder-golang/pkg/cmd/ringofavoritenumbers"
)

func main() {
	args := os.Args[1:]
	numbers, err := pkg.ArgumentValidate(args, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answer, err := ringofavoritenumbers.Of(numbers[0], numbers[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("A: %d", int(answer))
}
