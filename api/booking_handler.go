package api

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return err
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return errors.New("something went wrong")
	}

	filter := bson.M{
		"roomID":   rid,
		"canceled": bson.M{"$ne": true},
		"from":     bson.M{"$gte": params.From},
		"till":     bson.M{"$lte": params.Till},
	}
	bookings, err := h.BookingStore.Get(c.Context(), filter)
	if err != nil {
		return err
	}

	if len(bookings) > 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "This room is already booked at the chosen time",
		})
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
		return err
	}

	return c.JSON(res)
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return errors.New("something went wrong")
	}

	filter := bson.M{
		"userID": user.ID,
	}
	bookings, err := h.BookingStore.Get(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(bookings)
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return errors.New("something went wrong")
	}

	booking, err := h.BookingStore.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}

	if booking.UserID != user.ID {
		return c.Status(http.StatusForbidden).JSON("unauthorized")
	}

	if err := h.BookingStore.UpdateByID(
		c.Context(), c.Params("id"), &types.UpdateBookingParams{Canceled: true},
	); err != nil {
		return err
	}

	return c.JSON(map[string]string{
		"message": "booking canceled successfully!",
	})
}
