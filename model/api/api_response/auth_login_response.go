package api_response

type AuthLoginResponse struct {
	Id_User            int    `json:"id_user"`
	Nama               string `json:"nama"`
	Email              string `json:"email"`
	Token              string `json:"token"`
	Link_Foto          string `json:"link_foto"`
	No_Hp              string `json:"no_hp"`
	Jenjang_pendidikan string `json:"jenjang_pendidikan"`
	Alamat             string `json:"alamat"`
	Bahasa             string `json:"bahasa"`
	Kompetensi         string `json:"kompetensi"`
	Is_Login           string `json:"is_login"`
}
