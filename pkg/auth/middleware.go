package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
)

func Authentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			clientToken := r.Header.Get("Authorization")
			clientToken = strings.TrimPrefix(clientToken, "Bearer ")

			if clientToken == "" {
				log.Println("error no authorization header provided")
				err := errors.New("no authorization header provided")
				render.Render(w, r, bondApi.ErrRender(err, http.StatusUnauthorized, bondApi.ErrNoAuth))
				return
			}

			claims, err := ValidateToken(clientToken)
			if err != nil {
				log.Println("error validating token")
				render.Render(w, r, bondApi.ErrRender(err, http.StatusUnauthorized, bondApi.ErrNoAuth))
				return
			}

			ctx := context.WithValue(r.Context(), "email", claims.Email)
			ctx = context.WithValue(ctx, "id", claims.ID)
			r = r.WithContext(ctx)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
