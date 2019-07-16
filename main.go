package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"image"
	_ "image/jpeg"
	//Don't forgot to add import _ "image/jpeg" the image package itself doesn't know how to decode jpeg, you need to import image/jpeg to register the jpeg decoder.
)

func main() {
	r := gin.Default()
	r.POST("/load", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil{
			fmt.Println(err)
		}
		c.SaveUploadedFile(file, "/tmp/123.jpg")
		c.String(http.StatusOK, fmt.Sprintf("%s is uploaded!", file.Filename))
	})
	r.POST("/delete", func(c *gin.Context) {
		path := c.PostForm("path")
		fail := os.Remove(path)
		if fail == nil {
			fmt.Println("file is deleted!")
		} else {
			fmt.Println("Cannot delete file!")
			fmt.Println("%s", fail)
		}
		c.String(http.StatusOK, fmt.Sprintf("file is deleted!"))
	})
	r.GET("/imageinfo", func(c *gin.Context) {
		imagepath, _ := c.GetQuery("imagepath")
		fmt.Println(imagepath)
		file, err := os.Open(imagepath)
		if err != nil{
			fmt.Println(err)
		}
		image, _, err := image.DecodeConfig(file)
		if err != nil{
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"width":image.Width,
			"height":image.Height,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
