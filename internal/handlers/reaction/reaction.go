package handlers

import (
	"project/internal/service"
)

type ReactionHandler interface{}

type reactionHandler struct{}

func NewReactionHandler(userService service.UserService) *reactionHandler {
	return &reactionHandler{}
}
