package ports

import (
	"context"

	"github.com/Nidal-Bakir/hex/internal/adapters/framwork/left/grpc/pb"
)

type GrpcPort interface {
	Run()
	GetAdd(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error)
	GetSub(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error)
	GetMul(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error)
	GetDiv(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error)
}
