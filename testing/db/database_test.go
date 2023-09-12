package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type Users struct {
	ID        int     `gorm:"primaryKey;column:id_user"`
	Nama      string  `gorm:"column:nama"`
	Email     string  `gorm:"type:varchar(255);uniqueIndex"`
	IdLembaga int     `gorm:"column:id_lembaga"`
	Lembaga   Lembaga `gorm:"foreignKey:IdLembaga"`
}

type Lembaga struct {
	ID     int    `gorm:"primaryKey;column:id_lembaga"`
	Nama   string `gorm:"column:nama"`
	Alamat string `gorm:"column:alamat"`
}

func (Lembaga) TableName() string {
	return "ref_lembaga"
}

func (Users) TableName() string {
	return "ref_users"
}

func NewDb() *gorm.DB {
	dsn := "root:root@tcp(localhost:8889)/schoolshare?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func TestDB(t *testing.T) {
	db := NewDb()

	var users []Users
	db.Preload("Lembaga").Find(&users)

	for _, user := range users {
		// Anda dapat mengakses data yang dimuat menggunakan Preload
		nama := user.Nama
		fmt.Println(nama)
		id := user.Lembaga.ID
		fmt.Println(id)
		alamat := user.Lembaga.Alamat
		fmt.Println(alamat)
		fmt.Println()

		// Lakukan sesuatu dengan nama dan alamat
	}
}
