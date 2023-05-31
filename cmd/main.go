package main

import (
	"fmt"
	"github.com/amalmadhu06/video-streamer-go/pkg/streamer"
	"github.com/amalmadhu06/video-streamer-go/pkg/uploader"
	"net/http"
)

func main() {

	// route for uploading video to storage
	http.HandleFunc("/upload", uploader.Upload)

	//route for playing video with video_id
	http.HandleFunc("/play/", streamer.Play)

	http.HandleFunc("/hls/", streamer.ServePlaylistCheck)

	http.HandleFunc("/segment/", streamer.SegmentHandler)
	http.HandleFunc("/playlist/", streamer.PlaylistHandler)
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage/"))))

	fmt.Println("Server starting at port 3000.\nVisit http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("failed to start server\nerror: %q", err)
		return
	}
}
