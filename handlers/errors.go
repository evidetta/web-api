package handlers

import (
	e "errors"
)

var (
	ErrorInvalidParameter    = e.New("The parameter was invalid.")
	ErrorInvalidPayload      = e.New("The payload was invalid.")
	ErrorInternalServerError = e.New("There was an internal server error.")
	ErrorObjectNotFound      = e.New("The object cannot be found.")
)
