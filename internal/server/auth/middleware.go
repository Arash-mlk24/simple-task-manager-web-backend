package auth

import (
	"context"
	"fmt"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/service_errors"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server/handlers"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Middleware(next http.Handler, allowedRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			errUnauthorized := service_errors.ErrUnauthorized
			response := handlers.ApiFailure(errUnauthorized.Code, errUnauthorized.Message)
			utils.RespondJSON(w, errUnauthorized.HttpStatus, response)
			return
		}

		claims := &Claims{}

		fmt.Println("claims 1:", claims)

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		fmt.Println("claims 2:", claims)

		if err != nil || !token.Valid {
			errUnauthorized := service_errors.ErrUnauthorized
			response := handlers.ApiFailure(errUnauthorized.Code, errUnauthorized.Message)
			utils.RespondJSON(w, errUnauthorized.HttpStatus, response)
			return
		}

		for _, role := range claims.Roles {
			for _, allowed := range allowedRoles {
				if role == allowed {
					ctx := r.Context()
					ctx = setUserContext(ctx, claims)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
		}

		errForbidden := service_errors.ErrForbidden
		response := handlers.ApiFailure(errForbidden.Code, errForbidden.Message)
		utils.RespondJSON(w, errForbidden.HttpStatus, response)
	})
}

type contextKey string

var userContextKey = contextKey("user")

func setUserContext(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, userContextKey, claims)
}

func GetUserFromContext(ctx context.Context) *Claims {
	claims, ok := ctx.Value(userContextKey).(*Claims)
	if !ok {
		return nil
	}
	return claims
}
