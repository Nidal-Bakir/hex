package main

import (
	"context"
	"os"

	"github.com/Nidal-Bakir/hex/internal/adapters/app/api"
	"github.com/Nidal-Bakir/hex/internal/adapters/core/arithmatic"
	"github.com/Nidal-Bakir/hex/internal/adapters/framwork/left/grpc"
	"github.com/Nidal-Bakir/hex/internal/adapters/framwork/right/db"
	"github.com/Nidal-Bakir/hex/internal/ports"
)

var (
	database     = os.Getenv("DB_DATABASE")
	password     = os.Getenv("DB_PASSWORD")
	username     = os.Getenv("DB_USERNAME")
	port         = os.Getenv("DB_PORT")
	host         = os.Getenv("DB_HOST")
	poolMaxConns = os.Getenv("DB_POOL_MAX_CONNS")
)

func main() {
	ctx := context.TODO()

	var dbPort ports.DbPort
	var gRPCPort ports.GrpcPort
	var apiPort ports.ApiPort
	var arithmaticPort ports.ArithmaticPort

	dbPort = db.NewAdapter(ctx, username, password, host, port, database, poolMaxConns)
	defer dbPort.CloseCon(ctx)
	arithmaticPort = arithmatic.NewAdapter()
	apiPort = api.NewAdapter(dbPort, arithmaticPort)
	gRPCPort = grpc.NewAdapter(apiPort)
	gRPCPort.Run()
}
