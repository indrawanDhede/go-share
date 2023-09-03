package api_response

import "database/sql"

type AuthRegisterResponse struct {
	Email string         `json:"email"`
	Tiket sql.NullString `json:"tiket"`
}
