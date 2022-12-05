package model

type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"type: varchar(20);not null"`
	Password string `gorm:"type: varchar(20);not null"`
	Identify string `gorm:"type: varchar(20);not null"`
}
