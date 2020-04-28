package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var baseURL string

func main() {
	// init
	listenOn := os.Getenv("TVPROXY_LISTEN")
	if listenOn == "" {
		listenOn = "127.0.0.1:10086"
	}
	baseURL = os.Getenv("TVPROXY_BASE_URL")
	if baseURL == "" {
		baseURL = "http://" + listenOn + "/"
	}
	// webserver
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/i.ts", tsProxyHandler)
	tvb := r.Group("/tvb")
	{
		tvb.GET("/inews.m3u8", iNewsHandler)
		tvb.GET("/finance.m3u8", financeHandler)
	}
	rthk := r.Group("/rthk")
	{
		rthk.GET("/31.m3u8", rthk31Handler)
		rthk.GET("/32.m3u8", rthk32Handler)
	}
	r.Run(listenOn)
}
