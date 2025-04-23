package workenv

import (
	"context"
	"fmt"
	"github.com/9triver/cfn/proto"
	pb "github.com/9triver/cfn/proto/data"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

type Controller struct {
	envManager *VenvManager
	logger     *slog.Logger
	pb.UnimplementedFunctionServiceServer
}

func NewController(host string, port string) *Controller {
	manager, err := NewManager(context.Background())
	if err != nil {
		return nil
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	controller := &Controller{
		envManager: manager,
		logger:     logger,
	}

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterFunctionServiceServer(grpcServer, controller)

	logger.Info("gRPC server listening on " + host + ":" + port)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	return controller
}

func (controller *Controller) DeployPyFunc(ctx context.Context, pyFunc *pb.AppendPyFunc) (*proto.ServiceReplay, error) {
	// 初始化虚拟环境
	venv, err := controller.envManager.GetVenv(pyFunc.GetName(), pyFunc.GetRequirements()...)
	if err != nil {
		controller.logger.Error(fmt.Sprint(err))
		return nil, err
	}
	controller.logger.Info("虚拟环境初始化完成\n\t" + fmt.Sprint(venv))
	return &proto.ServiceReplay{}, nil
}
