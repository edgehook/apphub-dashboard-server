package v1

import (
	"github.com/edgehook/apphub-dashboard-server/common/dbm/model"
	v1 "github.com/edgehook/apphub-dashboard-server/common/types/v1"
	"github.com/edgehook/apphub-dashboard-server/common/utils"
	responce "github.com/edgehook/apphub-dashboard-server/webserver/types"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func GetScreenImagesByScreenIdAndType(c *gin.Context) {
	screenId := c.Query("screenId")
	ctype := c.Query("type")
	if screenId == "" || ctype == "" {
		responce.FailWithMessage("Parameter error", c)
		return
	}
	screenImages, err := model.GetScreenImagesByScreenIdAndType(screenId, ctype)
	if err != nil {
		responce.FailWithMessage("Get screen error", c)
		return
	}
	responce.OkWithData(screenImages, c)
}

func AddScreenImage(c *gin.Context) {
	var (
		screenImageData v1.ScreenImageData
	)
	if err := c.Bind(&screenImageData); err != nil {
		klog.Errorln("Parameter error: %v", err.Error())
		responce.FailWithMessage("Parameter error", c)
		return
	}
	screenImages, _ := model.GetScreenImagesByScreenIdAndType(screenImageData.ScreenId, screenImageData.Type)
	if screenImageData.Type == v1.AppHub_Dashboard_Image_Type && len(screenImages) > 0 {
		screenImage := screenImages[0]
		err := model.SaveScreenImage(screenImage.Id, screenImageData.Data)
		if err != nil {
			responce.FailWithMessage("Save db error", c)
			return
		}
	} else {
		screenImageId := utils.NewUUID()
		err := model.AddScreenImage(&model.ScreenImage{
			Id:       screenImageId,
			ScreenId: screenImageData.ScreenId,
			Type:     screenImageData.Type,
			Data:     screenImageData.Data,
			Filename: screenImageData.Filename,
		})
		if err != nil {
			responce.FailWithMessage("Add db error", c)
			return
		}
	}

	responce.Ok(c)
}

func DeleteScreenImage(c *gin.Context) {
	screenImageId := c.Param("id")
	if err := model.DeleteScreenImage(screenImageId); err != nil {
		responce.FailWithMessage("Delete db error", c)
		return
	}
	responce.Ok(c)
}
