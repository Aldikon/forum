package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"forum/internal/model"

	"github.com/gofrs/uuid"
)

func (a *auth) LogIn(ctx context.Context, user model.LogIn) (model.Session, error) {
	var session model.Session

	id, err := a.repo.LogIn(ctx, user)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return session, model.ErrNotFoundUser
	case err != nil:
		return session, err
	case id < 0:
		return session, model.ErrNotFoundUser
	}

	token, err := uuid.NewV4()
	if err != nil {
		return session, err
	}

	session.UserID = id
	session.Token = token.String()
	session.EndAtt = time.Now().Add(time.Hour)

	if err := a.repo.AddSession(ctx, session); err != nil {
		return session, err
	}

	return session, nil
}
