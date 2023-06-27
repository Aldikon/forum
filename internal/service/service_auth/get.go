package service

import "context"

func (a *auth) GetID(ctx context.Context, token string) (int64, error) {
	session, err := a.repo.GetID(ctx, token)
	if err != nil {
		return 0, err
	}

	if err := session.Validate(); err != nil {
		return 0, err
	}

	return session.UserID, nil
}

func (a *auth) GetByID(ctx context.Context, userID int64) (string, error) {
	if userID <= 0 {
		return "", nil
	}
	return a.repo.GetByID(ctx, userID)
}
