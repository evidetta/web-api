package config

import (
	"fmt"
	"strconv"
)

type APIConfig struct {
	Port     int
	PageSize int
}

func NewAPIConfig(port, pageSize string) (*APIConfig, error) {
	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, ErrorAPIPortInvalid
	}

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, ErrorAPIPageSizeInvalid
	}

	api_conf := APIConfig{
		Port:     p,
		PageSize: ps,
	}

	return &api_conf, nil
}

func (api_conf *APIConfig) GetAPIAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", api_conf.Port)
}
