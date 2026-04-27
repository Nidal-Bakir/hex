package arithmatic

import "context"

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ad Adapter) Add(ctx context.Context, a int, b int) (int, error) {
	return a + b, nil
}

func (ad Adapter) Sub(ctx context.Context, a int, b int) (int, error) {
	return a - b, nil
}

func (ad Adapter) Mul(ctx context.Context, a int, b int) (int, error) {
	return a * b, nil
}

func (ad Adapter) Div(ctx context.Context, a int, b int) (float64, error) {
	return float64(a) / float64(b), nil
}
