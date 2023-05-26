package main

import (
	"fmt"
	"github.com/amalmadhu06/video-streamer-go/pkg/streamer"
	"github.com/amalmadhu06/video-streamer-go/pkg/uploader"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploader.Upload)
	http.HandleFunc("/segment/", streamer.SegmentHandler)
	http.HandleFunc("/playlist/", streamer.PlaylistHandler)
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage/"))))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("failed to start server\nerror: %q", err)
		return
	}
	fmt.Println("Server started at port 3000.\nVisit http://localhost:3000/")
}
