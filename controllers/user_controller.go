package controllers

import (
	"day8/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserController struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewUserController(mySQL *gorm.DB, redisCli *redis.Client) UserController {
	return UserController{
		DB:    mySQL,    // Gán DB
		Redis: redisCli, // Gán Redis
	}
}

// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (u UserController) GetUsers(c *gin.Context) {
	var users []models.User
	tx := u.DB.Preload("Posts").Find(&users) // Preload Posts
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func (u UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// @Summary Get a user by ID
// @Description Get a specific user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (u UserController) GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Sử dụng Preload để lấy bài viết của người dùng
	if err := u.DB.Preload("Posts").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Update a user by ID
// @Description Update a specific user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (u UserController) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := u.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Delete a user by ID
// @Description Delete a specific user by ID
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Router /users/{id} [delete]
func (u UserController) DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := u.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := u.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
