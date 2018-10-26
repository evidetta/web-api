package config

import (
	e "errors"
)

var (
	ErrorHostUnspecified     = e.New("Host has not been specified.")
	ErrorInvalidPort         = e.New("Port is invalid.")
	ErrorUserUnspecified     = e.New("User has not been specified.")
	ErrorPasswordUnspecified = e.New("Password has not been specified.")
	ErrorNameUnspecified     = e.New("DB name has not been specified.")
	ErrorSSLModeUnspecified  = e.New("SSL mode has not been specified.")
)
