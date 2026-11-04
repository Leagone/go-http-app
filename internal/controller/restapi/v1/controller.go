package v1

import (
    "go-app/internal/usecase"
    "go-app/pkg/logger"
    "github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
    t  usecase.Translation
    u  usecase.User
    tk usecase.Task
    l  logger.Interface
    v  *validator.Validate
}