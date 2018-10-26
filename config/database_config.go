package config

import (
	"fmt"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewDatabaseConfig(host, port, user, password, dbname, sslmode string) (*DatabaseConfig, error) {

	if len(host) == 0 {
		return nil, ErrorHostUnspecified
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, ErrorInvalidPort
	}

	if len(user) == 0 {
		return nil, ErrorUserUnspecified
	}

	if len(password) == 0 {
		return nil, ErrorPasswordUnspecified
	}

	if len(dbname) == 0 {
		return nil, ErrorNameUnspecified
	}

	if len(sslmode) == 0 {
		return nil, ErrorSSLModeUnspecified
	}

	db_conf := &DatabaseConfig{
		Host:     host,
		Port:     p,
		User:     user,
		Password: password,
		Name:     dbname,
		SSLMode:  sslmode,
	}

	return db_conf, nil
}

func (db_conf *DatabaseConfig) GetConnectionString() string {
	conf_str := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		db_conf.Host,
		db_conf.Port,
		db_conf.User,
		db_conf.Password,
		db_conf.Name,
		db_conf.SSLMode,
	)

	return conf_str
}
