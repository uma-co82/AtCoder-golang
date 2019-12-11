package main

import (
	"fmt"
	"os"
	"strconv"

	"../../internal/cmd/patisserie"
	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
	numbers, err := argumentValidate(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	valueOfN := numbers[0]
	valueOfM := numbers[1]
	result := numbers[2:]

	if err := patisserie.Of(valueOfN, valueOfM, result); err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 引数がint型判定する
 */
func argumentValidate(args []string) ([]int64, error) {
	var numbers []int64

	for _, val := range args {
		num, ok := strconv.Atoi(val)
		if ok != nil {
			err := pkg.NewError("引数は整数値のみ有効です")
			return nil, err
		}
		tmp := int64(num)
		numbers = append(numbers, tmp)
	}
	return numbers, nil
}
