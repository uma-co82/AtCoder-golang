package main

import (
	"os"
	"strconv"

	"github.com/uma-co82/AtCoder-golang/pkg"
)

func main() {
	args := os.Args[1:]
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
