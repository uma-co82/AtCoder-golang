package patisserie

import (
	"fmt"
	"sort"

	"github.com/uma-co82/AtCoder-golang/pkg"
)

func Of(valueOfN, valueOfM int64, valueOfA []int64) error {
	if err := argumentValidate(valueOfN, valueOfM, valueOfA); err != nil {
		return err
	}
	values := [][]int64{}
	tmp := []int64{}

	for _, n := range valueOfA {
		tmp = append(tmp, n)

		if len(tmp) == 3 {
			values = append(values, tmp)
			tmp = []int64{}
		}
	}

	result1 := []int64{}
	result2 := []int64{}
	result3 := []int64{}
	result4 := []int64{}
	result5 := []int64{}
	result6 := []int64{}
	result7 := []int64{}
	result8 := []int64{}

	for _, n := range values {
		result1 = append(result1, calculation1(n))
		result2 = append(result2, calculation2(n))
		result3 = append(result3, calculation3(n))
		result4 = append(result4, calculation4(n))
		result5 = append(result5, calculation5(n))
		result6 = append(result6, calculation6(n))
		result7 = append(result7, calculation7(n))
		result8 = append(result8, calculation8(n))
	}

	results := [][]int64{}

	results = append(results, result1)
	results = append(results, result2)
	results = append(results, result3)
	results = append(results, result4)
	results = append(results, result5)
	results = append(results, result6)
	results = append(results, result7)
	results = append(results, result8)

	for _, result := range results {
		descendingOrder(result)
	}

	answer := []int64{}

	for _, result := range results {
		answer = append(answer, sliceTotal(result[:valueOfM]))
	}

	descendingOrder(answer)

	fmt.Printf("A: %d \n", answer[0])

	return nil
}

func argumentValidate(valueOfN, valueOfM int64, valueOfA []int64) error {
	if int64(len(valueOfA)/3) != valueOfN {
		return pkg.NewError("Nと種類の数が違うので演算が行えません")
	}

	if len(valueOfA)%3 != 0 {
		return pkg.NewError("種類の個数が不正です")
	}

	if !(1 <= valueOfN) || !(valueOfN <= 1000) {
		return pkg.NewError("Nには1以上1000以下の値を入力して下さい")
	}

	if !(0 <= valueOfM) || valueOfN <= valueOfM {
		return pkg.NewError("Mは0以上N以下の値を入力して下さい")
	}

	for _, n := range valueOfA {
		if !(-10000000000 <= n) || !(n <= 10000000000) {
			return pkg.NewError("aには-10000000000以上10000000000以下の値を入力して下さい")
		}
	}

	return nil
}

func calculation1(x []int64) int64 {
	return x[0] + x[1] + x[2]
}

func calculation2(x []int64) int64 {
	return x[0] + x[1] - x[2]
}

func calculation3(x []int64) int64 {
	return x[0] - x[1] + x[2]
}

func calculation4(x []int64) int64 {
	return x[0] - x[1] - x[2]
}

func calculation5(x []int64) int64 {
	return -x[0] + x[1] + x[2]
}

func calculation6(x []int64) int64 {
	return -x[0] - x[1] - x[2]
}

func calculation7(x []int64) int64 {
	return -x[0] - x[1] + x[2]
}

func calculation8(x []int64) int64 {
	return -x[0] + x[1] - x[2]
}

func descendingOrder(x []int64) {
	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})
}

func sliceTotal(x []int64) int64 {
	var total int64

	for _, n := range x {
		total = total + n
	}

	return total
}
