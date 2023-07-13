package configs

import (
	"gopkg.in/ini.v1"
)

var Instance *Config

func init() {
	Instance = new(Config)
	Instance.Init()
}

type Config struct {
	configs map[string]*ini.File
}

func (this *Config) Init() {
	this.configs = make(map[string]*ini.File)
}

func (this *Config) Get(key string, path string) (string, error) {
	_, exists := this.configs[path]
	if !exists {
		var err error
		this.configs[path], err = ini.Load("configs/" + path)
		if err != nil {
			return "", err
		}
	}

	return this.configs[path].Section("").Key(key).String(), nil
}
