package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/samarthasthan/go-hotel/db"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		cxt = context.Background()
	)
	user, err := h.userStore.GetUserByID(cxt, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	users, err := h.userStore.GetUsers(ctx)
	if err != nil {
		return err
	}
	return c.JSON(users)
}
