package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/api/custom_errors"
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
	// TODO: validate this and use params to filter
	var qparams types.HotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return custom_errors.Validation()
	}

	hotels, err := h.HotelStore.Get(c.Context(), nil)
	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(hotels)
}
