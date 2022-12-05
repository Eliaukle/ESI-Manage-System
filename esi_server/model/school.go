package model

type School struct {
	ID        uint   `gorm:"primary_key"`
	SchName   string `gorm:"type: varchar(20);not null;unique;"`
	SchType   string `gorm:"type: varchar(20);not null"`
	SchManage string `gorm:"type: varchar(20);not null"`
}

type SchRank struct {
	ID           uint   `gorm:"primary_key"`
	Year         uint   `gorm:"not null"`
	SchName      string `gorm:"type: varchar(20);not null;"`
	SchCountRank uint   `gorm:"not null"`
	SchWorldRank uint   `gorm:"not null"`
}
