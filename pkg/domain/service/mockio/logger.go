// Code generated by MockIO. DO NOT EDIT.
// Source: logger.go

// Package mockio_service is a generated GoMock package.
package mockio_service

import (
	context "context"
)

type LoggerError struct {
	Time   int
	ArgCtx context.Context
	ArgErr error
}

type LoggerSetUser struct {
	Time        int
	ArgCtx      context.Context
	ArgUserID   int64
	ArgUserName string
	Ret0        context.Context
}