package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-ApmCommon/internal"
	"sync"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Service   string
	Logconfig logconfig
	Databases map[string]databases
	Servers   map[string]servers
}

type logconfig struct {
	Logpath  string
	Loglevel logrus.Level
}

func (l *logconfig) UnmarshalTOML(p interface{}) error {
	data, _ := p.(map[string]interface{})
	l.Logpath, _ = data["logpath"].(string)
	//Parsing constant values
	var err error
	l.Loglevel, err = logrus.ParseLevel(data["loglevel"].(string))
	if err != nil {
		return err
	}
	return nil
}

type servers struct {
	IP   string
	PORT string
}

type databases struct {
	Server string
	Port   string
	Enable bool
}

// Parsing toml
var config TomlConfig

func Load(cp string) {
	cmdargs := internal.GetCmdargs()
	fpath := fmt.Sprintf(cp, cmdargs.Phase)
	if _, err := toml.DecodeFile(fpath, &config); err != nil {
		fmt.Println(err)
	}
}

func (t *TomlConfig) ApmServerUrl() string {
	return fmt.Sprintf("%s%s", t.Servers["APM"].IP, t.Servers["APM"].PORT)
}

// Singletone
var insTomlConfig *TomlConfig
var onceTomlConfig sync.Once

func GetConfig() *TomlConfig {
	onceTomlConfig.Do(func() {
		insTomlConfig = &config
	})
	return insTomlConfig
}
