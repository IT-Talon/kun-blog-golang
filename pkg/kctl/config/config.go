package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	ServerHost    string `yaml:"server_host"`
	TimeOutSecond int    `yaml:"time_out_second"`
}

func LoadConfigFile() *Config {
	configFilePath := getConfigFilePath()
	cfg := &Config{}
	err := yaml.Unmarshal(MustLoadFile(configFilePath), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func MustLoadFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
func getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configFile := filepath.Join(home, "/.kblog/config")
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		log.Fatalln(configFile, "配置文件未找到")
	}
	return configFile
}
