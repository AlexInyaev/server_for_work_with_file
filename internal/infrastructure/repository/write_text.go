package repository

import (
	"context"
	"fmt"
)

func (r *Repository) WriteTextToFile(ctx context.Context, name, email, text string) error {
	_ = ctx
	fmt.Printf("name: %s \n email %s \n text %s \n", name, email, text)
	return nil
}
