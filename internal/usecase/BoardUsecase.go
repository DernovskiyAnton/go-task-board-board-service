package usecase

import (
	"context"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
)

type BoardUsecase interface {
	Create(ctx context.Context, name, description, ownerID string) (domain.Board, error)

	GetByID(ctx context.Context, id string) (domain.Board, error)

	GetByOwnerID(ctx context.Context, ownerID string) ([]domain.Board, error)

	Update(ctx context.Context, board domain.Board) (domain.Board, error)

	Delete(ctx context.Context, id string) error
}
