package controller

import (
	"math/rand"
	"net/http"
	"time"

	"example.com/m/session"
	"github.com/gin-gonic/gin"
)

func GetResiter(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqlstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}
	return string(result)

}

func Register(c *gin.Context) {
	db := Dbinit()
	var loginuser LoginUser

	//loginuser.Account = c.Request.FormValue("account")
	//loginuser.Password = c.Request.FormValue("password")
	loginuser.Username = c.PostForm("username")
	loginuser.Account = c.PostForm("account")
	loginuser.Password = c.PostForm("password")

	// if loginuser.Username == "" {
	// 	c.JSON(202, "用户名不能为空")
	// 	return
	// }

	//如果用户名为空,生成随机字符串
	if len(loginuser.Username) == 0 {
		loginuser.Username = RandomString(10)
	}

	if len(loginuser.Account) == 0 {
		c.HTML(200, "login.html", gin.H{
			"acconull": "账号不能为空",
		})
		return
	}

	if len(loginuser.Account) < 11 {
		c.HTML(200, "login.html", gin.H{
			"numbererro": "您输入的手机号码有误",
		})
		return
	}
	if len(loginuser.Password) == 0 {
		c.HTML(200, "login.html", gin.H{
			"passnull": "密码不能为空",
		})
		return
	}

	if len(loginuser.Username) > 12 {
		c.HTML(200, "login.html", gin.H{
			"username": "用户名最大为12位",
		})
		return
	}

	if len(loginuser.Account) >= 12 {
		c.HTML(200, "login.html", gin.H{
			"phonenumbe": "您输入的不是11位数的手机号",
		})
		return
	}

	if len(loginuser.Password) < 6 {
		c.HTML(200, "login.html", gin.H{
			"passowrd": "密码不能小于6位",
		})
		return
	}

	// if hasSession := session.HasSession(c); hasSession == true {
	// 	c.JSON(200, "用户已登录")
	// 	return
	// }
	// db.Where("username = ?", loginuser.Username).First(&loginuser)
	// if loginuser.ID != 0 {
	// 	c.JSON(203, "用户名已存在请更换一个新的用户名")
	// 	return
	// }

	db.Where("account = ?", loginuser.Account).First(&loginuser)
	if loginuser.ID != 0 {
		c.HTML(200, "login.html", gin.H{
			"accoundz": "账号已存在请重新更换一个新账号",
		})

		return
	}

	if pwd, err := session.Encrypt(c.Request.FormValue("password")); err == nil {
		loginuser.Password = pwd
	}

	newAdduser := LoginUser{Username: loginuser.Username, Account: loginuser.Account, Password: loginuser.Password}
	db.Create(&newAdduser)

	session.SaveAuthSession(c, loginuser.ID)
	c.Redirect(http.StatusMovedPermanently, "/")
	// c.JSON(200, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	db := Dbinit()
	var user LoginUser
	// username := c.Request.FormValue("username")
	// account := c.Request.FormValue("account")
	// password := c.Request.FormValue("password")

	account := c.PostForm("account")
	password := c.PostForm("password")

	db.Where("account = ?", account).First(&user)
	if user.ID == 0 {
		c.HTML(200, "login.html", gin.H{
			"accountno": "账号不存在",
		})
		// c.JSON(203, "账号不存在,请核对账号")
		return
	}

	// if hasSession := session.HasSession(c); hasSession == true {
	// 	c.String(200, "用户登录过了")
	// 	// c.Redirect(http.StatusMovedPermanently, "/user/auth")
	// 	return
	// }

	//user := mysql.UserDetailByName(account)  //无效只能用这个判断来进行操作
	db.Where("password = ?", password).First(&user)

	if err := session.Compare(user.Password, password); err != nil {
		c.HTML(200, "login.html", gin.H{
			"passworderro": "密码错误",
		})
		return
	}

	session.SaveAuthSession(c, user.ID) //登录保存session信息
	session.SaveAuthSessionUserName(c, user.ID, user.Account)
	// c.JSON(200, gin.H{
	// 	"id:":  user.ID,
	// 	"用户名:": user.Username,
	// 	"账号:":  user.Account,
	// })

	// c.String(200, "登录成功")
	c.Redirect(http.StatusMovedPermanently, "/user/auth/admin")
}

func Logout(c *gin.Context) {
	if hasSession := session.HasSession(c); hasSession == false {
		c.String(401, "您还没登陆")
		return
	}
	session.ClearAuthSession(c) //清除session信息
	c.Redirect(http.StatusMovedPermanently, "/")
	c.String(200, "已经退出了")
}
