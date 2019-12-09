package happyBirthday

import (
	"github.com/uma-co82/AtCoder-golang/pkg"
)

type numberOfCakes struct {
	valueOfA int
	valueOfB int
}

const totalCakes = 16

func NewNumberOfCakes(valueOfA, valueOfB int) (*numberOfCakes, error) {
	if valueOfA+valueOfB > totalCakes {
		err := pkg.NewError("総数を16以下にして下さい")
		return nil, err
	}

	if valueOfA >= 9 || valueOfB >= 9 {
		err := pkg.NewError(":(")
		return nil, err
	}

	return &numberOfCakes{
		valueOfA: valueOfA,
		valueOfB: valueOfB,
	}, nil
}

func (numberOfCakes *numberOfCakes) GetValueOfA() int {
	return numberOfCakes.valueOfA
}

func (numberOfCakes *numberOfCakes) GetValueOfB() int {
	return numberOfCakes.valueOfB
}
