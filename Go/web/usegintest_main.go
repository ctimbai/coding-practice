// usegintest_main
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 2, Name: "王五"},
}

func main() {
	r := gin.Default()
	fmt.Println("hello gin")
	r.GET("/users", listUser)
	r.Run(":8080")
}

func listUser(c *gin.Context) {
	c.JSON(200, users)
}
