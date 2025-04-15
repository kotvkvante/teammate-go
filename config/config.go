package config

import (
  "fmt"
  "io/ioutil"
  "gopkg.in/yaml.v3"
)

type Config struct {
  Addr string `yaml:"addr"`
}

var AppConfig Config

func Init(config_path string) {
  file, err := ioutil.ReadFile(config_path)
  if err != nil {
    fmt.Println("No such file", config_path)
  }

  err = yaml.Unmarshal(file, &AppConfig)
  if err != nil {
    fmt.Println("Incorrect config", config_path)
  }

  fmt.Println(">", AppConfig.Addr)
}
