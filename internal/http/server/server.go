package server

import (
	"github.com/DemianShtepa/factorial-test/internal/http/handler"
	"github.com/julienschmidt/httprouter"
)

func InitializeServer() *httprouter.Router {
	router := httprouter.New()

	router.POST("/calculate", handler.CalculateHandler())

	return router
}
