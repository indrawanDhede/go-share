package api_response

import "database/sql"

type UserResponse struct {
	Id_User            int            `json:"id_user"`
	Nama               string         `json:"nama"`
	Email              string         `json:"email"`
	Token              sql.NullString `json:"token"`
	Link_Foto          sql.NullString `json:"link_foto"`
	No_Hp              sql.NullString `json:"no_hp"`
	Jenjang_pendidikan sql.NullString `json:"jenjang_pendidikan"`
	Alamat             sql.NullString `json:"alamat"`
	Bahasa             sql.NullString `json:"bahasa"`
	Kompetensi         sql.NullString `json:"kompetensi"`
	Is_Login           string         `json:"is_login"`
}
