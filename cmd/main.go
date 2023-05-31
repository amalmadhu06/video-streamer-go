package main

import (
	"github.com/amalmadhu06/video-streamer-go/pkg/streamer"
	"github.com/amalmadhu06/video-streamer-go/pkg/uploader"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// route for uploading video to storage
	r.POST("/upload", uploader.Upload)

	// route for streaming videos using hls
	r.GET("/play/:video_id/:playlist", streamer.Stream)

	r.Run()
}
