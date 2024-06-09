package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/db"
	"hotel.com/types"
)

type HotelHandler struct {
	hotelStore db.HotelStore
}

func NewHotelHandler(hotelStore db.HotelStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams types.HotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}

	hotels, err := h.hotelStore.Get(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}
