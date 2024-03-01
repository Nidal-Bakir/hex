package arithmetic

import "errors"

var ErrDividingByZero = errors.New("math: can not divide by zero")

type Adapter struct {
}

func (ad Adapter) Add(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (ad Adapter) Sub(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (ad Adapter) Mul(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (ad Adapter) Div(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, ErrDividingByZero
	}
	return a / b, nil
}
