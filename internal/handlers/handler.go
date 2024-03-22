package handlers

import (
	"context"
	"fmt"
	"github.com/KingKord/strange/internal/helpers"
	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/services"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	scheduleService services.ScheduleService
}

func NewHandlers(scheduleService services.ScheduleService) Handlers {
	return Handlers{
		scheduleService: scheduleService,
	}
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Root is an API endpoint that register a user and returns a JSON response.
// @Summary Get root endpoint
// @Description Get the root endpoint
// @Tags Root
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse
// @Router / [get]
func (h Handlers) Root(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
	_ = helpers.WriteJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "successfully got Root",
	})
}

func (h Handlers) DaySchedule(w http.ResponseWriter, r *http.Request) {
	requestPayload := struct {
		Name        string    `json:"name"`
		UserID      int       `json:"user_id"`
		Description string    `json:"description"`
		DateFrom    time.Time `json:"date_from"`
		DateTo      time.Time `json:"date_to"`
	}{}

	_ = helpers.ReadJSON(w, r, &requestPayload)
	err := h.scheduleService.AssignMeet(context.Background(), model.Card{})
	if err != nil {
		log.Println(err)
		_ = helpers.ErrorJSON(w, fmt.Errorf("scheduleService.AssignMeet: %w", err), http.StatusInternalServerError)
	}
	log.Println("success assigned!")

	_ = helpers.WriteJSON(w, http.StatusOK, jsonResponse{
		Message: "successfully assigned meet!",
	})
}

func (h Handlers) AssignMeet(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
}
