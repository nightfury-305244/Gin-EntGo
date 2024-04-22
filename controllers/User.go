package controller

import (
	"log"
	"main/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func CreateUser(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("user Post test")

		var user User
		if ctx.ShouldBind(&user) == nil {
			log.Println(user.Name)
			log.Println(user.Email)
			log.Println(user.Password)
		}

		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := client.User.
			Create().
			SetName(user.Name).
			SetEmail(user.Email).
			SetPassword(user.Password).
			Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"failed creating user": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, u)
	}
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func GetJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, json)
	}
}
