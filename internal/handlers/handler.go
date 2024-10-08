package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/KingKord/strange/internal/helpers"
	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/services"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
// @Success 200 {object} string
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
// @Param user-id query string true "user ID"
// @Success 200 {object} string "Successful assignment"
// @Failure 400 {object} string "Invalid date format"
// @Failure 500 {object} string "Internal server error"
// @Router /schedule/day [get]
func (h Handlers) DaySchedule(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	userID := r.URL.Query().Get("user-id")
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
	parsedUserID, err := strconv.Atoi(userID)
	if err != nil {
		_ = helpers.ErrorJSON(w, fmt.Errorf("invalid int format: %v", err), http.StatusBadRequest)
		return
	}
	// Ваш код для обработки запроса с использованием даты
	daySchedule, err := h.scheduleService.DaySchedule(context.Background(), parsedDate, parsedUserID)
	if err != nil {
		_ = helpers.ErrorJSON(w, fmt.Errorf("scheduleService.DaySchedule: %v", err), http.StatusInternalServerError)
		return
	}
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
// @Success 200 {object} string "Successful assignment"
// @Failure 500 {object} string "Internal server error"
// @Router /schedule/reserve [post]
func (h Handlers) AssignMeet(w http.ResponseWriter, r *http.Request) {
	req := requestPayload{}
	err := helpers.ReadJSON(w, r, &req)
	if err != nil {
		_ = helpers.ErrorJSON(w, fmt.Errorf("helpers.ReadJSON: %w", err), http.StatusBadRequest)
		return
	}

	err = validateAssignMeetRequest(req)
	if err != nil {
		_ = helpers.ErrorJSON(w, fmt.Errorf("validateAssignMeetRequest: %w", err), http.StatusBadRequest)
		return
	}

	err = h.scheduleService.AssignMeet(context.Background(), model.Card{
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

func validateAssignMeetRequest(req requestPayload) error {
	if req.DateFrom.After(req.DateTo) {
		return fmt.Errorf("invalid date interval: from-%v to-%v", req.DateFrom, req.DateTo)
	}

	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 1000)),
		validation.Field(&req.Description, validation.Length(0, 1000)),
	)
}
