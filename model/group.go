package model

import "time"

// Group 群列表
type Group struct {
	Id        int64     `xorm:"pk autoincr bigint(20)" from:"id" json:"id"`           // 表ID，自增
	UserId    int64     `xorm:"bigint(20)" from:"user_id" json:"user_id"`             // 记录是谁的即当前人添加的别人
	DstId     int64     `xorm:"bigint(20)" from:"dst_id" json:"dbt_id"`               // 添加的目标信息
	Type      int       `xorm:"int(11)" from:"type" json:"type"`                      // 记录类型：群聊
	Memo      string    `xorm:"varchar(120)" from:"memo" json:"memo"`                 // 备注
	CreatedAt time.Time `xorm:"datetime created" from:"created_at" json:"create_at"`  // created: 在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"datetime updated" from:"updated_at" json:"updated_at"` // updated: 在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"datetime deleted" from:"deleted_at" json:"deleted_at"` // deleted: 在Delete时设置为当前时间，并且当前记录不删除

}

const (
	CONTACT_TYPE_GROUP = 0x02
)
