package otter_core

import (
	"github.com/hashicorp/go-plugin"
	"os"
	"os/exec"
)

type Application struct {
	config *Config
	opts   *Options
}

func (a *Application) Init() {
	//TODO 使用config初始化opts
	a.opts.Client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: a.opts.Handshake,
		Plugins:         a.opts.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
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
	//TODO
	return nil
}
