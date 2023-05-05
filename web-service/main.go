package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"hello-go/web-service/controllers"
	"hello-go/web-service/database"
	"net/http"
)

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

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&database.DbUsername, "u", "root", "账号，默认为root")
	flag.StringVar(&database.DbPassword, "p", "root", "密码，默认为root")
	flag.StringVar(&database.DbHost, "h", "localhost", "主机名，默认为localhost")
	flag.StringVar(&database.DbPort, "P", "3306", "端口号，默认为3306")
	flag.StringVar(&database.DbName, "db", "my_db", "数据库，默认为my_db")
	// 解析命令行参数写入注册的flag里
	flag.Parse()

	// 输出结果
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbums)
	router.GET("/albums/:id", getAlbumById)
	userRepo := controllers.New()
	router.GET("/users", userRepo.GetUsers)
	router.POST("/users", userRepo.CreateUser)
	router.DELETE("/users/:id", userRepo.DeleteUser)
	router.Run("localhost:8080")
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func addAlbums(c *gin.Context) {
	newAlbum := new(album)
	if err := c.BindJSON(newAlbum); err != nil {
		return
	}

	albums = append(albums, *newAlbum)
	c.IndentedJSON(http.StatusCreated, *newAlbum)

}
