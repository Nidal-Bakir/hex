package arithmatic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ad Adapter) Add(a int, b int) (int, error) {
	return a + b, nil
}

func (ad Adapter) Sub(a int, b int) (int, error) {
	return a - b, nil
}

func (ad Adapter) Mul(a int, b int) (int, error) {
	return a * b, nil
}

func (ad Adapter) Div(a int, b int) (float64, error) {
	return float64(a) / float64(b), nil
}
