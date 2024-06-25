package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func Authentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientToken := r.Header.Get("Authorization")
			clientToken = strings.TrimPrefix(clientToken, "Bearer ")
			if clientToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Unauthorized: No authorization header provided")
				return
			}
			claims, err := ValidateToken(clientToken)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Unauthorized: Invalid token (%v)", err)
				return
			}
			// Set user data on context (assuming claims has Email and Uid)
			ctx := context.WithValue(r.Context(), "email", claims.Email)
			ctx = context.WithValue(ctx, "id", claims.ID)
			r = r.WithContext(ctx)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
