package middleware

import (
	"blog_server/common"
	"blog_server/model"
	"blog_server/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		// token为空
		if tokenString == "" {
			response.Fail(c, nil, "权限不足")
			c.Abort()
			return
		}
		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			response.Fail(c, nil, "权限不足")
			c.Abort()
			return
		}
		// 提取token的有效部分
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		// 非法token
		if err != nil || !token.Valid {
			response.Fail(c, nil, "权限不足")
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id =?", userId).First(&user)
		if user.Identify != "系统管理员" {
			response.Fail(c, nil, "权限不足")
			c.Abort()
			return
		}
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
