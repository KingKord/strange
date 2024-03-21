package handlers

import (
	"log"
	"net/http"
)

type Handlers struct {
}

func NewHandlers() Handlers {
	return Handlers{}
}

func (h Handlers) Root(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
}

func (h Handlers) DaySchedule(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
}

func (h Handlers) AssignMeet(w http.ResponseWriter, r *http.Request) {
	log.Println("got root successfully!")
}
