package model

type Paper struct {
	ID       uint   `gorm:"primary_key"`
	Year     uint   `gorm:"not null"`
	SchName  string `gorm:"type: varchar(20);not null;"`
	SubName  string `gorm:"type: varchar(20);not null;"`
	PaperNum uint   `gorm:"not null"`
	UsedNum  uint   `gorm:"not null"`
}
