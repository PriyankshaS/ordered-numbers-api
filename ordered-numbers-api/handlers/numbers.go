package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"ordered-numbers-api/services"
)

type NumbersHandler struct{
	numbersService *services.NumbersService
}

func NewNumbersHandler(apiToken, apiURL string) *NumbersHandler {
	numbersService := services.NewNumbersService(apiURL,apiToken)
	return &NumbersHandler{
		numbersService : numbersService,
	}
}

func (h *NumbersHandler) GetOrderedNo(w http.ResponseWriter, r *http.Request){
	numbers, err := h.numbersService.FetchOrderedNo()
	if(err != nil){
		log.Fatal("not able to fetch no.s", err)
		http.Error(w, "Failed to fetch no.s", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}
