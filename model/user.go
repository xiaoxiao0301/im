package model

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	// 用户ID
	Id int64 `xorm:"pk autoincr bigint(20)" json:"id"`
	// 用户手机号
	Mobile string `xorm:"varchar(20) comment('用户手机号')" json:"mobile"`
	// 用户密码 = f(plainpwd+salt), MD5
	Password string `xorm:"varchar(40) comment('用户密码')" json:"-"`
	// 头像
	Avatar string `xorm:"varchar(150) comment('头像')" json:"avatar"`
	// 性别
	Sex string `xorm:"varchar(2) comment('性别')" json:"sex"`
	// 昵称
	Nickname string `xorm:"varchar(20) comment('昵称')" json:"nick_name"`
	//加盐随机字符串6
	Salt   string `xorm:"varchar(10)" json:"-"`
	Online int    `xorm:"int(10)" json:"online"` //是否在线
	//前端鉴权因子, chat?id=1&token=x
	Token string `xorm:"varchar(40)" json:"token"`
	// 简介
	Memo string `xorm:"varchar(140)" json:"memo"`
	// 创建时间
	CreateAt  Mytimes `xorm:"timestamp created" json:"create_at"`
	UpdatedAt Mytimes `xorm:"timestamp updated" json:"updated_at"`
	DeletedAt Mytimes `xorm:"timestamp deleted" json:"deleted_at"`
}
