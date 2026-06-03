package persistence

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mxilia/go-user-microservice.git/pkg/apperror"
	"github.com/redis/go-redis/v9"
)

func translatePostgresError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// unique_violation
		case "23505":
			return apperror.ErrAlreadyExists
		// foreign_key_violation
		case "23503":
			return apperror.ErrConflict
		// check_violation
		case "23514":
			return apperror.ErrValidation
		// not_null_violation
		case "23502":
			return apperror.ErrRequiredField
		// string_data_right_truncation
		case "22001":
			return apperror.ErrOutOfRange
		}
	}
	return err
}

func translateRedisError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, redis.Nil):
		return apperror.ErrNotFound
	case errors.Is(err, context.DeadlineExceeded):
		return apperror.ErrTimeout
	case errors.Is(err, context.Canceled):
		return apperror.ErrTimeout
	}
	return apperror.ErrExternalService
}
