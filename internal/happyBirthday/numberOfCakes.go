package happyBirthday

import (
	"github.com/uma-co82/AtCoder-golang/pkg"
)

const TOTAL_CAKE_PIECES = 16

func NewNumberOfCakes(valueOfA, valueOfB int) error {
	if valueOfA+valueOfB > TOTAL_CAKE_PIECES {
		err := pkg.NewError("総数を16以下にして下さい")
		return err
	}

	if valueOfA >= 9 || valueOfB >= 9 {
		err := pkg.NewError(":(")
		return err
	}

	return nil
}
