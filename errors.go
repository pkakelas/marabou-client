package main

import (
	"errors"
)

var InternalError = errors.New("Internal Error")
var IncopatibleVersions = errors.New("Incopatible Versions")
var InvalidMessage = errors.New("Invalid Message Type")
var InvalidInputError = errors.New("Invalid Input")
