package conf

import (
	"sync"

	"github.com/caarlos0/env"
)

type EnvironmentConfig struct {
	LogLevel      string `env:"LOG_LEVEL" envDefault:"info"`
	ServerPort    int    `env:"SERVER_PORT" envDefault:"8888"`
	RaftNodeId    string `env:"RAFT_NODE_ID" envDefault:"node_id"`
	RaftPort      int    `env:"RAFT_PORT" envDefault:"7777"`
	RaftVolumeDir string `env:"RAFT_VOLUME_DIR" envDefault:"data"`
}

var instantiated *EnvironmentConfig
var once sync.Once

// EnvironmentConfigIns environment config
func EnvironmentConfigIns() *EnvironmentConfig {
	once.Do(func() {
		instantiated = &EnvironmentConfig{}
		if err := env.Parse(instantiated); err != nil {
			panic(err.Error())
		}
	})
	return instantiated
}
