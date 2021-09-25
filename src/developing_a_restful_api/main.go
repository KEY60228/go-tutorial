package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// https://pkg.go.dev/github.com/gin-gonic/gin#Default
	// 初期化
	router := gin.Default()
	// https://pkg.go.dev/github.com/gin-gonic/gin#RouterGroup.GET
	// ルート定義
	router.GET("/albums", getAlbums)

	// https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run
	// http.Server起動
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	// https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON
	// 第一引数でステータスコード指定、第二引数でJSON化するデータを渡す
	// Content-Type: application/jsonのレスポンスを返す？
	c.IndentedJSON(http.StatusOK, albums)
}
