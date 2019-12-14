package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/uma-co82/AtCoder-golang/pkg"
	"github.com/uma-co82/AtCoder-golang/pkg/cmd/patisserie"
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
 * 引数がint型か判定する
 */
func argumentValidate(args []string) ([]int64, error) {
	var numbers []int64

	if len(args) <= 0 {
		err := pkg.NewError("引数を入力して下さい")
		return nil, err
	}

	for _, val := range args {
		num, ok := strconv.Atoi(val)
		if ok != nil {
			err := pkg.NewError("引数は数値のみ有効です")
			return nil, err
		}
		tmp := int64(num)
		numbers = append(numbers, tmp)
	}
	return numbers, nil
}
