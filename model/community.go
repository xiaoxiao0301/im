package model

import "time"

// Community 群具体信息
type Community struct {
	Id        int64     `xorm:"pk autoincr bigint(20)" json:"id"`   // 表ID，自增
	UserId    int64     `xorm:"bigint(20)" json:"user_id"`          // 群主
	Name      string    `xorm:"varchar(30)" json:"name"`            // 群名称
	Icon      string    `xorm:"varchar(250)" json:"icon"`           // 群头像
	Cate      int       `xorm:"int(11)" json:"cate"`                //群类型
	Number    int       `xorm:"int(11)" json:"number"`              //群人数
	Memo      string    `xorm:"varchar(120)" json:"memo"`           // 群简介
	CreatedAt time.Time `xorm:"datetime created" json:"create_at"`  // created: 在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"datetime updated" json:"updated_at"` // updated: 在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"datetime deleted" json:"deleted_at"` // deleted: 在Delete时设置为当前时间，并且当前记录不删除
}

const (
	COMMUNITY_CATE_COM      = 0x01 // 通用型
	COMMUNITY_CATE_HOBBY    = 0x02 // 兴趣爱好
	COMMUNITY_CATE_INDUSTRY = 0x03 // 行业交流
	COMMUNITY_CATE_LIFE     = 0x04 // 生活休闲
	COMMUNITY_CATE_STUDY    = 0x05 // 学习考试
)
