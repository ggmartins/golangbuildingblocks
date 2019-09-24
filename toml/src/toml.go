package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Config Configuration
}

type Configuration struct {
	Control_server_ipaddr string
	Control_server_port   int
}

func main() {
	var c tomlConfig
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("control_server_ipaddr: %s\n", c.Config.Control_server_ipaddr)
	fmt.Printf("control_server_port: %d\n", c.Config.Control_server_port)
}
