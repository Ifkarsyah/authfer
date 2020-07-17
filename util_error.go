package main

import "errors"

var (
	ErrAuth       = errors.New("err auth")
	ErrBadRequest = errors.New("err bad request")
)
