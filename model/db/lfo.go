package db

import "time"

//Member o
type Member struct {
	Id       int       `xorm:"not null pk autoincr INTEGER" json:"id" `
	Username string    `xorm:"not null VARCHAR(50)" json:"username"`
	Password string    `xorm:"not null VARCHAR(50)" json:"password"`
	Time     time.Time `xorm:"default 'now()' DATETIME" json:"time"`
}
