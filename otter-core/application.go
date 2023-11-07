package otter_core

import (
	"github.com/hashicorp/go-plugin"
	"os"
	"os/exec"
)

type Application struct {
	config *Config
	client *plugin.Client
	opts   *Options
}

func (a *Application) Init() {
	a.client = plugin.NewClient(&plugin.ClientConfig{
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
