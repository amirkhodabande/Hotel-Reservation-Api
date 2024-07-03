package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/api/custom_errors"
	"hotel.com/db"
	"hotel.com/types"
)

type BookingHandler struct {
	*db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store,
	}
}

func (h *BookingHandler) HandleBookRoom(c *fiber.Ctx) error {
	params := c.Context().UserValue("params").(*types.BookRoomParams)

	rid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return custom_errors.NotFound()
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return custom_errors.Internal()
	}

	filter := &types.BookingQueryParams{
		RoomID:   rid,
		Canceled: false,
		From:     params.From,
		Till:     params.Till,
	}

	bookings, err := h.BookingStore.Get(c.Context(), filter)
	if err != nil {
		return custom_errors.Internal()
	}

	if len(bookings) > 0 {
		return custom_errors.NewErr(
			http.StatusBadRequest,
			"This room is already booked at the chosen time",
		)
	}

	booking := &types.Booking{
		UserID:     user.ID,
		RoomID:     rid,
		From:       params.From,
		Till:       params.Till,
		NumPersons: params.NumPersons,
	}

	res, err := h.BookingStore.Insert(c.Context(), booking)
	if err != nil {
		return custom_errors.Internal()
	}

	return c.Status(http.StatusCreated).JSON(SuccessResponse(res))
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return custom_errors.Internal()
	}

	params := c.Context().UserValue("query-params").(*types.BookingQueryParams)

	params.UserID = user.ID

	bookings, err := h.BookingStore.Get(c.Context(), params)
	if err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(
		SuccessResponse(bookings).WithPagination(int64(len(bookings)), params.GetPage()),
	)
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return custom_errors.Internal()
	}

	booking, err := h.BookingStore.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return custom_errors.Internal()
	}

	if booking.UserID != user.ID {
		return custom_errors.Forbidden()
	}

	if err := h.BookingStore.UpdateByID(
		c.Context(), c.Params("id"), &types.UpdateBookingParams{Canceled: true},
	); err != nil {
		return custom_errors.Internal()
	}

	return c.JSON(SuccessResponse(nil))
}
