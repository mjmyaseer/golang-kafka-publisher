package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	 "Company import path goes here"

	"net/http"

)

func RouterStart() {

	var publisher publisher_service.Publisher
	publisher = publisher_service.NewEventPublisher()
	publisher = instrumenting.TripLifeCycle{Publisher: publisher}

	r := mux.NewRouter()

	eventHandler := httpTransport.NewServer(
		endpoints.MakeEventEndpoint(publisher),
		request.DecodeEventRequest,
		response.EncodeResponse,
		httpTransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			var Error struct {
				Error struct {
					Code     int    `json:"code"`
					Messsage string `json:"messsage"`
				} `json:"error"`
			}
			Error.Error.Code = 1200
			Error.Error.Messsage = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Error)
		}),
	)
	port := config.AppConf.Port
	r.Handle("/publish", eventHandler).Methods(http.MethodPost)
	r.Handle("/metrics", promhttp.Handler())
	log.Info("Server Initialization Started...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, port), r))
}
