package service

type Config struct {
	//Telemetry TelemetryConfig `json:"telemetry" yaml:"telemetry" toml:"telemetry"`
	//Registry  registry.Config `json:"registry"  yaml:"registry"  toml:"registry"`
}

type TelemetryConfig struct {
	Enable bool `json:"enable" yaml:"enable" toml:"enable"`
}
