package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouters()*gin.Engine  {
	r := gin.Default()
	r.Static("/static", "static") //引用静态文件
	r.LoadHTMLGlob("template/*")  //寻找index文件
	r.GET("/", controller.IndexHandller)
	//待办事项
	v1Group := r.Group("v1")
	{
		//	待办事项
		//	增加
		v1Group.POST("/todo",controller.CreateTodo)
		//	查看 1.查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		//	2.查看单个
		v1Group.GET("/todo/:id", )
		//	修改
		v1Group.PUT("/todo/:id",controller.UpadateATodo)
		//	删除
		v1Group.DELETE("/todo/:id",controller.DeleteATodo)
	}
	return r
}
