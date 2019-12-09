package pkg

/**
 * 引数がint型か判定する
 */
func ArgumentValidate(args ...interface{}) error {
	for _, val := range args {
		if _, ok := val.(int); !ok {
			err := NewError("引数はint型のみ有効です")
			return err
		}
	}
	return nil
}
