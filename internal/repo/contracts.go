package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entities"
)

type (
	UserRepository interface {
		SaveUser(ctx context.Context, user *entities.User) (*entities.User, error)
		FindAllUsers(ctx context.Context, page int, limit int) ([]*entities.User, int64, error)
		FindUserById(ctx context.Context, id uuid.UUID) (*entities.User, error)
		FindUserByHandler(ctx context.Context, handler string) (*entities.User, error)
		FindUserByEmail(ctx context.Context, email string) (*entities.User, error)
		PatchUser(ctx context.Context, id uuid.UUID, user *entities.User) error
		DeleteUser(ctx context.Context, id uuid.UUID) error
	}

	SessionRepository interface {
		SaveSession(ctx context.Context, session *entities.Session) (*entities.Session, error)
		FindSessionByID(ctx context.Context, id uuid.UUID) (*entities.Session, error)
		RevokeSession(ctx context.Context, email string) error
		DeleteSession(ctx context.Context, id uuid.UUID) error
	}
)
