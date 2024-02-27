package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kudarap/foo"
)

type service interface {
	FighterByID(ctx context.Context, id string) (*foo.Fighter, error)
}

func GetFighterByID(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		c, err := s.FighterByID(r.Context(), id)
		if err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		encodeJSONResp(w, c, http.StatusOK)
	}
}

func ListFighters(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		encodeJSONResp(w, struct {
			Msg string `json:"message"`
		}{"no fighters yet implemented"}, http.StatusOK)
	}
}
