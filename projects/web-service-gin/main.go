package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album 代表一个专辑的数据
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

// getAlbums 使用专辑列表作为JSON响应
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums 添加一个album 从接受的JSON请求体中
func postAlbums(c *gin.Context) {
	var newAlbum album
	// 调用BindJSON去绑定接受的JSON，变成一个newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 添加新的album到切片中
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID 定位于客户端发送请求到id匹配的album，然后返回album作为响应
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// 遍历albums列表，查找与参数ID匹配的album
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album not found"})
}
