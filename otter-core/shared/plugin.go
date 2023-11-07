package shared

import (
	"context"
	"github.com/haetao/otter-core/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type IModulePlugin interface {
	Init()
	Run()
	Stop()
}

type ModuleClient struct {
	client proto.IPluginClient
}

func (m *ModuleClient) Init() {
	m.client.Init(context.Background(), &proto.Empty{})
}

func (m *ModuleClient) Run() {
	m.client.Run(context.Background(), &proto.Empty{})
}

func (m *ModuleClient) Stop() {
	m.client.Stop(context.Background(), &proto.Empty{})
}

type ModuleServer struct {
	Impl IModulePlugin
}

func (m *ModuleServer) Init(context.Context, *proto.Empty) (*proto.Empty, error) {
	m.Impl.Init()
	return &proto.Empty{}, nil
}

func (m *ModuleServer) Run(context.Context, *proto.Empty) (*proto.Empty, error) {
	m.Impl.Run()
	return &proto.Empty{}, nil
}

func (m *ModuleServer) Stop(context.Context, *proto.Empty) (*proto.Empty, error) {
	m.Impl.Stop()
	return &proto.Empty{}, nil
}

type ModulePlugin struct {
	plugin.GRPCPlugin
	Impl IModulePlugin
}

func (m *ModulePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterIPluginServer(s, &ModuleServer{
		Impl: m.Impl,
	})
	return nil
}

func (m *ModulePlugin) GRPCClient(c context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &ModuleClient{
		client: proto.NewIPluginClient(conn),
	}, nil
}
