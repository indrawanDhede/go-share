package domain

import (
	"database/sql"
	"time"
)

type User struct {
	IdUser            int
	IdSocket          sql.NullString
	Nama              string
	Email             string
	Password          string
	IdLembaga         int
	Token             sql.NullString
	Tiket             sql.NullString
	LinkFoto          sql.NullString
	NoHp              sql.NullString
	JenjangPendidikan sql.NullString
	Alamat            sql.NullString
	Bahasa            sql.NullString
	Kompetensi        sql.NullString
	Status            int8
	IsLogin           string
	Udcr              time.Time
	Udch              time.Time
}
