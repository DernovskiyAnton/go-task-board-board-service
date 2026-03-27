package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/DernovskiyAnton/go-task-board-board-service/internal/domain"
	"github.com/google/uuid"
)

type boardUsecase struct {
	boardRepository BoardRepository
}

func (b *boardUsecase) Create(ctx context.Context, name, description, ownerID string) (domain.Board, error) {
	if name == "" {
		return domain.Board{}, errors.New("name is required")
	}
	board := domain.Board{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := b.boardRepository.Create(ctx, board)
	if err != nil {
		return domain.Board{}, err
	}

	return board, nil
}

func (b *boardUsecase) GetByID(ctx context.Context, id string) (domain.Board, error) {
	if id == "" {
		return domain.Board{}, errors.New("id is required")
	}

	board, err := b.boardRepository.GetByID(ctx, id)
	if err != nil {
		return domain.Board{}, err
	}
	return board, nil
}

func (b *boardUsecase) GetByOwnerID(ctx context.Context, ownerID string) ([]domain.Board, error) {
	if ownerID == "" {
		return nil, errors.New("ownerID is required")
	}
	boards, err := b.boardRepository.GetByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (b *boardUsecase) Update(ctx context.Context, board domain.Board) (domain.Board, error) {
	if board.ID == "" {
		return domain.Board{}, errors.New("board is required")
	}
	board.UpdatedAt = time.Now()
	board, err := b.boardRepository.Update(ctx, board)
	if err != nil {
		return domain.Board{}, err
	}
	return board, nil
}

func (b *boardUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return b.boardRepository.Delete(ctx, id)
}

func NewBoardUsecase(boardRepository BoardRepository) BoardUsecase {
	return &boardUsecase{
		boardRepository: boardRepository,
	}
}
