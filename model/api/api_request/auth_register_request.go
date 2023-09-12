package api_request

type AuthRegisterRequest struct {
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IdLembaga int    `json:"id_lembaga"`
}
