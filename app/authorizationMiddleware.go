package app

import (
	"net/http"
	"strings"

	"github.com/adityapwr/banking-lib/errs"
	"github.com/adityapwr/go-banking/domain"
	"github.com/gorilla/mux"
)

type AuthorizationMiddleware struct {
	repo domain.AuthorizationRepository
}

func (s AuthorizationMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader != "" {
				authorizationToken := getTokenFromHeader(authorizationHeader)

				isAuthorized := s.repo.IsAuthorized(authorizationToken, currentRoute.GetName(), currentRouteVars)
				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}

			} else {
				writeResponse(w, http.StatusUnauthorized, "missing authorization header")

			}

		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
