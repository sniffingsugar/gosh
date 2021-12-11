package config

var GConfig *Config

type Config struct {
	Created bool   `json:"created"`
	Server  string `json:"server"`
}
