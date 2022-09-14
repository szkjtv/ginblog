package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// gin session key
const KEY = "AEN233"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

// session中间件 AuthSessionMiddle
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			// c.HTML(200, "index.html", gin.H{
			// 	"notice": "请登录后再访问",
			// })
			// c.JSON(200, "你无权访问,请登录后再访问,我是负责拦你的管事员")
			//c.String(200,"你无权访问,请登录后再访问!!")
			c.HTML(200, "notlogged.html", gin.H{
				"notlogin": "对不起，您请求的页面不存在、或已被删除、或暂时不可用",
			})
			c.Abort()
			return
		}
		// 设置简单的变量
		c.Set("userId", sessionValue.(uint))

		c.Next()
		return
	}
}

// 注册和登陆时都需要保存seesion信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)

}

// 注册和登陆时都需要保存seesion信息
func SaveAuthSessionUserName(c *gin.Context, id uint, account string) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Set("username", account)
	session.Save()
	// return
}

func GetSessionUserAccunt(c *gin.Context) (id int, account string) {
	session := sessions.Default(c)
	sessionValue := session.Get("account")
	if sessionValue == nil {
		c.JSON(200, "请求错误返回的数据")
		return 0, ""
	}
	return 0, sessionValue.(string)

}

// func GetSessionUserAcoount(c *gin.Context) {
// 	session := sessions.Default(c)

// 	userInfo = session.Get("account").(mysql.LoginUser)

// 	// if userInfo != userInfo {
// 	// 	return
// 	// }
// 	return userInfo

// }
