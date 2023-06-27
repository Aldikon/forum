package service

import (
	"context"

	"github.com/Aldikon/forum/internal/model"
)

func (a *auth) Add(ctx context.Context, r model.Registration) error {
	if err := r.Validate(); err != nil {
		return err
	}

	if err := a.repo.AddUser(ctx, r); err != nil {
		return err
	}
	return nil
}
