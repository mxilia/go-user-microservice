package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entity"
)

type (
	UserUseCase interface {
		CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
		FindAllUsers(ctx context.Context, page int, limit int) ([]*entity.User, int64, error)
		FindUserById(ctx context.Context, id uuid.UUID) (*entity.User, error)
		FindUserByHandler(ctx context.Context, handler string) (*entity.User, error)
		FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
		PatchUser(ctx context.Context, id uuid.UUID, user *entity.User) error
		DeleteUser(ctx context.Context, id uuid.UUID) error
	}

	SessionUseCase interface {
		CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error)
		FindSessionByID(ctx context.Context, id uuid.UUID) (*entity.Session, error)
		RevokeSession(ctx context.Context, email string) error
		DeleteSession(ctx context.Context, id uuid.UUID) error
	}
)
