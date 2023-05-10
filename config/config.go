package config

import (
	"github.com/cro4k/authorize/common/services/api"
	"github.com/cro4k/authorize/common/services/rpc"
	"github.com/cro4k/authorize/internal/db"
)

type Config struct {
	DB  db.Config  `json:"db" yaml:"db" toml:"db"`
	API api.Config `json:"api" yaml:"api" toml:"api"`
	RPC rpc.Config `json:"rpc" yaml:"rpc" toml:"rpc"`
}
