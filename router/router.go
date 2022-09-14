package router

import (

	//跨域

	"example.com/m/controller"
	"example.com/m/session"
	"github.com/gin-gonic/gin"
)

func TTest(c *gin.Context) {
	c.HTML(200, "TTest.html", nil)
}

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.Static("/static", "./static")
	r.Any("/ueditor/controller", controller.Action) // 富文本
	// r.GET("/admin", controller.Aritcle) //获取后台添加文章页面
	// r.GET("/query", controller.QueryAricle) //查询文章
	// r.GET("/delec/:id", controller.DeleAriticle)  //删除接口
	// r.POST("/ar", controller.PostAricle)  //添加文章
	// r.POST("/upaticle/:id", controller.Uparticle) //文章修改接口
	// r.GET("/showeditor/:id", controller.ShowEditor) //显示修改文章页面获取要修改的文章
	r.GET("/deltel/:id", controller.Deltel)                 //普通用户和登录用户都可以查看详情的接口
	r.GET("/", controller.UserQueryAricle)                  //不需要登录就可以查询搜索的接口
	r.GET("/tt", TTest)                                     //测试页面测试好后要删除的接口
	r.GET("/adminarticle", controller.AdminQueryAricle)     //管理员查看所有文章内容和删除
	r.GET("/admindelete/:id", controller.AdminDeleAriticle) //管理员删除接口

	u := r.Group("/user", session.EnableCookieSession())
	{
		u.Static("/static", "./static")
		u.GET("/register", controller.GetResiter) //获取登录注册页面
		u.POST("/lopost", controller.Login)       //登录
		u.POST("/repost", controller.Register)    //注册
		//登录后才能访问的路由
		f := u.Group("/auth", session.MiddleWare())
		{

			f.GET("/admin", controller.Aritcle)             //后台首页添加页面
			f.POST("/ar", controller.PostAricle)            //登录后添加文章接口
			f.GET("/query", controller.QueryAricle)         //登录后的用户查询文章
			f.GET("/logout", controller.Logout)             //退出登录
			f.POST("/upaticle/:id", controller.Uparticle)   //修改接口
			f.GET("/showeditor/:id", controller.ShowEditor) //显示修改文章的页面
			f.GET("/delec/:id", controller.DeleAriticle)    //删除文章接口
		}
	}

	r.Run(":81")

}
