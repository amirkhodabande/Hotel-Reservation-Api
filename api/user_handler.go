package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/api/custom_errors"
	"hotel.com/db"
	"hotel.com/types"
)

type UserHandler struct {
	*db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		store,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.Get(c.Context())
	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	params := c.Context().UserValue("params").(*types.CreateUserParams)

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return custom_errors.Internal()
	}

	insertedUser, err := h.UserStore.Insert(c.Context(), user)

	if err != nil {
		return custom_errors.Internal()
	}

	return c.Status(fiber.StatusCreated).JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.UserStore.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return custom_errors.NotFound()
		}

		return custom_errors.Internal()
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	params := c.Context().UserValue("params").(*types.UpdateUserParams)

	if err := h.UserStore.UpdateByID(c.Context(), c.Params("id"), params); err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(map[string]string{"message": "user updated successfully"})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.UserStore.DeleteByID(c.Context(), id)

	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(map[string]string{"message": "user deleted successfully"})
}
