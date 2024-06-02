package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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

	return c.Status(fiber.StatusCreated).JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(404).JSON(map[string]string{"error": "not found"})
		}

		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var params types.UpdateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := h.userStore.UpdateUserByID(c.Context(), c.Params("id"), params); err != nil {
		return err
	}

	return c.JSON(map[string]string{"message": "user updated successfully"})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userStore.DeleteUserByID(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"message": "user deleted successfully"})
}
