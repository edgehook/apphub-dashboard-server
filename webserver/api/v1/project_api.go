package v1

import (
	"github.com/edgehook/ithings/common/dbm/model"
	v1 "github.com/edgehook/ithings/common/types/v1"
	responce "github.com/edgehook/ithings/webserver/types"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func GetProjects(c *gin.Context) {
	var (
		err      error
		projects []*model.Project
		count    int64
	)
	keywords := c.Query("keywords")
	projects, err = model.GetProjectByKeywords(keywords)
	count, err = model.GetProjectCount(keywords)
	if err != nil {
		responce.FailWithMessage("Get count error", c)
		return
	}

	if err != nil {
		responce.FailWithMessage("Get project error", c)
		return
	}

	responce.OkWithData(map[string]interface{}{
		"list":  projects,
		"total": count,
	}, c)
}

func GetProjectDataById(c *gin.Context) {
	projectId := c.Param("id")
	if projectId == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	projectDetails, err := model.GetProjectDetailsByProjectId(projectId)
	if err != nil {
		responce.FailWithMessage("Get project content error", c)
		return
	}
	responce.OkWithData(projectDetails.Content, c)
}

func SaveProjectData(c *gin.Context) {
	var (
		projectData v1.ProjectData
	)
	projectId := c.Param("id")
	if projectId == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if err := c.Bind(&projectData); err != nil {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	projectDetails, _ := model.GetProjectDetailsByProjectId(projectId)
	if projectDetails != nil {
		projectDetails.Content = projectData.Content
		if err := model.SaveProjectDetails(projectDetails.ID, projectDetails); err != nil {
			responce.FailWithMessage("save project content error", c)
			return
		}

	} else {
		projectDetails := &model.ProjectDetails{
			Content:   projectData.Content,
			ProjectId: projectId,
		}
		if err := model.AddProjectDetails(projectDetails); err != nil {
			responce.FailWithMessage("add project content error", c)
			return
		}
	}

	responce.Ok(c)
}
func UpdateProject(c *gin.Context) {
	var (
		projectData v1.ProjectData
	)
	projectId := c.Param("id")
	if projectId == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if err := c.Bind(&projectData); err != nil {
		responce.FailWithMessage("Parameter error", c)
		return
	}

	if err := model.SaveProject(projectId, projectData.Name, *projectData.State); err != nil {
		responce.FailWithMessage("Update project error", c)
		return
	}
	responce.Ok(c)
}
func AddProject(c *gin.Context) {
	var (
		projectData v1.ProjectData
	)
	if err := c.Bind(&projectData); err != nil {
		klog.Errorln("Parameter error: %v", err.Error())
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if isExist := model.IsExistProjectByName(projectData.Name); isExist {
		responce.FailWithMessage("Project already exists", c)
		return
	}
	err := model.AddProject(&model.Project{
		Name:        projectData.Name,
		Description: projectData.Description,
		State:       *projectData.State,
	})
	if err != nil {
		responce.FailWithMessage("Add db error", c)
		return
	}
	responce.Ok(c)
}

func DeleteProject(c *gin.Context) {
	projectId := c.Param("id")
	if err := model.DeleteProject(projectId); err != nil {
		responce.FailWithMessage("Delete db error", c)
		return
	}
	responce.Ok(c)
}
