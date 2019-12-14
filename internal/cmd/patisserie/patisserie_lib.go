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
	hoge := [][]int64{}
	tmp := []int64{}

	for _, n := range valueOfA {
		tmp = append(tmp, n)

		if len(tmp) == 3 {
			hoge = append(hoge, tmp)
			tmp = []int64{}
		}
	}

	result := []int64{}
	result2 := []int64{}
	result3 := []int64{}
	result4 := []int64{}
	result5 := []int64{}
	result6 := []int64{}
	result7 := []int64{}
	result8 := []int64{}

	for _, n := range hoge {
		total := calculation1(n)
		total2 := calculation2(n)
		total3 := calculation3(n)
		total4 := calculation4(n)
		total5 := calculation5(n)
		total6 := calculation6(n)
		total7 := calculation7(n)
		total8 := calculation8(n)
		result = append(result, total)
		result2 = append(result2, total2)
		result3 = append(result3, total3)
		result4 = append(result4, total4)
		result5 = append(result5, total5)
		result6 = append(result6, total6)
		result7 = append(result7, total7)
		result8 = append(result8, total8)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	sort.Slice(result2, func(i, j int) bool {
		return result2[i] > result2[j]
	})
	sort.Slice(result3, func(i, j int) bool {
		return result3[i] > result3[j]
	})
	sort.Slice(result4, func(i, j int) bool {
		return result4[i] > result4[j]
	})
	sort.Slice(result5, func(i, j int) bool {
		return result5[i] > result5[j]
	})
	sort.Slice(result6, func(i, j int) bool {
		return result6[i] > result6[j]
	})
	sort.Slice(result7, func(i, j int) bool {
		return result7[i] > result7[j]
	})
	sort.Slice(result8, func(i, j int) bool {
		return result8[i] > result8[j]
	})

	fmt.Println("result")
	fmt.Println(result)

	answer := result[:valueOfM]
	answer2 := result2[:valueOfM]
	answer3 := result3[:valueOfM]
	answer4 := result4[:valueOfM]
	answer5 := result5[:valueOfM]
	answer6 := result6[:valueOfM]
	answer7 := result7[:valueOfM]
	answer8 := result8[:valueOfM]

	var foo int64
	var foo2 int64
	var foo3 int64
	var foo4 int64
	var foo5 int64
	var foo6 int64
	var foo7 int64
	var foo8 int64

	for _, n := range answer {
		foo = foo + n
	}
	for _, n := range answer2 {
		foo2 = foo2 + n
	}
	for _, n := range answer3 {
		foo3 = foo3 + n
	}
	for _, n := range answer4 {
		foo4 = foo4 + n
	}
	for _, n := range answer5 {
		foo5 = foo5 + n
	}
	for _, n := range answer6 {
		foo6 = foo6 + n
	}
	for _, n := range answer7 {
		foo7 = foo7 + n
	}
	for _, n := range answer8 {
		foo8 = foo8 + n
	}

	buz := []int64{}

	buz = append(buz, foo)
	buz = append(buz, foo2)
	buz = append(buz, foo3)
	buz = append(buz, foo4)
	buz = append(buz, foo5)
	buz = append(buz, foo6)
	buz = append(buz, foo7)
	buz = append(buz, foo8)

	sort.Slice(buz, func(i, j int) bool {
		return buz[i] > buz[j]
	})

	fmt.Println(buz[0])
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
	return - x[0] + x[1] + x[2]
}

func calculation6(x []int64) int64 {
	return - x[0]  - x[1] - x[2]
}

func calculation7(x []int64) int64 {
	return - x[0] - x[1] + x[2]
}

func calculation8(x []int64) int64 {
	return - x[0] + x[1] - x[2]
}