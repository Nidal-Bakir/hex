package ports 

type ArithmeticPort interface {
  Add(a int32, b int32) (int32,error)
  Sub(a int32, b int32) (int32,error)
  Mul(a int32, b int32) (int32,error)
  Div(a int32, b int32) (int32,error)
}
