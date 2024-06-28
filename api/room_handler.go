package api

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/db"
)

type RoomHandler struct {
	*db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store,
	}
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	hid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	rooms, err := h.RoomStore.Get(c.Context(), bson.M{"hotelID": hid})
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}
