package ports

type ArithmaticPort interface {
	Add(a int, b int) (int, error)
	Sub(a int, b int) (int, error)
	Mul(a int, b int) (int, error)
	Div(a int, b int) (float64, error)
}

