package config

type Config struct {
	Mode ModeName // 环境
}

type ModeName string

func DefaultModeName() ModeName {
	return "test"
}

func DefaultConfig(mode ModeName) *Config {
	return &Config{Mode: mode}
}
