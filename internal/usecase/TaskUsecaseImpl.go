package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
	"github.com/google/uuid"
)

type taskUsecase struct {
	taskRepository TaskRepository
}

func (t *taskUsecase) Create(ctx context.Context, columnID, header string) (domain.Task, error) {
	if columnID == "" || header == "" {
		return domain.Task{}, errors.New("columnID or header is required")
	}

	task := domain.Task{
		ID:        uuid.New().String(),
		ColumnID:  columnID,
		Header:    header,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := t.taskRepository.Create(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *taskUsecase) GetByID(ctx context.Context, id string) (domain.Task, error) {
	if id == "" {
		return domain.Task{}, errors.New("id is required")
	}
	task, err := t.taskRepository.GetByID(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *taskUsecase) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	if task.ID == "" {
		return domain.Task{}, errors.New("id is required")
	}
	task.UpdatedAt = time.Now()
	_, err := t.taskRepository.Update(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *taskUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return t.taskRepository.Delete(ctx, id)
}

func (t *taskUsecase) Move(ctx context.Context, taskID, newColumnID string, newPosition int64) error {
	if taskID == "" || newColumnID == "" {
		return errors.New("taskID or newColumnID is required")
	}

	task, err := t.taskRepository.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	task.ColumnID = newColumnID
	task.Position = newPosition
	_, err = t.taskRepository.Update(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func NewTaskUsecase(taskRepository TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}
