package api

import (
	"context"

	"github.com/Nidal-Bakir/hex/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmaticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (ad Adapter) GetAdd(ctx context.Context, a, b int) (int, error) {
	res, err := ad.arith.Add(ctx, a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AppendHistory(ctx, "add", float64(res))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (ad Adapter) GetSub(ctx context.Context, a, b int) (int, error) {
	res, err := ad.arith.Sub(ctx, a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AppendHistory(ctx, "sub", float64(res))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (ad Adapter) GetMul(ctx context.Context, a, b int) (int, error) {
	res, err := ad.arith.Mul(ctx, a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AppendHistory(ctx, "mul", float64(res))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (ad Adapter) GetDiv(ctx context.Context, a, b int) (float64, error) {
	res, err := ad.arith.Div(ctx, a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AppendHistory(ctx, "div", float64(res))
	if err != nil {
		return 0, err
	}
	return res, nil
}
