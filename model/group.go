package model

// Group 群列表
type Group struct {
	Id        int64   `xorm:"pk autoincr bigint(20)" json:"id"`    // 表ID，自增
	UserId    int64   `xorm:"bigint(20)" json:"user_id"`           // 记录是谁的即当前人添加的别人
	DstId     int64   `xorm:"bigint(20)" json:"dst_id"`            // 添加的目标信息
	Type      int     `xorm:"int(11)" json:"type"`                 // 记录类型：群聊
	Memo      string  `xorm:"varchar(120)" json:"memo"`            // 备注
	CreatedAt Mytimes `xorm:"timestamp created" json:"create_at"`  // created: 在Insert时自动赋值为当前时间
	UpdatedAt Mytimes `xorm:"timestamp updated" json:"updated_at"` // updated: 在Insert或Update时自动赋值为当前时间
	DeletedAt Mytimes `xorm:"timestamp deleted" json:"deleted_at"` // deleted: 在Delete时设置为当前时间，并且当前记录不删除

}

const (
	CONTACT_TYPE_COMMUNITY = 0x02
)
