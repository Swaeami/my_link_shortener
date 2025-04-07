package main

import (
	"my_link_shortener/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	urlRepository := repository.NewInMemoryUrlRepository()
	r.GET("/:short", func(c *gin.Context) {
		short := c.Param("short")
		url, err := urlRepository.GetByShort(short)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, url.Original)
	})
	r.POST("/doshort", func(c *gin.Context) {
		var json struct {
			Original string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		url, err := urlRepository.DoShort(json.Original)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"short": url.Short})
	})
	r.Run()
}
