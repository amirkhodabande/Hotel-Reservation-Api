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
	params := c.Context().UserValue("query-params").(*types.HotelQueryParams)

	hotels, err := h.HotelStore.Get(c.Context(), params)
	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(SuccessResponse(hotels))
}
