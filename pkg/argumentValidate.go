package pkg

import (
	"fmt"
	"strconv"
)

/**
 * 引数がint型か規定数以上与えられたか判定する
 */
func ArgumentValidate(args []string, argsCount int) ([]int, error) {
	var numbers []int

	if len(args) < argsCount {
		return nil, NewError(fmt.Sprintf("引数が足りません%d個入力して下さい", argsCount))
	}

	for _, val := range args {
		num, ok := strconv.Atoi(val)
		if ok != nil {
			err := NewError("引数は整数値のみ有効です")
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
