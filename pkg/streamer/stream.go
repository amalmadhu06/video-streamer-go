package streamer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Stream(c *gin.Context) {

	//fetch video id from path parameter
	videoID := c.Param("video_id")

	//return error if video id is not there
	if videoID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read video id from request",
			"error":   "invalid request",
		})
		return
	}

	// fetch playlist name from path parameter
	playlist := c.Param("playlist")

	// file path for locating playlists
	playlistPath := fmt.Sprintf("pkg/storage/%s/%s", videoID, playlist)
	playlistData, err := ioutil.ReadFile(playlistPath)

	// handle error that may error while reading file
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read file from server",
			"error":   err.Error(),
		})
		return
	}

	// Set the response headers
	c.Header("Content-Type", "application/vnd.apple.mpegurl")
	c.Header("Content-Disposition", "inline")

	// Write the playlist data to the response body
	c.Writer.Write(playlistData)

}
