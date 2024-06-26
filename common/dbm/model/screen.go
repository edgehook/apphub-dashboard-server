package model

import (
	"github.com/edgehook/apphub-dashboard-server/common/global"
	"github.com/edgehook/apphub-dashboard-server/common/utils"
	"k8s.io/klog/v2"
)

type Screen struct {
	Id              string `gorm:"column:id; type:varchar(36); primary_key;" json:"id"`
	Name            string `gorm:"column:name; not null; type:varchar(256);" json:"name"`
	Data            string `gorm:"column:data; type:text;" json:"data"`
	State           int    `gorm:"column:state; not null; type:int;" json:"state"`
	IsAside         string `gorm:"column:is_aside;type:varchar(256)" json:"isAside"`
	Image           string `gorm:"column:img; type:text;" json:"img"`
	Description     string `gorm:"column:description; default:null; type:varchar(512);" json:"description"`
	Locked          string `gorm:"column:locked; type:varchar(256);" json:"locked"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp;" json:"createTimeStamp"`
	UpdateTimeStamp int64  `gorm:"column:update_time_stamp;autoUpdateTime:milli" json:"updateTimeStamp"`
}

func (d *Screen) TableName() string {
	return "screen"
}

func GetScreenByPageAndKeywords(page int, limit int, keywords string) ([]*Screen, error) {
	var screens []*Screen
	tx := global.DBAccess.Model(&Screen{}).Offset((page - 1) * limit).Limit(limit)
	if keywords != "" {
		tx = tx.Where("name LIKE ?", "%"+keywords+"%")
	}
	err := tx.Order("create_time_stamp desc").Find(&screens).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return screens, err
}

func GetScreenByKeywords(keywords string) ([]*Screen, error) {
	var screens []*Screen
	tx := global.DBAccess.Model(&Screen{})
	if keywords != "" {
		tx = tx.Where("name LIKE ?", "%"+keywords+"%")
	}
	err := tx.Order("create_time_stamp desc").Find(&screens).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return screens, err
}

func GetScreenByIsAside(isAside string) ([]*Screen, error) {
	var screens []*Screen
	tx := global.DBAccess.Model(&Screen{})
	if isAside != "" {
		tx = tx.Where("is_aside = ?", isAside)
	}
	err := tx.Find(&screens).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return screens, err
}

func GetScreenCount(keywords string) (int64, error) {
	var count int64
	tx := global.DBAccess.Model(&Screen{})
	if keywords != "" {
		tx = tx.Where("name LIKE ?", "%"+keywords+"%")
	}
	err := tx.Count(&count).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return -1, err
	}
	return count, err
}

func GetScreenById(id string) (*Screen, error) {
	screen := &Screen{}
	err := global.DBAccess.Where("id = ?", id).First(screen).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}

	return screen, nil
}

func IsExistScreenByName(name string) bool {
	var count int64
	err := global.DBAccess.Model(&Screen{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
func AddScreen(Screen *Screen) error {
	Screen.CreateTimeStamp = utils.GetNowTimeStamp()
	err := global.DBAccess.Create(&Screen).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func SaveScreen(id string, name string, state *int, data string, img string, isAside string) error {
	vals := make(map[string]interface{})
	if name != "" {
		vals["name"] = name
	}
	if state != nil {
		vals["state"] = state
	}

	if data != "" {
		vals["data"] = data
	}

	if img != "" {
		vals["img"] = img
	}

	if isAside != "" {
		vals["is_aside"] = isAside
	}

	err := global.DBAccess.Model(&Screen{}).Where("id = ?", id).Updates(vals).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func DeleteScreen(id string) error {
	err := global.DBAccess.Where("id = ?", id).Delete(&Screen{}).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}
