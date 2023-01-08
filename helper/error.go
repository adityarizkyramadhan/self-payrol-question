package helper

import "errors"

var (
	ErrWrongId = errors.New("wrong id")
	ErrBalance = errors.New("balance is not enough")
	ErrAmount  = errors.New("amount cannot be negative")
)
