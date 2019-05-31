package events

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// EventsServiceHandler defines object model
type EventsServiceHandler struct {
	Message string `json:"message,omitempty"`
}

// AllEventHandler returns all events
func (eh *EventsServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := EventsServiceHandler{Message: "alleventhandler"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

// FindEventHandler will find events
func (eh *EventsServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := EventsServiceHandler{Message: "findeventhandler"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

// AddEventHandler will add new events. @// TODO: move to POST
func (eh *EventsServiceHandler) AddEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := EventsServiceHandler{Message: "addeventhandler"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

// Index serves default.
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Handlers handles events related to events
func Handlers() *mux.Router {
	eh := EventsServiceHandler{}
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	eventsrouter := router.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("").HandlerFunc(eh.AllEventHandler)
	eventsrouter.Methods("GET").Path("/findevents").HandlerFunc(eh.FindEventHandler)
	eventsrouter.Methods("GET").Path("/addevent").HandlerFunc(eh.AddEventHandler)
	return router
}
