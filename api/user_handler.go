package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/db"
	"hotel.com/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)

	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)

	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userStore.DeleteUserByID(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"message": "user deleted successfully"})
}
