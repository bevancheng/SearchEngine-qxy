package handlers

import (
	"userser/cmd/api/rpc"
	"userser/kitex_gen/user"
	"userser/pkg/errno"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var signupVar UserParam
	if err := c.ShouldBind(&signupVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(signupVar.UserName) == 0 || len(signupVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(c, &user.UserSignUpRequest{
		Username: signupVar.UserName,
		Password: signupVar.PassWord,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)

}
