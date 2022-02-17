package controller

import (
	"blogspot/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type POSTINPUT struct {
	Title string `json:"title" validation:"required,min:20"`
	Content string `json:"content" validation:"required,min:200"`
	Category string `json:"category" validation:"required,min:3" `
	Status string `json:"status"   validation:"required"`
}

func GetPostAll(c *gin.Context){
	db:=c.MustGet("db").(*gorm.DB)
	var Post []model.Post
	db.Find(&Post)

	c.JSON(http.StatusOK,gin.H{"data":Post})
}
func CreatePostData(c *gin.Context){
	var input POSTINPUT
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return  
	}
	Post:=model.Post{
		Title: input.Title,
		Content: input.Content,
		Category: input.Category,
		Status: input.Status,
	}
	db:=c.MustGet("db").(*gorm.DB)
	db.Create(Post)

	c.JSON(http.StatusOK,gin.H{"data":Post})
}
func GetById(c *gin.Context){
	var Post model.Post
	db:=c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Post).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan !!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": Post})
}
func UpdateDataId(c *gin.Context){
	var Post model.Post
	var DataPost POSTINPUT

	db := c.MustGet("db").(*gorm.DB)

    if err := db.Where("id = ?", c.Param("id")).First(&Post).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }
	if err := c.ShouldBindJSON(&DataPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	var datachanged model.Post
	datachanged.Category=DataPost.Category
	datachanged.Content=DataPost.Content
	datachanged.Title=DataPost.Title
	datachanged.Status=DataPost.Status

	db.Model(&Post).Updates(datachanged)
	

	
    c.JSON(http.StatusOK, gin.H{"data": Post})
}

func DeletebyID(c *gin.Context){
	   // Get model if exist
	   db := c.MustGet("db").(*gorm.DB)
	   var POST model.Post
	   if err := db.Where("id = ?", c.Param("id")).First(&POST).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		   return
	   }
   
	   db.Delete(&POST)
   
	   c.JSON(http.StatusOK, gin.H{"data": true})
}