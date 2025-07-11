package model

// User 用户模型
type UserBase struct {
	UserID     int64  `gorm:"primaryKey;autoIncrement;column:user_id;comment:用户唯一ID" json:"user_id"`
	Account    string `gorm:"column:account;size:20;not null;default:'';comment:用户登陆账号" json:"account"`
	Password   string `gorm:"column:password;size:64;default:'';comment:登陆密码" json:"-"`
	Name       string `gorm:"column:name;size:255;default:'';comment:用户昵称" json:"name"`
	Avatar     string `gorm:"column:avatar;size:255;default:'';comment:头像" json:"avatar"`
	CreateTime string `gorm:"column:create_time;size:30;comment:创建时间" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;size:30;comment:更新时间" json:"update_time"`
	Status     int8   `gorm:"column:status;default:0;comment:状态" json:"status"`
}

type User struct {
	UserBase
	Email         string `gorm:"column:email;size:255;comment:邮箱" json:"email"`
	Mobile        string `gorm:"column:mobile;size:255;comment:手机" json:"mobile"`
	Sex           int8   `gorm:"column:sex;default:0;comment:性别" json:"sex"`
	LastLoginTime string `gorm:"column:last_login_time;size:30;comment:上次登录时间" json:"last_login_time"`
}

func (User) TableName() string {
	return "t_user"
}
