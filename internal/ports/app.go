package ports

type ApiPort interface {
	GetAdd(a, b int) (int, error)
	GetSub(a, b int) (int, error)
	GetMul(a, b int) (int, error)
	GetDiv(a, b int) (float64, error)
}
