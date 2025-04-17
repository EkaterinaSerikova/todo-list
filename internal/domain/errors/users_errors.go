package errors

import "errors"

var ErrEmptyUsersList = errors.New("empty user list")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrInvalidEmail = errors.New("invalid email format")
var ErrWeakPassword = errors.New("password is too weak")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrAccountLocked = errors.New("account locked")
