package controllers

import (
	"day8/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type PostController struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewPostController(db *gorm.DB, redisCli *redis.Client) PostController {
	return PostController{DB: db, Redis: redisCli}
}

// @Summary Get all posts
// @Description Get a list of all posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func (p PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	tx := p.DB.Find(&posts)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// @Summary Create a new post
// @Description Create a new post
// @Accept json
// @Produce json
// @Param post body models.Post true "Post data"
// @Success 201 {object} models.Post
// @Router /posts [post]
func (p PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra xem user_id có hợp lệ không
	var user models.User
	if err := p.DB.First(&user, post.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := p.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// @Summary Get a post by ID
// @Description Get a specific post by ID
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func (p PostController) GetPostByID(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := p.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// @Summary Update a post by ID
// @Description Update a specific post by ID
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Post data"
// @Success 200 {object} models.Post
// @Router /posts/{id} [put]
func (p PostController) UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := p.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// @Summary Delete a post by ID
// @Description Delete a specific post by ID
// @Param id path int true "Post ID"
// @Success 204 "No Content"
// @Router /posts/{id} [delete]
func (p PostController) DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := p.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := p.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
