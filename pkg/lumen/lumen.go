package lumen

import "errors"

type Error struct {
	appError error
	svcError error
}

func NewError(svcError, appError error) Error {
	return Error{
		appError: appError,
		svcError: svcError,
	}
}

func (e Error) Error() string {
	return errors.Join(e.svcError, e.appError).Error()
}

func (e Error) AppError() error {
	return e.appError
}

func (e Error) SvcError() error {
	return e.svcError
}
