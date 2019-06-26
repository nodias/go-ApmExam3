package model

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Service   string
	Logpaths  logpaths
	Databases map[string]databases
	Servers   map[string]servers
}

type servers struct {
	IP   string
	PORT string
}

type logpaths struct {
	Logpath string
}

type databases struct {
	Server string
	Port   string
	Enable bool
}

// TomlConfig.Load is a function, select config with args
func (t *TomlConfig) Load() {
	var phase string
	if len(os.Args) < 2 {
		phase = "local"
	}
	phase = os.Args[1]
	fpath := fmt.Sprintf("config/%s/config.toml", phase)
	if _, err := toml.DecodeFile(fpath, &t); err != nil {
		fmt.Println(err)
	}
}

func (t *TomlConfig) ApmServerUrl() string {
	return fmt.Sprintf("%s%s", t.Servers["APM_TESTSERVER"].IP, t.Servers["APM_TESTSERVER"].PORT)
}
