package streamer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SegmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[SegmentHandler] url:", r.URL, "  method:", r.Method)

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
	fmt.Println("[Playlist handler] url --> ", r.URL, "method -->", r.Method)
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

func Play(w http.ResponseWriter, r *http.Request) {

	fmt.Println("[func Play] url -->", r.URL, "method -->", r.Method)

	path := r.URL.Path
	// Remove the "/play/" prefix from the path
	videoID := strings.TrimPrefix(path, "/play/")

	// Use the videoID as needed
	fmt.Printf("[func Play] Playing video with ID: %s\n", videoID)

	videoFilePath := fmt.Sprintf("pkg/storage/%s/video.mp4", videoID)
	//	serve the video file
	http.ServeFile(w, r, videoFilePath)
}

//func ServePlaylist(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("[func ServePlaylist] url --->", r.URL, "  method --->", r.Method)
//
//	path := r.URL.Path
//	// Remove the "/play/" prefix from the path
//	videoID := strings.TrimPrefix(path, "/hls/")
//
//	// Use the videoID as needed
//	fmt.Printf("[func ServePlaylist] Playing video with ID: %s\n", videoID)
//
//	playlistPath := fmt.Sprintf("pkg/storage/%s/playlist.m3u8", videoID)
//	fmt.Println(playlistPath)
//	playlistData, err := ioutil.ReadFile(playlistPath)
//
//	if err != nil {
//		http.Error(w, "failed to read playlist file", http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
//	w.Header().Set("Content-Disposition", "inline")
//
//	w.Write(playlistData)
//}

func ServePlaylistCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[func ServePlaylistCheck] url --->", r.URL, "  method --->", r.Method)

	path := r.URL.Path
	//	// Remove the "/play/" prefix from the path
	playlist := strings.TrimPrefix(path, "/hls/")
	playlistPath := fmt.Sprintf("pkg/storage/d804d1e6-32e1-44bf-a7fe-48d35a84df36/%s", playlist)
	playlistData, err := ioutil.ReadFile(playlistPath)

	if err != nil {
		http.Error(w, "failed to read playlist file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	w.Header().Set("Content-Disposition", "inline")

	w.Write(playlistData)

}
