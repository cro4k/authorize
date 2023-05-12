package rpc

import "github.com/cro4k/authorize/common/services/service"

type Config struct {
	service.Config
	Name      string                  `json:"name"      yaml:"name"      toml:"name"`
	Addr      string                  `json:"addr"      yaml:"addr"      toml:"addr"`
	Telemetry service.TelemetryConfig `json:"telemetry" yaml:"telemetry" toml:"telemetry"`
}
