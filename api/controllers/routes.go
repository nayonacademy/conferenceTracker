package controllers

import "github.com/nayonacademy/conferenceTracker/api/middlewares"

func (server *Server) InitializeRoutes(){
	server.Router.HandleFunc("/status", middlewares.SetMiddlewareJSON(server.HealthCheck)).Methods("GET")
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")
}