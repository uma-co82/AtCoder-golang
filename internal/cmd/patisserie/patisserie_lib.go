package patisserie

import "github.com/uma-co82/AtCoder-golang/pkg"

func Of(valueOfN, valueOfM int, valueOfA []int64) error {
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

	return nil
}

func argumentValidate(valueOfN, valueOfM int, valueOfA []int64) error {
	if len(valueOfA)/3 != valueOfN {
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
