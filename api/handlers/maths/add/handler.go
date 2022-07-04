package add

import (
	"encoding/json"
	"net/http"

	"github.com/PeteXC/maths-api/api/respond"
	"github.com/joerdav/zapray"
)

type Handler struct {
	Log *zapray.Logger
}

type Input struct {
	NumberA int `json:"numberA"`
	NumberB int `json:"numberB"`
}

type Output struct {
	Result int `json:"result"`
}

func (h Handler) Handle(a, b int) (x int) {
	return a + b
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Contains code to handle request
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.Log.Error("Error decoding request body")
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}

	x := h.Handle(input.NumberA, input.NumberB)
	result := Output{
		Result: x,
	}
	respond.WithJSON(w, result)
}

func NewHandler(log *zapray.Logger) (h Handler, err error) {
	h = Handler{
		Log: log,
	}
	return
}
