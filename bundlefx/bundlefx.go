package bundlefx

import (
	"context"
	"github.com/rishabhgpt/fx-app/configfx"
	"github.com/rishabhgpt/fx-app/handler"
	"github.com/rishabhgpt/fx-app/loggerfx"
	pb "github.com/rishabhgpt/fx-app/proto"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var Module = fx.Options(configfx.Module, loggerfx.Module, fx.Invoke(registerHooks))

func registerHooks(lc fx.Lifecycle, cfg *configfx.Config, logger *zap.SugaredLogger, rpcServer handler.Handler) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on ", cfg.Address)
				lis, _ := net.Listen("tcp", cfg.Address)
				var opts []grpc.ServerOption
				grpcServer := grpc.NewServer(opts...)
				pb.RegisterMessageServiceServer(grpcServer, rpcServer)
				reflection.Register(grpcServer)
				go grpcServer.Serve(lis)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)

}
