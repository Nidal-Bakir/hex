package ports

import "context"

type ArithmaticPort interface {
	Add(ctx context.Context, a int, b int) (int, error)
	Sub(ctx context.Context, a int, b int) (int, error)
	Mul(ctx context.Context, a int, b int) (int, error)
	Div(ctx context.Context, a int, b int) (float64, error)
}
