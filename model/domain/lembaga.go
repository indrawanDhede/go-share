package domain

type Lembaga struct {
	ID            int    `gorm:"primaryKey;autoIncrement:true;column:id_lembaga"`
	Nama          string `gorm:"column:nama"`
	Alamat        string `gorm:"column:alamat"`
	IdTipeLembaga int    `gorm:"column:id_tipe_lembaga"`
}

func (Lembaga) TableName() string {
	return "ref_lembaga"
}
