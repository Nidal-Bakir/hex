package api

import "github.com/Nidal-Bakir/hex/internal/ports"

type Adapter struct {
	arith ports.ArithmaticPort
}

func NewAdapter(arith ports.ArithmaticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (ad Adapter) GetAdd(a, b int) (int, error) {
	return ad.arith.Add(a, b)
}

func (ad Adapter) GetSub(a, b int) (int, error) {
	return ad.arith.Sub(a, b)
}

func (ad Adapter) GetMul(a, b int) (int, error) {
	return ad.arith.Mul(a, b)
}

func (ad Adapter) GetDiv(a, b int) (float64, error) {
	return ad.arith.Div(a, b)

}
