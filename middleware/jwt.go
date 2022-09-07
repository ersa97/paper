package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ersa97/paper-test/helpers"
	"github.com/ersa97/paper-test/models"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		if request.Header["Authorization"] != nil {
			authorizationToken := strings.Split(request.Header["Authorization"][0], " ")
			if len(authorizationToken) != 2 {
				json.NewEncoder(response).Encode(models.Response{
					Message: "invalid authorization token",
					Data:    nil,
				})

				return
			}

			if authorizationToken[0] != "Bearer" {
				json.NewEncoder(response).Encode(models.Response{
					Message: "authorization token type does not match",
					Data:    nil,
				})

				return
			}

			var err error
			var token *jwt.Token
			var jwtKeyStatus string
			jwtKeys := helpers.JWTKeys
			for i := 0; i < len(jwtKeys); i++ {
				token, err = jwt.Parse(authorizationToken[1], func(token *jwt.Token) (interface{}, error) {
					_, _ = token.Method.(*jwt.SigningMethodHMAC)

					return []byte(jwtKeys), nil
				})
				if err != nil {
					jwtKeyStatus = "error"
				} else {
					jwtKeyStatus = "success"

					break
				}
			}
			if jwtKeyStatus != "success" {
				json.NewEncoder(response).Encode(models.Response{
					Message: "authorization token credentials do not match",
					Data:    nil,
				})

				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				json.NewEncoder(response).Encode(models.Response{
					Message: "invalid authorization token credentials",
					Data:    nil,
				})

				return
			}

			if token.Valid {
				ctx := context.WithValue(request.Context(), "authorizationToken", claims)

				request = request.WithContext(ctx)

				timeData, _ := time.Parse("2006-01-02 15:04:05", request.Context().Value("authorizationToken").(jwt.MapClaims)["expired"].(string))
				currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
				if currentTime.After(timeData) {
					json.NewEncoder(response).Encode(models.Response{
						Message: "authorization token has expired",
						Data:    nil,
					})

					return
				}

				next(response, request)
			}
		} else {
			json.NewEncoder(response).Encode(models.Response{
				Message: "unauthorized",
				Data:    nil,
			})

			return
		}
	})
}
