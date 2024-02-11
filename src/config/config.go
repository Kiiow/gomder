package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

/* Views */
type View int

const (
	DirectoryIndex View = iota
	TerminalIndex
)

/* Style */
const (
	DirectoryWidthPercent = 30
	TerminalWidthPercent  = 70
)

/* Config File */

var (
	Config *ConfigIni
)

func createDefaultConfig() *ConfigIni {
	defaultConfig := &ConfigIni{}
	defaultConfig.Application.Name = "Gmder"
	return defaultConfig
}

func Load() {
	inidata, err := ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Failed to read ini file: %v", err)
		Config = createDefaultConfig()
		return
	}

	Config = &ConfigIni{}
	err = inidata.MapTo(Config)
	if err != nil {
		fmt.Printf("Failed to map ini file: %v", err)
		os.Exit(1)
	}
}

type ConfigIni struct {
	Application struct {
		Name string `ini:"APP_NAME"`
	} `ini:"application"`
}

const DefaultViewIndex = DirectoryIndex
