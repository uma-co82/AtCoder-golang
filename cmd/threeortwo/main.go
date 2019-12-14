package main

import (
	"fmt"
	"os"
	"strconv"

	"../../internal/cmd/threeortwo"
	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
	numbers, err := argumentValidate(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answer, err := threeortwo.Of(numbers[0], numbers[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("A: %d", answer)
}

func argumentValidate(args []string) ([]int, error) {
	var numbers []int

	if len(args) <= 0 {
		err := pkg.NewError("引数を入力して下さい")
		return nil, err
	}
	
	for _, val := range args {
		num, ok := strconv.Atoi(val)
		if ok != nil {
			err := pkg.NewError("引数は整数値のみ有効です")
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
