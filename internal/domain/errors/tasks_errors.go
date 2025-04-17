package errors

import "errors"

var ErrEmptyTasksList = errors.New("empty task list")
var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExist = errors.New("task already exists")

var ErrConflict = errors.New("conflict error")
