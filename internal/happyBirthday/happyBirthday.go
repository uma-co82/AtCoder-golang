package happyBirthday

import "fmt"

func Of(valueOfA, valueOfB int) {
	err := NewNumberOfCakes(valueOfA, valueOfB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Yay!")
	return
}
