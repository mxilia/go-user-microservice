package session

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/go-user-microservice.git/internal/entity"
	"github.com/mxilia/go-user-microservice.git/internal/repo"
	"github.com/mxilia/go-user-microservice.git/pkg/redisclient"
)

type RedisSessionRepository struct {
	redis *redisclient.Client
}

func NewRedisSessionRepository(redis *redisclient.Client) repo.SessionRepository {
	return &RedisSessionRepository{
		redis: redis,
	}
}

func (r *RedisSessionRepository) SaveSession(ctx context.Context, session *entity.Session) (*entity.Session, error) {
	panic("unimplemented")
}

func (r *RedisSessionRepository) FindSessionByID(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	panic("unimplemented")
}

func (r *RedisSessionRepository) RevokeSession(ctx context.Context, email string) error {
	panic("unimplemented")
}

func (r *RedisSessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
