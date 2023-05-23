package model

import (
	"github.com/edgehook/ithings/common/global"
	"k8s.io/klog/v2"
	"time"
)

type Project struct {
	Id              string `gorm:"column:id; type:varchar(36); primary_key;" json:"id"`
	Name            string `gorm:"column: name; not null; type:varchar(256);" json:"name"`
	State           int    `gorm:"column: state; not null; type:int;" json:"state"`
	Description     string `gorm:"column: description; default:null; type:varchar(256);" json:"description"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp;" json:"createTimeStamp"`
	UpdateTimeStamp int64  `gorm:"column:update_time_stamp;autoUpdateTime:milli" json:"updateTimeStamp"`
}

func (d *Project) TableName() string {
	return "project"
}

func GetProjectByPageAndKeywords(page int, limit int, keywords string) ([]*Project, error) {
	var projects []*Project
	tx := global.DBAccess.Model(&Project{}).Offset((page - 1) * limit).Limit(limit)
	if keywords != "" {
		tx = tx.Where("name LIKE ?", "%"+keywords+"%")
	}
	err := tx.Order("update_time_stamp desc").Find(&projects).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return projects, err
}

func GetProjectByKeywords(keywords string) ([]*Project, error) {
	var projects []*Project
	tx := global.DBAccess.Model(&Project{})
	if keywords != "" {
		tx = tx.Where("name LIKE ?", "%"+keywords+"%")
	}
	err := tx.Order("update_time_stamp desc").Find(&projects).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}
	return projects, err
}

func GetProjectCount(keywords string) (int64, error) {
	var count int64
	tx := global.DBAccess.Model(&Project{})
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
func IsExistProjectByName(name string) bool {
	var count int64
	err := global.DBAccess.Model(&Project{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
func AddProject(Project *Project) error {
	Project.CreateTimeStamp = time.Now().UnixNano() / 1e6
	err := global.DBAccess.Create(&Project).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func SaveProject(id string, name string, state int) error {
	vals := make(map[string]interface{})

	vals["update_time_stamp"] = time.Now().UnixNano() / 1e6
	if name != "" {
		vals["Name"] = name
	}
	if state > -1 {
		vals["state"] = state
	}
	err := global.DBAccess.Model(&Project{}).Where("id = ?", id).Updates(vals).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func DeleteProject(id string) error {
	tx := global.DBAccess.Begin()
	if err := global.DBAccess.Where("project_id = ?", id).Delete(&ProjectDetails{}).Error; err != nil {
		klog.Errorf("err: %v", err)
		tx.Rollback()
		return err
	}
	err := global.DBAccess.Delete(&Project{}, id).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
