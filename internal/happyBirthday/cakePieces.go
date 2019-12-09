package happyBirthday

type cakePieces struct {
	values []int
}

func newCakePieces() cakePieces {
	cakes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	return cakePieces{values: cakes}
}

/**
 * １つ食べる
 */
func (cakePieces *cakePieces) eatCake(removeNum int) {
	var result []int
	for _, v := range cakePieces.values {
		if v != removeNum {
			result = append(result, v)
		}
	}
	cakePieces.values = result
}

func (cakePieces *cakePieces) getValues() []int {
	return cakePieces.values
}
