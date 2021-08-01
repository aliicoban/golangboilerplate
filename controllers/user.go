package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alicobanserver/models"
	"github.com/alicobanserver/utils"
	jwtgocobra "github.com/aliicoban/jwt-go-cobra"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.Users
	var usr = models.Users{}
	var error models.Error
	c.BindJSON(&user)
	if user.Email == "" {
		error.Message = "Email is missing"
		return
	}

	if user.Password == "" {
		error.Message = "Password missing"
		return
	}

	errUsr := utils.Usercollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&usr)
	if errUsr != nil {
		log.Fatal(errUsr)
	}

	if usr.Email != "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User already exist",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hash)
	user.Role = "user"

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := utils.Usercollection.InsertOne(ctx, user)

	if result != nil {
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
func Signin(c *gin.Context) {
	var user = models.Users{}
	c.BindJSON(&user)

	email := user.Email
	password := user.Password

	err := utils.Usercollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exist",
		})
		return
	}
	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Password mistake",
		})
		return
	}
	token, errToken := jwtgocobra.CreateToken(user.ID)
	if errToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
