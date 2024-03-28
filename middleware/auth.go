package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

const AuthUserID = "middleware.auth.userID" // Unique key

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer") {
			writeUnauthed(w)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, "Bearer ")

		// Decode the token from base 64
		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			writeUnauthed(w)
			return
		}

		// Just assuming a valid base64 token is a valid user id.
		userID := string(token)

		ctx := context.WithValue(r.Context(), AuthUserID, userID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
