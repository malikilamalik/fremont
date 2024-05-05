package lumen

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")      //400
	ErrUnauthorized    = errors.New("unauthorized")     //401
	ErrNotFound        = errors.New("not found")        //404
	ErrConflict        = errors.New("conflict")         //409
	ErrInternalFailure = errors.New("internal failure") //500
)
