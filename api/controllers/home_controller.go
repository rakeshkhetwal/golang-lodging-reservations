package controllers

import (
	"net/http"
	
	"golang-lodging-reservations/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Lodging Reservation Portal")
}
