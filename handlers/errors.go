package handlers

import (
	e "errors"
)

var (
	ErrorNoTagFound          = e.New("The message did not contain a tag.")
	ErrorInvalidPayload      = e.New("The payload was invalid.")
	ErrorInternalServerError = e.New("There was an internal server error.")
	ErrorObjectNotFound      = e.New("The object cannot be found.")
	ErrorInvalidEndpoint     = e.New("The endpoint is invalid.")
)
