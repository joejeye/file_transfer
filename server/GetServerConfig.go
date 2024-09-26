package main

import (
	"file_transfer_naive/myutils"
	"fmt"
	"github.com/BurntSushi/toml"
)

type ServerConfig struct {
	DlDir      string `toml:"download_directory"`
	ListenPort string `toml:"listen_port"`
}

func GetServerConfig() ServerConfig {
	var conf ServerConfig
	configFilePath := myutils.MyPathJoin(myutils.GetRootDir(), "server", "config.toml")
	_, err := toml.DecodeFile(configFilePath, &conf)
	if err != nil {
		panic(fmt.Errorf("error decoding config file: %w", err))
	}
	return conf
}
