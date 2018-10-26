package models

import (
	e "errors"
)

var (
	ErrorEntryNotFound = e.New("The entry cannot be found.")
)
