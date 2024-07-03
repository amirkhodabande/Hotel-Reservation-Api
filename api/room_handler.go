package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/api/custom_errors"
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
	rooms, err := h.RoomStore.GetByHotelID(c.Context(), c.Params("id"))
	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(rooms)
}
