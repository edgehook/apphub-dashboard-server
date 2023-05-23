package model

import (
	"github.com/edgehook/ithings/common/global"
	"k8s.io/klog/v2"
	"time"
)

type ProjectDetails struct {
	ID              int64  `gorm:"primary_key; auto_increment" json:"id"`
	Content         string `gorm:"column: content; not null; type:text;" json:"content"`
	ProjectId       string `gorm:"column: project_id; not null; type:varchar(256);" json:"project_id"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp;" json:"createTimeStamp"`
	UpdateTimeStamp int64  `gorm:"column:update_time_stamp;autoUpdateTime:milli" json:"updateTimeStamp"`
}

func (d *ProjectDetails) TableName() string {
	return "project_details"
}

func GetProjectDetailsByProjectId(projectId string) (*ProjectDetails, error) {
	projectDetails := &ProjectDetails{}
	err := global.DBAccess.Where("project_id = ?", projectId).First(projectDetails).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return nil, err
	}

	return projectDetails, nil
}

func AddProjectDetails(projectDetails *ProjectDetails) error {
	projectDetails.CreateTimeStamp = time.Now().UnixNano() / 1e6
	err := global.DBAccess.Create(&projectDetails).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}

func SaveProjectDetails(id int64, projectDetails *ProjectDetails) error {
	err := global.DBAccess.Model(&ProjectDetails{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Content": projectDetails.Content,
	}).Error
	if err != nil {
		klog.Errorf("err: %v", err)
		return err
	}
	return nil
}
