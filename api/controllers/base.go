package controllers

import (

	"net/http"
	"errors"

	log "golang-lodging-reservations/api/logger"
	"github.com/gorilla/mux"
	
)

type Server struct {
	Router *mux.Router
}	

var server = Server{}
var standardLogger = log.Logger()
//null error handler message
var nilerr = errors.New("")


func Run() {
    // Router  initialize
    server.Router = mux.NewRouter()
	server.initializeRoutes()
	// Server Listen
	server.RunServer(":8082")

}

func (server *Server) RunServer(addr string) {
	standardLogger.SuccessMessage("Starting server at", "8082")
	standardLogger.FatalErrorMessage(http.ListenAndServe(addr, server.Router),"")
}
