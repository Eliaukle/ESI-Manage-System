package model

type Subject struct {
	ID      uint   `gorm:"primary_key"`
	SubName string `gorm:"type: varchar(20);not null;unique;"`
	SubType string `gorm:"type: varchar(20);not null"`
}

type SubRank struct {
	ID           uint   `gorm:"primary_key"`
	Year         uint   `gorm:"not null"`
	SchName      string `gorm:"type: varchar(20);not null;"`
	SubName      string `gorm:"type: varchar(20);not null;"`
	SubCountRank uint   `gorm:"not null"`
	SubWorldRank uint   `gorm:"not null"`
}
