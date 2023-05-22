package v1

import (
	"github.com/edgehook/ithings/common/dbm/model"
	"github.com/edgehook/ithings/webserver/types"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"strconv"
)

type modelBody struct {
	Name         string `form: name json:"name  binding:"required"`
	ModelId      string `form: "modelId" json:"modelId"  binding:"required"`
	Manufacturer string `form:"manufacturer" json:"manufacturer"`
	Industry     string `form:"industry" json:"industry"`
	DataType     string `form:"dataType" json:"dataType"`
	Description  string `form:"description" json:"description"`
}

func GetModelsByPage(c *gin.Context) {
	var (
		err    error
		models []model.DeviceModel
	)
	keywords := c.Query("keywords")
	currentPage := c.Query("currentPage")
	limit := c.Query("limit")
	klog.Infof("page: %s, limit: %s", currentPage, limit)

	pageInt, err := strconv.Atoi(currentPage)
	if err != nil {
		types.FailWithMessage("Parameter error", c)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		types.FailWithMessage("Parameter error", c)
		return
	}
	if keywords == "" {
		models, err = model.GetModelByPage(pageInt, limitInt)
	} else {
		models, err = model.GetModelByPageAndKeywords(pageInt, limitInt, keywords)
	}

	if err != nil {
		types.FailWithMessage("Get models error", c)
		return
	}

	count, err := model.GetDeviceModelCount()
	if err != nil {
		types.FailWithMessage("Get count error", c)
		return
	}

	types.OkWithData(map[string]interface{}{
		"list":  models,
		"total": count,
	}, c)
}

func AddModels(c *gin.Context) {
	var modelData modelBody
	if err := c.Bind(&modelData); err != nil {
		types.FailWithMessage("Parameter error", c)
		return
	}

	err := model.AddDeviceModel(&model.DeviceModel{
		Name:         modelData.Name,
		ModelId:      modelData.ModelId,
		Manufacturer: modelData.Manufacturer,
		Industry:     modelData.Industry,
		Description:  modelData.Description,
	})
	if err != nil {
		types.FailWithMessage("Add db error", c)
		return
	}

	types.Ok(c)
}
