package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong %s\n %s\n %s\n %s\n %s\n %s\n %s\n",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,/Users/xqzhang/Desktop/codes/helloworld/go_gin>Finished running tool: /Users/xqzhang/go/bin/staticcheck",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,/Users/xqzhang/Desktop/codes/helloworld/go_gin>Finished running tool: /Users/xqzhang/go/bin/staticcheck",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,/Users/xqzhang/Desktop/codes/helloworld/go_gin>Finished running tool: /Users/xqzhang/go/bin/staticcheck",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,/Users/xqzhang/Desktop/codes/helloworld/go_gin>Finished running tool: /Users/xqzhang/go/bin/staticcheck",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,func (g *gzipHandler) shouldCompress(req *http.Request) bool {",
			"Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,Starting linting the current package at /Users/xqzhang/Desktop/codes/helloworld/go_gin,func (g *gzipHandler) shouldCompress(req *http.Request) bool {",
			fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
