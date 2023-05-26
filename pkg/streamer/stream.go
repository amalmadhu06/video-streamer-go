package streamer

import (
	"fmt"
	"net/http"
	"strings"
)

func SegmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url:", r.URL, "  method:", r.Method)

	urlParts := strings.Split(r.URL.Path, "/")
	fmt.Println("url parts:", urlParts)

	if len(urlParts) < 2 {
		http.Error(w, "Video ID not found in the request URL", http.StatusBadRequest)
		return
	}

	videoID := urlParts[len(urlParts)-1]
	fmt.Println("video ID:", videoID)

	segmentsDir := "./storage/" + videoID
	fileServer := http.FileServer(http.Dir(segmentsDir))
	http.Handle("/"+videoID+"/", http.StripPrefix("/"+videoID+"/", fileServer))
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url --> ", r.URL, "method -->", r.Method)
	urlParts := strings.Split(r.URL.Path, "/")
	fmt.Println("url parts:", urlParts)

	if len(urlParts) < 2 {
		http.Error(w, "Video ID not found in the request URL", http.StatusBadRequest)
		return
	}

	videoID := urlParts[len(urlParts)-1]
	fmt.Println("video ID:", videoID)

	segmentsDir := "./storage/" + videoID
	playlistFile := segmentsDir + "/playlist.m3u8"

	w.Header().Set("Content-Type", "application/x-mpegURL")
	http.ServeFile(w, r, playlistFile)
}
