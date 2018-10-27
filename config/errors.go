package config

import (
	e "errors"
)

var (
	ErrorDBHostUnspecified     = e.New("DB Config: Host has not been specified.")
	ErrorDBPortInvalid         = e.New("DB Config: Port is invalid.")
	ErrorDBUserUnspecified     = e.New("DB Config: User has not been specified.")
	ErrorDBPasswordUnspecified = e.New("DB Config: Password has not been specified.")
	ErrorDBNameUnspecified     = e.New("DB Config: DB name has not been specified.")
	ErrorDBSSLModeUnspecified  = e.New("DB Config: SSL mode has not been specified.")

	ErrorAPIPortInvalid     = e.New("API Config: Port is invalid.")
	ErrorAPIPageSizeInvalid = e.New("API Config: Page size is invalid.")
)
