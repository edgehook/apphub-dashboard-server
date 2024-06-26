package model

import (
	"github.com/edgehook/apphub-dashboard-server/common/global"
	"github.com/edgehook/apphub-dashboard-server/common/utils"
	"k8s.io/klog/v2"
)

type ScreenImage struct {
	Id              string `gorm:"column:id; type:varchar(36); primary_key;" json:"id"`
	ScreenId        string `gorm:"column:screen_id; not null; type:varchar(256);" json:"screen_id"`
	Data            string `gorm:"column:data; type:text;" json:"data"`
	Type            string `gorm:"column:type; not null; type:varchar(256);" json:"type"`
	Filename        string `gorm:"column:filename; default:null; type:varchar(512);" json:"filename"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp;" json:"createTimeStamp"`
	UpdateTimeStamp int64  `gorm:"column:update_time_stamp;autoUpdateTime:milli" json:"updateTimeStamp"`
}

func (d *ScreenImage) TableName() string {
	return "screen_image"
}

func GetScreenImagesByScreenId(screenId string) ([]*ScreenImage, error) {
	var screenImages []*ScreenImage
	err := global.DBAccess.Where("screen_id = ?", screenId).Find(&screenImages).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}

	return screenImages, nil
}

func GetScreenImagesByScreenIdAndType(screenId string, ctype string) ([]*ScreenImage, error) {
	var screenImages []*ScreenImage
	err := global.DBAccess.Where("screen_id = ? and type = ?", screenId, ctype).Find(&screenImages).Error
	if err != nil {
		klog.Errorf("err: %v", err)
	}

	return screenImages, err
}

func AddScreenImage(ScreenImage *ScreenImage) error {
	ScreenImage.CreateTimeStamp = utils.GetNowTimeStamp()
	err := global.DBAccess.Create(&ScreenImage).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func SaveScreenImage(id, data string) error {
	err := global.DBAccess.Model(&ScreenImage{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Data": data,
	}).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func DeleteScreenImage(id string) error {
	err := global.DBAccess.Where("id = ?", id).Delete(&ScreenImage{}).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}
