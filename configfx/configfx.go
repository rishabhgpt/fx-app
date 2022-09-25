package configfx

import (
	"io/ioutil"

	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApplicationConfig `yaml:"application"`
}

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

func ConfigProvider() (*Config, error) {
	conf := &Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	return conf, err
}

var Module = fx.Options(fx.Provide(ConfigProvider))
