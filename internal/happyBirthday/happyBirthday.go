package happyBirthday

import "fmt"

func Of(valueOfA, valueOfB int) {
	var A_EAT_PIECE_NUM = map[int]int{
		1: 1,
		2: 3,
		3: 5,
		4: 7,
		5: 9,
		6: 11,
		7: 13,
		8: 15,
	}
	var B_EAT_PIECE_NUM = map[int]int{
		1: 2,
		2: 4,
		3: 6,
		4: 8,
		5: 10,
		6: 12,
		7: 14,
		8: 16,
	}
	numberOfCakes, err := NewNumberOfCakes(valueOfA, valueOfB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cakePieces := newCakePieces()

	for i := 0; i < numberOfCakes.valueOfA; i++ {
		cakePieces.eatCake(A_EAT_PIECE_NUM[i])
	}

	for i := 0; i < numberOfCakes.valueOfB; i++ {
		cakePieces.eatCake(B_EAT_PIECE_NUM[i])
	}

	if len(cakePieces.getValues()) > 0 {
		fmt.Println(cakePieces.getValues())
		fmt.Println("Yay!")
	} else {
		fmt.Println(cakePieces.getValues())
		fmt.Println(":(")
	}
}
