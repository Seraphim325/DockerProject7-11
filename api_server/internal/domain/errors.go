package domain

import "errors"

var (
	ErrDuplicate = errors.New("field already exists")
	ErrBigNumber = errors.New("number is too big")
)
