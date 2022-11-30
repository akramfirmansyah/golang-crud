package models

import "time"

type Mahasiswa struct {
	NIM      string    `gorm:"primaryKey" json:"nim"`
	Name     string    `json:"name"`
	Angkatan uint      `json:"angkatan"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
