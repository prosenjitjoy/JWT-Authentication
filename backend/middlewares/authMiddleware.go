package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"server/internal/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("token")
		if err == http.ErrNoCookie {
			log.Println("Unauthorized attempt! No auth cookie")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
			return
		} else if err != nil {
			log.Println("Unable to fetch cookie")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
			return
		}

		claims, err := validateToken(authCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims.ID)
		ctx = context.WithValue(ctx, "user_name", claims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func validateToken(clientToken string) (claims *user.SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(clientToken, &user.SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		err = errors.New("failed to parse token")
		return
	}

	claims, ok := token.Claims.(*user.SignedDetails)
	if !ok {
		err = errors.New("the token is invalid")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		err = errors.New("the token is expired")
		return
	}

	return claims, nil
}
