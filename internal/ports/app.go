package ports

import "context"

type ApiPort interface {
	GetAdd(ctx context.Context, a, b int) (int, error)
	GetSub(ctx context.Context, a, b int) (int, error)
	GetMul(ctx context.Context, a, b int) (int, error)
	GetDiv(ctx context.Context, a, b int) (float64, error)
}
