package api

import "github.com/cro4k/authorize/common/services/service"

type Config struct {
	service.Config
	Addr      string                  `json:"addr"      yaml:"addr"      toml:"addr"`
	Timeout   int                     `json:"timeout"   yaml:"timeout"   toml:"timeout"`
	Telemetry service.TelemetryConfig `json:"telemetry" yaml:"telemetry" toml:"telemetry"`
}

func (c *Config) setDefault() {
	if c.Addr == "" {
		c.Addr = "0.0.0.0:8088"
	}
	if c.Timeout == 0 {
		c.Timeout = 5000
	}
}
