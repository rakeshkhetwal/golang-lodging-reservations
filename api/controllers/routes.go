package controllers

import ( "golang-lodging-reservations/api/middlewares"
	    )  

//routes intialize
func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.HttpLogging(s.Home)).Methods("GET")
	
	//logging routes initialization
	standardLogger.SuccessMessage("Routes successfully initialized", "")
}
