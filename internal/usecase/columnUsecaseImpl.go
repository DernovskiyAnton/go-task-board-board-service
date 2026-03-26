package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
	"github.com/google/uuid"
)

type columnUsecase struct {
	columnRepository ColumnRepository
}

func (c *columnUsecase) Create(ctx context.Context, name, boardID string, position int64) (domain.Column, error) {
	if name == "" || boardID == "" {
		return domain.Column{}, errors.New("name or boardID is empty")
	}
	column := domain.Column{
		ID:        uuid.New().String(),
		Name:      name,
		BoardID:   boardID,
		Position:  position,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := c.columnRepository.Create(ctx, column)
	if err != nil {
		return domain.Column{}, err
	}

	return column, nil
}

func (c *columnUsecase) GetByID(ctx context.Context, id string) (domain.Column, error) {
	if id == "" {
		return domain.Column{}, errors.New("id is empty")
	}
	columns, err := c.columnRepository.GetByID(ctx, id)
	if err != nil {
		return domain.Column{}, err
	}

	return columns, nil
}

func (c *columnUsecase) GetByBoardID(ctx context.Context, boardID string) ([]domain.Column, error) {
	if boardID == "" {
		return nil, errors.New("boardID is empty")
	}
	columns, err := c.columnRepository.GetByBoardID(ctx, boardID)
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func (c *columnUsecase) Update(ctx context.Context, column domain.Column) (domain.Column, error) {
	if column.ID == "" {
		return domain.Column{}, errors.New("id is empty")
	}
	column.UpdatedAt = time.Now()
	updatedColumn, err := c.columnRepository.Update(ctx, column)
	if err != nil {
		return domain.Column{}, err
	}
	return updatedColumn, nil
}

func (c *columnUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	return c.columnRepository.Delete(ctx, id)
}

func NewColumnUsecase(columnRepository ColumnRepository) ColumnUsecase {
	return &columnUsecase{
		columnRepository: columnRepository,
	}
}
