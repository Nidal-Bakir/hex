package ports

import "context"

type DbPort interface {
	CloseCon(ctx context.Context)
	AppendHistory(ctx context.Context, operation string, answer float64) error
}
