package do

import "time"

type User struct {
	Id          uint64     `gorm:"column:id;primary_key;NOT NULL"`                   // 主键
	Name        string     `json:"name" gorm:"column:name;NOT NULL"`                 //账户
	Password    string     `json:"password" gorm:"column:password;NOT NULL"`         //密码
	Gender      int        `json:"gender" gorm:"column:gender;NOT NULL"`             //性别
	Nick        string     `json:"nick" gorm:"column:nick;NOT NULL"`                 //昵称
	Avatar      uint32     `json:"avatar" gorm:"column:avatar;NOT NULL"`             //头像
	AvatarState int        `json:"avatar_state" gorm:"column:avatar_state;NOT NULL"` //头像状态
	Nationality string     `json:"nationality" gorm:"column:nationality;NOT NULL"`   //国籍
	Birthday    time.Time  `json:"birthday" gorm:"column:birthday;NOT NULL"`         //生日
	Deleted     uint32     `json:"deleted" gorm:"column:deleted;default:0;NOT NULL"` // 删除状态：0-正常
	CreateTime  *time.Time `json:"create_time" gorm:"column:create_time;NOT NULL;default:CURRENT_TIMESTAMP(3)"`
	UpdateTime  *time.Time `json:"update_time" gorm:"column:update_time;NOT NULL;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
}
