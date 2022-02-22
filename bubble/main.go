package main

import (
	"bubble/dao"
	"bubble/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Todo model
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Statule bool   `json:"statule"`
}


func main() {
	//创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&Todo{})
	//gin框架启动
	r:=routers.SetupRouters()
	r.Run(":8080")
}
