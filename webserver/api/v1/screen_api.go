package v1

import (
	"fmt"

	"github.com/edgehook/apphub-dashboard-server/common/dbm/model"
	v1 "github.com/edgehook/apphub-dashboard-server/common/types/v1"
	"github.com/edgehook/apphub-dashboard-server/common/utils"
	responce "github.com/edgehook/apphub-dashboard-server/webserver/types"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func GetScreens(c *gin.Context) {
	var (
		err     error
		screens []*model.Screen
		count   int64
	)
	keywords := c.Query("keywords")
	screens, err = model.GetScreenByKeywords(keywords)
	if err != nil {
		responce.FailWithMessage("Get dashboard error", c)
		return
	}
	count, err = model.GetScreenCount(keywords)
	if err != nil {
		responce.FailWithMessage("Get count error", c)
		return
	}
	data := make([]*model.Screen, 0)
	for _, screen := range screens {
		screen.Data = ""
		data = append(data, screen)
	}

	responce.OkWithData(map[string]interface{}{
		"list":  data,
		"total": count,
	}, c)
}

func GetScreensByIsAside(c *gin.Context) {
	var (
		err     error
		screens []*model.Screen
	)
	isAside := c.Query("isAside")
	screens, err = model.GetScreenByIsAside(isAside)
	if err != nil {
		responce.FailWithMessage("Get dashboard error", c)
		return
	}
	data := make([]*model.Screen, 0)
	for _, screen := range screens {
		screen.Data = ""
		data = append(data, screen)
	}

	responce.OkWithData(map[string]interface{}{
		"list":  data,
		"total": -1,
	}, c)
}

func GetScreenById(c *gin.Context) {
	screenId := c.Param("id")
	if screenId == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	screen, err := model.GetScreenById(screenId)
	if err != nil {
		responce.FailWithMessage("Get dashboard error", c)
		return
	}
	responce.OkWithData(screen, c)
}

func UpdateScreen(c *gin.Context) {
	var (
		screenData v1.ScreenData
	)
	screenId := c.Param("id")
	if screenId == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if err := c.Bind(&screenData); err != nil {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if err := model.SaveScreen(screenId, screenData.Name, screenData.State, screenData.Data, screenData.Image, screenData.IsAside); err != nil {
		responce.FailWithMessage("Update dashboard error", c)
		return
	}

	responce.Ok(c)
}

func AddScreen(c *gin.Context) {
	var (
		screenData v1.ScreenData
	)

	if err := c.Bind(&screenData); err != nil {
		klog.Errorln("Parameter error: %v", err.Error())
		responce.FailWithMessage("Parameter error", c)
		return
	}
	if isExist := model.IsExistScreenByName(screenData.Name); isExist {
		responce.FailWithMessage("Dashboard already exists", c)
		return
	}
	state := 0
	screenId := utils.NewUUID()
	err := model.AddScreen(&model.Screen{
		Id:          screenId,
		Name:        screenData.Name,
		Description: screenData.Description,
		State:       state,
		IsAside:     "false",
	})
	if err != nil {
		responce.FailWithMessage("Add db error", c)
		return
	}
	responce.OkWithData(map[string]interface{}{
		"id": screenId,
	}, c)
}

func CopyScreen(c *gin.Context) {
	screenId := c.Param("id")
	screen, err := model.GetScreenById(screenId)
	if err != nil {
		responce.FailWithMessage("Dashboard does not exist", c)
		return
	}
	name := screen.Name
Retry:
	name = fmt.Sprintf("%s-%s", name, "copy")
	if isExist := model.IsExistScreenByName(name); isExist {
		klog.V(4).Infof("Screen name: %s already exists", name)
		goto Retry
	}
	state := 0
	id := utils.NewUUID()
	if err := model.AddScreen(&model.Screen{
		Id:          id,
		Name:        name,
		Description: screen.Description,
		State:       state,
		Data:        screen.Data,
		IsAside:     screen.IsAside,
	}); err != nil {
		responce.FailWithMessage("Copy dashboard error", c)
		return
	}
	responce.OkWithData(map[string]interface{}{
		"id": screenId,
	}, c)
}

func DeleteScreen(c *gin.Context) {
	screenId := c.Param("id")
	if err := model.DeleteScreen(screenId); err != nil {
		responce.FailWithMessage("Delete db error", c)
		return
	}

	responce.Ok(c)
}
