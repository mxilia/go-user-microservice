package session

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entities"
	"github.com/mxilia/go-user-microservice.git/internal/repo"
	"github.com/mxilia/go-user-microservice.git/internal/usecase"
	appError "github.com/mxilia/go-user-microservice.git/pkg/apperror"
)

type SessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) usecase.SessionUseCase {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	if session == nil {
		return nil, appError.ErrInvalidInput
	}
	return s.sessionRepo.SaveSession(ctx, session)
}

func (s *SessionService) FindSessionByID(ctx context.Context, id uuid.UUID) (*entities.Session, error) {
	if id == uuid.Nil {
		return nil, appError.ErrInvalidInput
	}
	return s.sessionRepo.FindSessionByID(ctx, id)
}

func (s *SessionService) RevokeSession(ctx context.Context, email string) error {
	if email == "" {
		return appError.ErrInvalidInput
	}
	return s.sessionRepo.RevokeSession(ctx, email)
}

func (s *SessionService) DeleteSession(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return appError.ErrInvalidInput
	}
	return s.sessionRepo.DeleteSession(ctx, id)
}
