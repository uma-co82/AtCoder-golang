package pkg

import (
	"strconv"
)

/**
 * 引数がint型か判定する
 */
func ArgumentValidate(args []string) ([]int, error) {
	var numbers []int
	for _, val := range args {
		num, ok := strconv.Atoi(val)
		if ok != nil {
			err := NewError("引数はint型のみ有効です")
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
