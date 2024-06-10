package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/db"
	"hotel.com/types"
)

type HotelHandler struct {
	*db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams types.HotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}

	hotels, err := h.HotelStore.Get(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}
