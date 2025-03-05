package service

import (
	"errors"
	// "github.com/tsubasa66739/gin-nextjs-webapp/dto"
	// "github.com/tsubasa66739/gin-nextjs-webapp/repository"
	// "github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
)

var (
	ErrNotFound       = errors.New("resource not found")
	ErrInternalServer = errors.New("unknown error")
)
