package models

import (
	"bubble/dao"
)

//Todo model
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Statule bool   `json:"statule"`
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}
func GetAllTodo() (todoList *[]Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
func UpadateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}
func DeleteATodo(id string)(err error)  {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}