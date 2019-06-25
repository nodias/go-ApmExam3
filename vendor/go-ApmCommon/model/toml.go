package model

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Title     string
	Writer    writerinfo
	Databases map[string]databases
	Servers   map[string]servers
	Logpaths  map[string]logpaths
}

type writerinfo struct {
	Name string
}

type databases struct {
	Server string
	Port   string
	Enable bool
}

type servers struct {
	IP   string
	PORT string
}

type logpaths struct {
	Path string
}

func (t *TomlConfig) New(path string) {
	if _, err := toml.DecodeFile("config.toml", &t); err != nil {
		fmt.Println(err)
	}
}
