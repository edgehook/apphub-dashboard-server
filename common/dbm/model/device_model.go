package model

import (
	"github.com/edgehook/ithings/common/global"
	"k8s.io/klog/v2"
	"time"
)

type DeviceModel struct {
	ID                     int64  `gorm:"primary_key; auto_increment" json:"id"`
	Name                   string `gorm:"column: name; unique; not null; type:varchar(256); unique" json:"name"`
	ModelId                string `gorm:"column: modelId; not null; type:varchar(256); unique" json:"modelId"`
	UserGroup              string `gorm:"column: userGroup; default:null; type:varchar(256);" json:"userGroup"`
	Manufacturer           string `gorm:"column: manufacturer; default:null; type:varchar(256);" json:"manufacturer"`
	Industry               string `gorm:"column: industry;default:null; type:varchar(256);" json:"industry"`
	DataType               string `gorm:"column: dataType; default:null; type:varchar(256);" json:"dataType"`
	Description            string `gorm:"column: description; default:null; type:varchar(256);" json:"description"`
	Time                   int64  `gorm:"column: time; default:null" json:"createTime"`
	RegisteredDeviceNumber int64  `gorm:"column: registeredDeviceNumber; type:integer" json:"registeredDeviceNumber"`
}

func (d *DeviceModel) TableName() string {
	return "deviceModel"
}

func GetModelByPage(page int, limit int) ([]DeviceModel, error) {
	var models []DeviceModel
	global.DBAccess.Begin()
	err := global.DBAccess.Offset((page - 1) * limit).Limit(limit).Find(&models).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		global.DBAccess.Rollback()
		return nil, err
	}
	return models, err
}

func GetModelByPageAndKeywords(page int, limit int, keywords string) ([]DeviceModel, error) {
	var models []DeviceModel
	global.DBAccess.Begin()
	err := global.DBAccess.Offset((page-1)*limit).Limit(limit).Where("name LIKE ?", "%"+keywords+"%").Find(&models).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		global.DBAccess.Rollback()
		return nil, err
	}
	return models, err
}
func GetDeviceModelCount() (int64, error) {
	var count int64
	global.DBAccess.Begin()
	err := global.DBAccess.Model(&DeviceModel{}).Count(&count).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		global.DBAccess.Rollback()
		return -1, err
	}
	return count, err
}

func AddDeviceModel(deviceModel *DeviceModel) error {
	global.DBAccess.Begin()
	deviceModel.Time = time.Now().Unix()
	err := global.DBAccess.Create(&deviceModel).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		global.DBAccess.Rollback()
		return err
	}
	global.DBAccess.Commit()
	return nil
}

func SaveDeviceModel(deviceModel *DeviceModel) error {
	global.DBAccess.Begin()
	err := global.DBAccess.Save(&deviceModel).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		global.DBAccess.Rollback()
		return err
	}
	global.DBAccess.Commit()
	return nil
}
