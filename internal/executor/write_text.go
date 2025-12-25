package executor

import (
	"context"
)

type WriteTextRepository interface {
	WriteTextToFile(ctx context.Context, name, email, text string) error
}

type WriteText struct {
	repo WriteTextRepository
}

func NewWriteText(repo WriteTextRepository) *WriteText {
	return &WriteText{
		repo: repo,
	}
}

func (e *WriteText) Execute(ctx context.Context, name, email, text string) error {
	err := e.repo.WriteTextToFile(ctx, name, email, text)
	if err != nil {
		return err
	}

	return nil
}
