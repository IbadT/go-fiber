package service

import "errors"

// ErrInvalidCredentials возвращается при неверных учетных данных
var ErrInvalidCredentials = errors.New("invalid email or password")

// ErrNotFound возвращается, когда запрашиваемый ресурс не найден
var ErrNotFound = errors.New("resource not found")

// ErrUnauthorized возвращается, когда у пользователя нет прав для выполнения действия
var ErrUnauthorized = errors.New("unauthorized access")

// ErrValidation возвращается при ошибке валидации данных
var ErrValidation = errors.New("validation error")
