package errors

import "errors"

var ErrEmptyTasksList = errors.New("empty task list")
var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExist = errors.New("task already exists")

var ErrEmptyUsersList = errors.New("empty user list")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")
