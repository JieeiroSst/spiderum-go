package config

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Config struct {
	Server  ServerConfig
	S3		S3
}

type ServerConfig struct {
	PortServer        string
}

type S3 struct {
	S3Region string
	S3Bucket string
	S3ACL   string
}

func ReadConf(filename string) (*Config, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	return config, nil
}
