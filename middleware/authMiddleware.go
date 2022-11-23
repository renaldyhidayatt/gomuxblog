package middleware

import (
	"muxblog/helpers"
	"muxblog/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			helpers.ResponseWithError(w, http.StatusUnauthorized, "Authentication failure")
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := utils.VerifyToken(tokenString)

		if err != nil {
			helpers.ResponseWithError(w, http.StatusUnauthorized, "Error Verify JWT Token: "+err.Error())
			return
		}

		userId := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}
