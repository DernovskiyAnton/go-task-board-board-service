package usecase

import (
	"context"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
)

type TaskUsecase interface {
	Create(ctx context.Context, columnID, header string) (domain.Task, error)

	GetByID(ctx context.Context, id string) (domain.Task, error)

	Update(ctx context.Context, task domain.Task) (domain.Task, error)
	Delete(ctx context.Context, id string) error
	Move(ctx context.Context, taskID, newColumnID string, newPosition int64) error
}
