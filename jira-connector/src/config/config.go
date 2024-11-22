package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database  DataBaseConfig  `yaml:"data_base"`
	Connector ConnectorConfig `yaml:"connector"`
}

func (c *Config) Parse() error {
	configFileParse := flag.String("config", "config/config.yaml", "Path to the configuration file")
	flag.Parse()

	data, err := os.ReadFile(*configFileParse)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}
