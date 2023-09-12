package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"go_share/helper"
	"go_share/model/api"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
	DB      *gorm.DB
}

func NewAuthMiddleware(handler http.Handler, db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, DB: db}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.RequestURI == "/api/v1/auth/login" || request.RequestURI == "/api/v1/auth/register" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		token := request.Header.Get("Authorization")
		parts := strings.Split(token, "Bearer ")

		if len(parts) > 1 {
			tokenValue := parts[1]
			result, err := helper.VerifyToken(tokenValue)
			if err != nil {
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				apiResponse := api.ApiResponseGeneral{
					Code:   http.StatusUnauthorized,
					Status: "USER UNAUTHORIZED",
				}

				helper.WriteToResponse(writer, apiResponse)
			}

			if _, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
				//id := claims["sub"]

				//query := "SELECT id_user FROM ref_users WHERE id_user = ? AND token = ?"
				//rows, _ := middleware.DB.Query(query, id, tokenValue)
				//defer rows.Close()
				//
				//user := domain.User{}
				//if rows.Next() {
				//	rows.Scan(&user.IdUser)
				//	middleware.Handler.ServeHTTP(writer, request)
				//} else {
				//	writer.Header().Set("Content-Type", "application/json")
				//	writer.WriteHeader(http.StatusUnauthorized)
				//
				//	apiResponse := api.ApiResponseGeneral{
				//		Code:   http.StatusUnauthorized,
				//		Status: "USER UNAUTHORIZED",
				//	}
				//
				//	helper.WriteToResponse(writer, apiResponse)
				//}

				middleware.Handler.ServeHTTP(writer, request)
			} else {
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				apiResponse := api.ApiResponseGeneral{
					Code:   http.StatusUnauthorized,
					Status: "USER UNAUTHORIZED",
				}

				helper.WriteToResponse(writer, apiResponse)
			}
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			apiResponse := api.ApiResponseGeneral{
				Code:   http.StatusUnauthorized,
				Status: "USER UNAUTHORIZED",
			}

			helper.WriteToResponse(writer, apiResponse)
		}
	}
}
