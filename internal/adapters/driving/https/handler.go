package https

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Resul-Necefli/gophersub/internal/core/ports"
)

// SubscriptionHandler handles HTTP requests related to subscriptions.

type SubscriptionHandler struct {
	servs ports.Subscribe
}

func NewSubscriptionHandler(serv ports.Subscribe) *SubscriptionHandler {
	return &SubscriptionHandler{
		servs: serv,
	}
}

func (s *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Println("method not supported")
		http.Error(w, "method not supported", http.StatusBadRequest)
		return
	}

	var dto = SubscribeRequest{}

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		log.Println("error decoding request body:", err)
		http.Error(w, "error decoding request body: ", http.StatusBadRequest)
		return
	}

	err = s.servs.Subscribe(dto.UserID, dto.PlanName, dto.Amount, dto.Currency)
	if err != nil {
		log.Println("subscribe error:", err)
		http.Error(w, "subscribe error: ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("subscription created successfully"))

}
