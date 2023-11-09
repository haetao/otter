package otter_core

import "github.com/hashicorp/go-plugin"

type Config struct {
	Version  string `yaml:"version"`
	AppName  string `yaml:"appName"`
	Endpoint string `yaml:"endpoint"`
}

type Options struct {
	Client    *plugin.Client
	Handshake plugin.HandshakeConfig
	PluginMap map[string]plugin.Plugin
}

type Option func(cfg *Config)
