package config

import "fmt"

type System struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Env    string `yaml:"env"`
	SslPem string `yaml:"ssl-pem"`
	SslKey string `yaml:"ssl-key"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
