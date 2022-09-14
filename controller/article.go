package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"example.com/m/session"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//获取添加文章页面
func Aritcle(c *gin.Context) {
	db := Dbinit()
	var user LoginUser
	UserID := int(session.GetSessionUserId(c))
	db.First(&user, UserID).Find(&user)
	c.HTML(200, "index.html", gin.H{
		"username": user.Username,
		"account":  user.Account,
	})
}

// 添加文章
func PostAricle(c *gin.Context) {
	db := Dbinit()
	title := c.PostForm("title")
	content := c.PostForm("content")
	UserID := int(session.GetSessionUserId(c)) //用户sesion
	new := Aritclee{Title: title, Content: content, UserID: UserID}
	db.Create(&new)
	c.Redirect(http.StatusMovedPermanently, "/user/auth/admin")
	defer db.Close()

}

//查询文章 后台账号查询的文章内容  需要登录
func QueryAricle(c *gin.Context) {

	db := Dbinit()
	var article []Aritclee
	var seacrharticle []Aritclee
	var user LoginUser
	query := c.Query("title")

	UserID := int(session.GetSessionUserId(c))
	db.First(&user, UserID).Find(&user)
	currentPage, err := strconv.Atoi(c.Query("currentPage")) //第几页
	if err != nil {
		currentPage = 1 //第几页
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize")) //一页显示多少条
	if err != nil {
		pageSize = 15 //默认一页显示多少条
	}

	var total int = 0 //总条数据
	db.Order("id desc").Where("title LIKE  ?", "%"+query+"%").Where(&Aritclee{UserID: UserID}).Model(&seacrharticle).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&article)

	a := total/15 + 1
	renovate := "/user/auth/query?&pageSize=15&currentPage=1"
	if currentPage > a {
		c.HTML(200, "artilcle.html", gin.H{
			"datanull":    "没有数据了,不要再点了！！！",
			"renovate":    renovate,
			"currentPage": 0,
		})
		return
	}

	if currentPage == 0 {
		c.HTML(200, "artilcle.html", gin.H{
			"zuosi":       "不要走回头路",
			"renovate":    renovate,
			"currentPage": 0,
		})
		return
	}

	c.HTML(200, "artilcle.html", gin.H{
		"relut":       article,
		"searult":     seacrharticle,
		"query":       query,           //搜索出来的内容,可以拼接在上一页下一页
		"add":         currentPage + 1, //下一页加1
		"reduce":      currentPage - 1, //上一页减1
		"pagenumber":  a,               //共多少页
		"total":       total,           //总条数
		"currentPage": currentPage,     //前端显示在第几页
		"renovate":    renovate,        //刷新
		"username":    user.Username,
		"account":     user.Account,
	})

	defer db.Close()

}

// 通过id删除删除数据
func DeleAriticle(c *gin.Context) {
	db := Dbinit()
	id := c.Param("id")
	var delete Aritclee
	db.Where(id).Delete(&delete) //调试前端时添加了id
	// db.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/user/auth/query")
	// c.JSON(200, "删除成功")
}

// 修改更新文章
func Uparticle(c *gin.Context) {
	db := Dbinit()
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	db.Model(&Aritclee{}).Where(id).Update(&Aritclee{Title: title, Content: content})
	c.Redirect(http.StatusMovedPermanently, "/user/auth/query")
	defer db.Close()
}

// 显示文章回显修改
func ShowEditor(c *gin.Context) {
	db := Dbinit()
	var querarticle Aritclee
	var user LoginUser
	UserID := int(session.GetSessionUserId(c))
	db.First(&user, UserID).Find(&user)
	id := c.Param("id")
	db.Find(&querarticle, id)
	c.HTML(http.StatusOK, "uparticle.html", gin.H{
		"ID":       querarticle.ID,
		"title":    querarticle.Title,
		"content":  querarticle.Content,
		"username": user.Username,
		"account":  user.Account,
	})
}

// 查看详情
func Deltel(c *gin.Context) {
	db := Dbinit()
	var detile Aritclee
	var auth LoginUser
	id := c.Param("id")

	db.Find(&detile, id).First(&auth, detile.UserID)
	c.HTML(http.StatusOK, "detel.html", gin.H{
		"dlet":    detile,
		"auth":    auth.Username,
		"content": template.HTML(detile.Content),
	})
}

// 不需要登录就可以查看和搜索的，不用管这个接口
// 用户查询文章 不需要用户所有人可以查询的文章内容
//查询文章
func UserQueryAricle(c *gin.Context) {

	db := Dbinit()
	var article []Aritclee
	query := c.Query("title")

	// var u Aritclee
	// db.First(&u)
	// db.Select([]string{"username", "account"}).Model(&u).Find(&u.User, &article)
	// fmt.Println("查询的结果:", u)
	// c.JSON(200, gin.H{
	// 	"u":        u,
	// 	"ariticle": article,
	// })
	currentPage, err := strconv.Atoi(c.Query("currentPage")) //第几页
	if err != nil {
		currentPage = 1 //第几页
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize")) //一页显示多少条
	if err != nil {
		pageSize = 15 //默认一页显示多少条
	}

	var total int = 0 //总条数据
	db.Order("id desc").Where("title LIKE  ?", "%"+query+"%").Model(&article).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&article)

	a := total/15 + 1
	renovate := "/"
	if currentPage > a {
		c.HTML(200, "blog.html", gin.H{
			"datanull":    "没有数据了,不要再点了！！！",
			"renovate":    renovate,
			"currentPage": 0,
		})
		return
	}

	if currentPage == 0 {
		c.HTML(200, "blog.html", gin.H{
			"zuosi":       "不要走回头路",
			"renovate":    renovate, //刷新
			"currentPage": 0,
		})
		return
	}

	c.HTML(200, "blog.html", gin.H{
		"relut": article,
		// "searult":     seacrharticle,
		"query":       query,           //搜索出来的内容,可以拼接在上一页下一页
		"add":         currentPage + 1, //下一页加1
		"reduce":      currentPage - 1, //上一页减1
		"pagenumber":  a,               //共多少页
		"total":       total,           //总条数
		"currentPage": currentPage,     //前端显示在第几页
		"renovate":    renovate,        //刷新

	})

	defer db.Close()

}
