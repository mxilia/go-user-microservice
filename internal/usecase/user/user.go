package user

import (
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

func (u *UserService) CreateUser(user *entities.User) (*entities.User, error) {
	if user == nil {
		return nil, appError.ErrInvalidInput
	}
	return u.userRepository.SaveUser(user)
}

func (u *UserService) FindAllUsers(page int, limit int) ([]*entities.User, int64, error) {
	if page < 1 || limit < 1 {
		return nil, 0, appError.ErrInvalidInput
	}

	users, totalCount, err := u.userRepository.FindAllUsers(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return users, totalCount, err
}

func (u *UserService) FindUserByEmail(email string) (*entities.User, error) {
	panic("unimplemented")
}

func (u *UserService) FindUserByHandler(handler string) (*entities.User, error) {
	panic("unimplemented")
}

func (u *UserService) FindUserById(id uuid.UUID) (*entities.User, error) {
	panic("unimplemented")
}

func (u *UserService) PatchUser(id uuid.UUID, user *entities.User) error {
	panic("unimplemented")
}

func (u *UserService) DeleteUser(id uuid.UUID) error {
	panic("unimplemented")
}
