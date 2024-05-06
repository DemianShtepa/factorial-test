package handler

import (
	"encoding/json"
	"github.com/DemianShtepa/factorial-test/internal/factorial"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync"
)

type requestBody struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
}

type responseBody struct {
	A uint64 `json:"a!"`
	B uint64 `json:"b!"`
}

type errorResponseBody struct {
	Err string `json:"error"`
}

func CalculateHandler() func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	return validateBody()
}

func validateBody() func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var body requestBody

		writer.Header().Add("Content-Type", "application/json")

		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			response := errorResponseBody{Err: "Incorrect input"}
			if err = json.NewEncoder(writer).Encode(response); err != nil {
				log.Fatalln(err)
			}

			return
		}

		handler := handle(body)
		handler(writer, request, params)
	}
}

func handle(body requestBody) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var aFactorial, bFactorial uint64
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			aFactorial = factorial.Calculate(body.A)
		}()
		go func() {
			defer wg.Done()

			bFactorial = factorial.Calculate(body.B)
		}()

		wg.Wait()

		response := responseBody{
			A: aFactorial,
			B: bFactorial,
		}

		if err := json.NewEncoder(writer).Encode(response); err != nil {
			log.Fatalln(err)
		}
	}
}
