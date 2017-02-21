package controllers

import (
	"fmt"
	"net/http"

	model "../models"
)

//Admin GET teste
func Admin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}
		w.Write([]byte(fmt.Sprintf("protected area. USER ID = %v", claims.UserID)))
	}
}
