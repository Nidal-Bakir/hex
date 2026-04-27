package ports

type GrpcPort interface {
	Run()
	GetAdd()
	GetSub()
	GetMul()
	GetDiv()
}