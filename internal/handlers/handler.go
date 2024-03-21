package handlers

import (
	"github.com/KingKord/strange/internal/helpers"
	"log"
	"net/http"
)

type Handlers struct {
}

func NewHandlers() Handlers {
	return Handlers{}
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
	log.Println("got root successfully!")
}

func (h Handlers) AssignMeet(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
}
