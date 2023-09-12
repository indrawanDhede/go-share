package api_response

import "database/sql"

type LembagaResponse struct {
	Id_Lembaga int            `json:"id_lembaga"`
	Nama       string         `json:"nama"`
	Alamat     sql.NullString `json:"alamat"`
}
