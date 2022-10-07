package service

import (
	"go_web/db/mysql"
	"go_web/model"
	"go_web/request"
	"go_web/response"
	"net/http"

	"gorm.io/gorm"
)

// 视图
func UserRegister(reg request.UserReq) (response.Response, error) {
	var user model.User
	err := mysql.MysqlDB.Where(&model.User{UserName: reg.UserName}).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		user.UserName = reg.UserName
		err := user.SetPassword(reg.Password)

		if err != nil {
			return response.Response{
				Code:  http.StatusInternalServerError,
				Msg:   "密码加密出错",
				Error: err.Error(),
			}, err
		}

		err = mysql.MysqlDB.Create(&user).Error

		if err != nil {
			return response.Response{
				Code:  http.StatusInternalServerError,
				Msg:   "数据库添加数据出错",
				Error: err.Error(),
			}, err
		}

		return response.Response{
			Code: http.StatusCreated,
			Msg:  "用户注册成功",
		}, nil

	} else if err != nil {
		return response.Response{
			Code:  http.StatusInternalServerError,
			Msg:   "数据库内部出错",
			Error: err.Error(),
		}, err
	}

	return response.Response{
		Code:  http.StatusForbidden,
		Msg:   "用户名重复，注册失败",
		Error: err.Error(),
	}, err
}

func UserLogin(login request.UserReq) (response.Response, error) {
	var user model.User
	err := mysql.MysqlDB.Where(&model.User{UserName: login.UserName}).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Response{
				Code:  http.StatusBadRequest,
				Msg:   "该用户不存在，请先注册",
				Error: err.Error(),
			}, err
		}

		// 不是用户不存在却还是出错，就是其他原因
		return response.Response{
			Code:  http.StatusInternalServerError,
			Msg:   "查询数据库出现错误",
			Error: err.Error(),
		}, err
	}

	// 用户从数据库中找到了，校验密码
	ok, err := user.CheckPassword(login.Password)
	if err != nil {
		return response.Response{
			Code:  http.StatusInternalServerError,
			Msg:   "登录失败 ",
			Error: err.Error(),
		}, err
	}

	if !ok {
		return response.Response{
			Code:  http.StatusInternalServerError,
			Msg:   "登录失败",
			Error: err.Error(),
		}, err
	}

	return response.Response{
		Code: http.StatusOK,
		Msg:  "登录成功",
		Data: map[string]string{"token": "token"},
	}, nil

}
