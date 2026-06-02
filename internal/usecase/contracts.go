package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entities"
)

type (
	UserUseCase interface {
		CreateUser(user *entities.User) (*entities.User, error)
		FindAllUsers(page int, limit int) ([]*entities.User, int64, error)
		FindUserById(id uuid.UUID) (*entities.User, error)
		FindUserByHandler(handler string) (*entities.User, error)
		FindUserByEmail(email string) (*entities.User, error)
		PatchUser(id uuid.UUID, user *entities.User) error
		DeleteUser(id uuid.UUID) error
	}

	SessionUseCase interface {
		CreateSession() (*entities.Session, error)
		FindSessionByID(id uuid.UUID) (*entities.Session, error)
		RevokeSession(email string) error
		DeleteSession(id uuid.UUID) error
	}
)
