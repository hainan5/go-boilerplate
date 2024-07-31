// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/app/http/controllers"
	"go-boilerplate/app/http/controllers/auth"
	"go-boilerplate/app/http/middlewares"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	var v1 *gin.RouterGroup
	v1 = r.Group("/api")
	coc := new(controllers.CommonController)
	v1.GET("/ping", coc.SayHello)

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)

			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			uc := new(controllers.UsersController)

			// 获取当前用户
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
				usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
				usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
				usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
				usersGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
				usersGroup.PUT("/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
			}
		}
	}
}
