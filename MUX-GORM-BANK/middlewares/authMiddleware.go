package middlewares

import (
	"bankingApp/models"
	"net/http"
)

var secretKey = []byte("secret_key")

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*models.Claims)

		if claims == nil || !claims.IsAdmin {
			http.Error(w, "Unauthorized: admin access required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func VerifyCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*models.Claims)

		if claims == nil || claims.IsAdmin {
			http.Error(w, "Unauthorized: customer access required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
