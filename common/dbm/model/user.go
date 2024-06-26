package model

import (
	"time"

	"github.com/edgehook/apphub-dashboard-server/common/global"
	"k8s.io/klog/v2"
)

type User struct {
	ID              string `gorm:"column:id; type:varchar(36); primary_key;" json:"id"`
	Name            string `gorm:"column:name; type:varchar(256);" json:"name"`
	Password        string `gorm:"column:password; type:varchar(256);" json:"password"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp;" json:"createTimeStamp"`
	UpdateTimeStamp int64  `gorm:"column:update_time_stamp;autoUpdateTime:milli" json:"updateTimeStamp"`
}

func (d *User) TableName() string {
	return "project_details"
}

func GetUserByPage(page int, limit int, keywords string) ([]*User, error) {
	var users []*User
	err := global.DBAccess.Where("name LIKE ?", "%"+keywords+"%").Offset((page - 1) * limit).Order("update_time_stamp desc").Limit(limit).Find(&users).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return users, err
}
func GetUserCountByKeywords(keywords string) (int64, error) {
	var count int64
	err := global.DBAccess.Model(&User{}).Where("name LIKE ?", "%"+keywords+"%").Count(&count).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return 0, err
	}
	return count, err
}

func AddUser(user *User) error {
	user.CreateTimeStamp = time.Now().UnixNano() / 1e6
	err := global.DBAccess.Create(&user).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func SaveUser(id int64, user *User) error {
	err := global.DBAccess.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"password": user.Password,
	}).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}
