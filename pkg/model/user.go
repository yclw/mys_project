package model

// User 用户模型
type User struct {
	ID              int64  `gorm:"primaryKey;autoIncrement;column:id;comment:系统前台用户表" json:"id"`
	Account         string `gorm:"column:account;size:20;not null;default:'';comment:用户登陆账号" json:"account"`
	Password        string `gorm:"column:password;size:64;default:'';comment:登陆密码" json:"-"`
	Name            string `gorm:"column:name;size:255;default:'';comment:用户昵称" json:"name"`
	Mobile          string `gorm:"column:mobile;size:255;comment:手机" json:"mobile"`
	RealName        string `gorm:"column:realname;size:255;comment:真实姓名" json:"realname"`
	CreateTime      string `gorm:"column:create_time;size:30;comment:创建时间" json:"create_time"`
	Status          int8   `gorm:"column:status;default:0;comment:状态" json:"status"`
	LastLoginTime   string `gorm:"column:last_login_time;size:30;comment:上次登录时间" json:"last_login_time"`
	Sex             int8   `gorm:"column:sex;default:0;comment:性别" json:"sex"`
	Avatar          string `gorm:"column:avatar;size:255;default:'';comment:头像" json:"avatar"`
	IDCard          string `gorm:"column:idcard;size:255;comment:身份证" json:"idcard"`
	Province        int    `gorm:"column:province;default:0;comment:省" json:"province"`
	City            int    `gorm:"column:city;default:0;comment:市" json:"city"`
	Area            int    `gorm:"column:area;default:0;comment:区" json:"area"`
	Address         string `gorm:"column:address;size:255;comment:所在地址" json:"address"`
	Description     string `gorm:"column:description;type:text;comment:备注" json:"description"`
	Email           string `gorm:"column:email;size:255;comment:邮箱" json:"email"`
	DingTalkOpenID  string `gorm:"column:dingtalk_openid;size:50;comment:钉钉openid" json:"dingtalk_openid"`
	DingTalkUnionID string `gorm:"column:dingtalk_unionid;size:50;comment:钉钉unionid" json:"dingtalk_unionid"`
	DingTalkUserID  string `gorm:"column:dingtalk_userid;size:50;comment:钉钉用户id" json:"dingtalk_userid"`
}

// TableName 指定表名
func (User) TableName() string {
	return "t_user"
}
