package service

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrInternalServer = errors.New("unknown error")
)
