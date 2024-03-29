package threeortwo

import (
	"github.com/uma-co82/AtCoder-golang/pkg"
)

func Of(valueOfN int, valueOfA []int) (int, error) {
	if err := argumentValidate(valueOfN, valueOfA); err != nil {
		return 0, err
	}

	operationCount := 0

	for i := 0; i < valueOfN; i++ {
		quotient := countDivisionBy2(valueOfA[i])
		operationCount = operationCount + quotient
	}

	return operationCount, nil
}

func argumentValidate(valueOfN int, valueOfA []int) error {
	if len(valueOfA) != valueOfN {
		return pkg.NewError("Nとaiの個数が違うので演算が行えません")
	}

	if !(1 <= valueOfN) || !(valueOfN <= 10000) {
		return pkg.NewError("Nには1以上10000以下の値を入力して下さい")
	}

	for _, n := range valueOfA {
		if !(1 <= n) || !(n <= 1000000000) {
			return pkg.NewError("aには1以上1000000000以下の値を入力して下さい")
		}
	}

	return nil
}

func countDivisionBy2(numerator int) int {
	n := numerator
	count := 0

	for n%2 == 0 {
		n = n / 2
		count++
	}

	return count
}
