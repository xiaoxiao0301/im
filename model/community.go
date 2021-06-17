package model

import "time"

// Community 群具体信息
type Community struct {
	Id        int64     `xorm:"pk autoincr bigint(20)" from:"id" json:"id"`           // 表ID，自增
	UserId    int64     `xorm:"bigint(20)" from:"user_id" json:"user_id"`             // 群主
	Name      string    `xorm:"varchar(30)" from:"name" json:"name"`                  // 群名称
	Icon      string    `xorm:"varchar(250)" from:"icon" json:"icon"`                 // 群头像
	Cate      int       `xorm:"int(11)" from:"cate" json:"cate"`                      //群类型
	Number    int       `xorm:"int(11)" from:"number" json:"number"`                  //群人数
	Memo      string    `xorm:"varchar(120)" from:"memo" json:"memo"`                 // 群简介
	CreatedAt time.Time `xorm:"datetime created" from:"created_at" json:"create_at"`  // created: 在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"datetime updated" from:"updated_at" json:"updated_at"` // updated: 在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"datetime deleted" from:"deleted_at" json:"deleted_at"` // deleted: 在Delete时设置为当前时间，并且当前记录不删除
}

const (
	COMMUNITY_CATE_COM = 0x01 // 通用型
)
