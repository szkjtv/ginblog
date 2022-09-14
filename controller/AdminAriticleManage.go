package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//查询文章 后台账号查询的文章内容  需要登录
func AdminQueryAricle(c *gin.Context) {

	db := Dbinit()
	var article []Aritclee
	var seacrharticle []Aritclee

	query := c.Query("title")

	currentPage, err := strconv.Atoi(c.Query("currentPage")) //第几页
	if err != nil {
		currentPage = 1 //第几页
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize")) //一页显示多少条
	if err != nil {
		pageSize = 15 //默认一页显示多少条
	}

	var total int = 0 //总条数据
	db.Order("id desc").Where("title LIKE  ?", "%"+query+"%").Model(&seacrharticle).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&article)

	a := total/15 + 1
	renovate := "/adminarticle?&pageSize=15&currentPage=1"
	if currentPage > a {
		c.HTML(200, "AdminAriticleManage.html", gin.H{
			"datanull":    "没有数据了,不要再点了！！！",
			"renovate":    renovate,
			"currentPage": 0,
		})
		return
	}

	if currentPage == 0 {
		c.HTML(200, "AdminAriticleManage.html", gin.H{
			"zuosi":       "不要走回头路",
			"renovate":    renovate,
			"currentPage": 0,
		})
		return
	}

	c.HTML(200, "AdminAriticleManage.html", gin.H{
		"relut":       article,
		"searult":     seacrharticle,
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

// 通过id删除删除数据
func AdminDeleAriticle(c *gin.Context) {
	// db := Dbinit()
	// id := c.Param("id")
	// var delete Aritclee
	// db.Where(id).Delete(&delete) //调试前端时添加了id
	// // db.Delete(id)
	// c.Redirect(http.StatusMovedPermanently, "/adminarticle")
	// // c.JSON(200, "删除成功")

	db := Dbinit()
	id := c.Param("id")
	var delete Aritclee
	db.Where(id).Delete(&delete) //调试前端时添加了id
	// db.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/adminarticle")
}
