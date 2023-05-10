package rpc

import "github.com/cro4k/authorize/common/services/service"

type Config struct {
	service.Config

	Addr      string                  `json:"addr" yaml:"addr" toml:"addr"`
	Telemetry service.TelemetryConfig `json:"telemetry" yaml:"telemetry" toml:"telemetry"`
}
