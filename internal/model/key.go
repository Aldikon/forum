package model

import "context"

type contextKey int

const UserID contextKey = iota

func GetID(ctx context.Context) int64 {
	id, ok := ctx.Value(UserID).(int64)
	if !ok {
		return 0
	}

	return id
}
