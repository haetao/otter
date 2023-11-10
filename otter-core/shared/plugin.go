package shared

import (
	"context"
	"github.com/haetao/otter-core/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type IModulePlugin interface {
	Init(data []byte)
	Run()
	Stop()
}

// ============================================= GRPC ================================================

type ModuleGrpcClient struct {
	client proto.IPluginClient
}

func (m *ModuleGrpcClient) Init(data []byte) {
	m.client.Init(context.Background(), &proto.InitRequest{Data: data})
}

func (m *ModuleGrpcClient) Run() {
	m.client.Run(context.Background(), &proto.Empty{})
}

func (m *ModuleGrpcClient) Stop() {
	m.client.Stop(context.Background(), &proto.Empty{})
}

type ModuleGrpcServer struct {
	Impl IModulePlugin
}

func (m *ModuleGrpcServer) Init(_ context.Context, req *proto.InitRequest) (*proto.Empty, error) {
	m.Impl.Init(req.Data)
	return &proto.Empty{}, nil
}

func (m *ModuleGrpcServer) Run(context.Context, *proto.Empty) (*proto.Empty, error) {
	m.Impl.Run()
	return &proto.Empty{}, nil
}

func (m *ModuleGrpcServer) Stop(context.Context, *proto.Empty) (*proto.Empty, error) {
	m.Impl.Stop()
	return &proto.Empty{}, nil
}

type ModuleGrpcPlugin struct {
	plugin.Plugin
	Impl IModulePlugin
}

func (m *ModuleGrpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterIPluginServer(s, &ModuleGrpcServer{
		Impl: m.Impl,
	})
	return nil
}

func (m *ModuleGrpcPlugin) GRPCClient(c context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &ModuleGrpcClient{
		client: proto.NewIPluginClient(conn),
	}, nil
}

// ============================================= GRPC ================================================
