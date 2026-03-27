package usecase

import (
	"context"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
)

type ColumnRepository interface {
	Create(ctx context.Context, column domain.Column) (domain.Column, error)

	GetByID(ctx context.Context, id string) (domain.Column, error)

	GetByBoardID(ctx context.Context, boardID string) ([]domain.Column, error)

	Update(ctx context.Context, column domain.Column) (domain.Column, error)

	Delete(ctx context.Context, id string) error
}
