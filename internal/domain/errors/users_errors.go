package errors

import "errors"

var ErrEmptyUsersList = errors.New("empty user list")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")
