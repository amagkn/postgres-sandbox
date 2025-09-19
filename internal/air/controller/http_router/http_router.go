package http_router

import (
	"github.com/amagkn/postgres-sandbox/internal/air/controller/http_router/ver1"
	"github.com/amagkn/postgres-sandbox/internal/air/usecase"
	"github.com/go-chi/chi/v5"
)

func AirRoutes(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Route("/api/v1/air", func(r chi.Router) {
		r.Get("/passenger/names", v1.PassengerNames)
		r.Get("/passenger/all", v1.PassengerAll)
		r.Get("/trip/all", v1.TripAll)
	})
}
