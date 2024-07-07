package main

import (
	"bubble/gormcli"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() {
	DB = gormcli.GetDB()
}

func GetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	initMySQL()

	// 模型绑定(其实就是绑定表)
	err := DB.AutoMigrate(&Todo{})
	if err != nil {
		panic("DB migrate err")
	}

	r := gin.Default()
	// 告诉gin模板文件引用的静态文件去哪找
	r.Static("/static", "static")
	// 告诉gin去哪找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", GetHandler)

	// v1分组的api
	v1Group := r.Group("v1")
	{
		// 待办事项

		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项 点击提交 会发送请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)

			// 2. 存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK, gin.H{
				//	"code": 2000,
				//	"msg": "success",
				//	"data": todo,
				//})
			}

			// 3. 返回响应
		})

		// 查看
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todos表中的所有数据
			var todolist []Todo
			if err = DB.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todolist)
			}
		})

		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})

		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})

		// 删除
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			if err = DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	if err := r.Run(); err != nil {
		panic("r.Run() failed.")
	}
}
