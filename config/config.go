package config

import "fmt"

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
)

type DBConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

func (c *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type FSConfig struct {
	Hostname string `yaml:"hostname"`
}

type Config struct {
	Env   string      `yaml:"env"`
	DB    DBConfig    `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
	FS    FSConfig    `yaml:"fs"`
}

func (c *Config) Develop() bool { return c.Env == Develop }
func (c *Config) Debug() bool   { return c.Env == Debug }
func (c *Config) Produce() bool { return c.Env == Produce }

func (c *Config) KEY() []byte {
	switch c.Env {
	case Produce:
		return []byte{234, 211, 45, 47, 242, 70, 173, 220, 26, 163, 56, 176, 128, 191, 166, 136}
	default:
		return []byte{41, 66, 164, 19, 56, 145, 63, 42, 118, 59, 163, 61, 180, 27, 45, 43}
	}
}
