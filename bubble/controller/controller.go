package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandller(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//从请求中拿出值
	var todo models.Todo
	c.BindJSON(&todo)
	//存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func GetTodoList(c *gin.Context) {
	//查询表里的所有数据
	todoList,err:=models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}else {
		c.JSON(http.StatusOK, todoList)
	}

}
func UpadateATodo(c *gin.Context) {
	//传一个值过来
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "错误的ID"})
		return
	}
	todo,err :=models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.BindJSON(&todo)
	if err = models.UpadateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "错误的ID"})
		return
	}
	if err:=models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}
