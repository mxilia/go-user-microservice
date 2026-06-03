package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entities"
	"github.com/mxilia/go-user-microservice.git/internal/repo"
	"github.com/mxilia/go-user-microservice.git/internal/usecase"
	appError "github.com/mxilia/go-user-microservice.git/pkg/apperror"
)

type UserService struct {
	userRepository repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) usecase.UserUseCase {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, appError.ErrInvalidInput
	}
	return u.userRepository.SaveUser(ctx, user)
}

func (u *UserService) FindAllUsers(ctx context.Context, page int, limit int) ([]*entities.User, int64, error) {
	if page < 1 || limit < 1 {
		return nil, 0, appError.ErrInvalidInput
	}

	users, totalCount, err := u.userRepository.FindAllUsers(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return users, totalCount, err
}

func (u *UserService) FindUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	if email == "" {
		return nil, appError.ErrInvalidInput
	}
	return u.userRepository.FindUserByEmail(ctx, email)
}

func (u *UserService) FindUserByHandler(ctx context.Context, handler string) (*entities.User, error) {
	if handler == "" {
		return nil, appError.ErrInvalidInput
	}
	return u.userRepository.FindUserByHandler(ctx, handler)
}

func (u *UserService) FindUserById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	if id == uuid.Nil {
		return nil, appError.ErrInvalidInput
	}
	return u.userRepository.FindUserById(ctx, id)
}

func (u *UserService) PatchUser(ctx context.Context, id uuid.UUID, user *entities.User) error {
	if id == uuid.Nil || user == nil {
		return appError.ErrInvalidInput
	}
	return u.userRepository.PatchUser(ctx, id, user)
}

func (u *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return appError.ErrInvalidInput
	}
	return u.userRepository.DeleteUser(ctx, id)
}
