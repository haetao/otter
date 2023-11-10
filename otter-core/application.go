package otter_core

import (
	"github.com/haetao/otter-core/shared"
	"github.com/hashicorp/go-plugin"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
)

type Application struct {
	config *Config
	opts   *Options
}

func (a *Application) AppName() string {
	return a.config.AppName
}

func (a *Application) Version() string {
	return a.config.Version
}

func (a *Application) Init() {
	for k, _ := range a.config.ModConfigs {
		a.opts.PluginMap[k] = &shared.ModuleGrpcPlugin{}
	}
	a.opts.Client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: a.opts.Handshake, //TODO 处理默认握手配置
		Plugins:         a.opts.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("OTTER_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
	})
}

func (a *Application) Run() {

}

func (a *Application) Stop() {

}

func NewApplication() *Application {
	config := loadConfig()
	return &Application{
		config: config,
	}
}

func loadConfig() *Config {
	var config = &Config{}
	data, err := os.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	return config
}
