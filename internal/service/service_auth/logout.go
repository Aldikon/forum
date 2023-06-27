package service

import "context"

func (a *auth) LogOut(ctx context.Context, userID int64) error {
	return a.repo.DeleteSession(ctx, userID)
}
