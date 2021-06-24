package model

import "time"

type TestTime struct {
	Id        int64     `xorm:"bigint(20) pk autoincr" json:"id"`
	Name      string    `xorm:"varchar(10)" json:"name"`
	CreatedAt Mytimes   `xorm:"timestamp created" json:"created_at"`
	UpdatedAt time.Time `xorm:"timestamp updated" json:"updated_at"`
}
