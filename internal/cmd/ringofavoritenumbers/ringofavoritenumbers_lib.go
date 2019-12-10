package ringofavoritenumbers

import (
	"math"

	"github.com/uma-co82/AtCoder-golang/pkg"
)

func Of(valueOfD, valueOfN int) (float64, error) {
	if err := argumentValidate(valueOfD, valueOfN); err != nil {
		return 0, err
	}

	if valueOfD == 0 {
		return float64(valueOfN), nil
	}

	return math.Pow(100, float64(valueOfD)) * float64(valueOfN), nil
}

func argumentValidate(valueOfD, valueOfN int) error {
	if !(0 <= valueOfD) && !(valueOfD <= 2) {
		return pkg.NewError("Dには0,1,2いずれかの値を入力して下さい")
	}

	if !(0 <= valueOfN) && !(valueOfN <= 100) {
		return pkg.NewError("Nには1以上100以下の値を入力して下さい")
	}

	return nil
}
