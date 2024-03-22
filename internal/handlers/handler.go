package handlers

import (
	"context"
	"errors"
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

type requestPayload struct {
	Name        string    `json:"name"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	DateFrom    time.Time `json:"date_from"`
	DateTo      time.Time `json:"date_to"`
}

// DaySchedule godoc
// @Summary Get meets for day
// @Description Get meets for day
// @Tags Schedule
// @Accept json
// @Produce json
// @Param date query string true "Date of the day (YYYY-MM-DD)"
// @Success 200 {object} jsonResponse "Successful assignment"
// @Failure 400 {object} jsonResponse "Invalid date format"
// @Failure 500 {object} jsonResponse "Internal server error"
// @Router /schedule/day [get]
func (h Handlers) DaySchedule(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		_ = helpers.ErrorJSON(w, errors.New("date parameter is required"), http.StatusBadRequest)
		return
	}

	// Парсинг даты
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		_ = helpers.ErrorJSON(w, fmt.Errorf("invalid date format: %v", err), http.StatusBadRequest)
		return
	}

	// Ваш код для обработки запроса с использованием даты
	daySchedule, err := h.scheduleService.DaySchedule(context.Background(), parsedDate)
	if err != nil {
		return
	}
	// Пример:
	_ = helpers.WriteJSON(w, http.StatusOK, jsonResponse{
		Message: fmt.Sprintf("Successfully got meets for %s", parsedDate.Format("2006-01-02")),
		Data:    daySchedule,
	})
}

// AssignMeet godoc
// @Summary Assign meet for the day
// @Description Assign meet for the day
// @Tags Schedule
// @Accept json
// @Produce json
// @Param requestPayload body requestPayload true "Request payload"
// @Success 200 {object} jsonResponse "Successful assignment"
// @Failure 500 {object} jsonResponse "Internal server error"
// @Router /schedule/reserve [post]
func (h Handlers) AssignMeet(w http.ResponseWriter, r *http.Request) {
	req := requestPayload{}

	_ = helpers.ReadJSON(w, r, &req)
	err := h.scheduleService.AssignMeet(context.Background(), model.Card{
		Name:        req.Name,
		UserID:      req.UserID,
		Description: req.Description,
		From:        req.DateFrom,
		To:          req.DateTo,
	})
	if err != nil {
		log.Println(err)
		_ = helpers.ErrorJSON(w, fmt.Errorf("scheduleService.AssignMeet: %w", err), http.StatusInternalServerError)
	}
	log.Println("success assigned!")

	_ = helpers.WriteJSON(w, http.StatusOK, jsonResponse{
		Message: "successfully assigned meet!",
	})
}
