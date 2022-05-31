package main

import (
	"context"
	"net/http"
	"time"
	"userser/cmd/api/handlers"
	"userser/cmd/api/rpc"
	"userser/kitex_gen/user"
	"userser/pkg/constants"
	"userser/pkg/tracer"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}
func main() {
	Init()

	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data any) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}

		},
		Authenticator: func(c *gin.Context) (any, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &user.UserLogInRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/signup", handlers.SignUp)
	user1.POST("/logout", authMiddleware.LogoutHandler)

	favorite1 := v1.Group("/like")
	favorite1.Use(authMiddleware.MiddlewareFunc())

	//todo

	//searcher1 := v1.Group("/searcher")
	//searcher1.POST("/index") //todo of searcher
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
