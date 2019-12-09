package main

import (
	"fmt"
	"os"

	"github.com/uma-co82/AtCoder-golang/internal/happyBirthday"

	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
	numbers, err := pkg.ArgumentValidate(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	happyBirthday.Of(numbers[0], numbers[1])
}
