package main

import (
	"fmt"
	"os"

	"../../internal/cmd/happybirthday"
	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
	numbers, err := pkg.ArgumentValidate(args, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	happybirthday.Of(numbers[0], numbers[1])
}
