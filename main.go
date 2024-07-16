package main

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int `json:"id"`
	SecretField int `json:"secret_field"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		// https://dummyjson.com/products/2
		
		req, err := http.NewRequest("GET", "https://dummyjson.com/products/2", nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}

		req.Header.Set("User-Agent", "curl/7.64.1")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Data(http.StatusOK, "application/json", body)
	})
	
	r.GET("/ping", func(c *gin.Context) {
		user := User{
			ID:          1,
			SecretField: 2,
		}

		user.SecretField = 23

		c.JSON(http.StatusOK, user)
	})

	r.GET("/bin", func(c *gin.Context) {
		mbSize, _ := strconv.Atoi(c.Query("mb"))
		if mbSize == 0 || mbSize > 1000 {
			mbSize = 1
		}

		mbToBytes := mbSize * 1024 * 1024

		bytes := &bytes.Buffer{}

		for i := 0; i < mbToBytes; i++ {
			bytes.WriteRune('a')
		}

		c.Data(http.StatusOK, "application/octet-stream", bytes.Bytes())
	})

	r.Run(":8000")
}
