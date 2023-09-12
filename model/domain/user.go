package domain

import (
	"database/sql"
)

type StatusLogin string

const (
	StatusOffline StatusLogin = "N"
	StatusOnline  StatusLogin = "Y"
)

type User struct {
	ID                int            `gorm:"primaryKey;autoIncrement:true;column:id_user"`
	IdSocket          sql.NullString `gorm:"column:id_socket"`
	Nama              string         `gorm:"column:nama"`
	Email             string         `gorm:"type:varchar(255);uniqueIndex"`
	Password          string         `gorm:"column:password"`
	IdLembaga         int            `gorm:"column:id_lembaga"`
	Token             sql.NullString `gorm:"column:token"`
	Tiket             sql.NullString `gorm:"column:tiket"`
	LinkFoto          sql.NullString `gorm:"column:link_foto"`
	NoHp              sql.NullString `gorm:"column:no_hp"`
	JenjangPendidikan sql.NullString `gorm:"column:jenjang_pendidikan"`
	Alamat            sql.NullString `gorm:"column:alamat"`
	Bahasa            sql.NullString `gorm:"column:bahasa"`
	Kompetensi        sql.NullString `gorm:"column:kompetensi"`
	Status            int8           `gorm:"column:status"`
	IsLogin           StatusLogin    `gorm:"type:enum('N', 'Y')"`
	Lembaga           *Lembaga       `gorm:"foreignKey:IdLembaga"`
}

func (User) TableName() string {
	return "ref_users"
}
