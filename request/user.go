package request

type UserReq struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"username" binding:"required"`
}
