//Package routes 注册路由
package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine){

	//测试一个 v1 路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发生验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}
	}

}