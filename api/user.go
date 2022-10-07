package api

import (
	"go_web/request"
	"go_web/service"
	"net/http"

	logging "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userRegister request.UserReq
	err := ctx.ShouldBind(&userRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "json数据解析失败",
			"error": err.Error(),
		})
		logging.Info(err)
		return
	}

	res, err := service.UserRegister(userRegister)
	ctx.JSON(res.Code, res)
}

func Login(ctx *gin.Context) {
	var userLogin request.UserReq
	err := ctx.ShouldBind(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "json数据解析失败",
			"error": err.Error(),
		})
		logging.Info(err)
		return
	}

	res, err := service.UserLogin(userLogin)
	ctx.JSON(res.Code, res)
}
