package grpc

import (
	"context"

	"github.com/Nidal-Bakir/hex/internal/adapters/framwork/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ad Adapter) GetAdd(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error) {
	if params.GetA() == 0 || params.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments bro....")
	}
	answer, err := ad.api.GetAdd(ctx, int(params.GetA()), int(params.GetB()))
	if err != nil {
		return nil, status.Error(codes.Internal, "my bad bro...")
	}
	return &pb.Answer{Value: float64(answer)}, nil
}

func (ad Adapter) GetSub(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error) {
	if params.GetA() == 0 || params.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments bro....")
	}
	answer, err := ad.api.GetSub(ctx, int(params.GetA()), int(params.GetB()))
	if err != nil {
		return nil, status.Error(codes.Internal, "my bad bro...")
	}
	return &pb.Answer{Value: float64(answer)}, nil
}

func (ad Adapter) GetMul(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error) {
	if params.GetA() == 0 || params.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments bro....")
	}
	answer, err := ad.api.GetMul(ctx, int(params.GetA()), int(params.GetB()))
	if err != nil {
		return nil, status.Error(codes.Internal, "my bad bro...")
	}
	return &pb.Answer{Value: float64(answer)}, nil
}

func (ad Adapter) GetDiv(ctx context.Context, params *pb.OperationParametars) (*pb.Answer, error) {
	if params.GetA() == 0 || params.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments bro....")
	}
	answer, err := ad.api.GetDiv(ctx, int(params.GetA()), int(params.GetB()))
	if err != nil {
		return nil, status.Error(codes.Internal, "my bad bro...")
	}
	return &pb.Answer{Value: answer}, nil
}
