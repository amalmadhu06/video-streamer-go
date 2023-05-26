package uploader

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

const (
	storageLocation = "storage"
)

type Response struct {
	Message string
	Data    interface{}
}

//var wg *sync.WaitGroup

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url : %v --> Method : %v \n", r.URL, r.Method)
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	// parse video file from request
	file, header, err := r.FormFile("video")

	// handle the error that may occur while parsing video content from request
	if err != nil {
		http.Error(w, "failed to retrieve video file from request", http.StatusBadRequest)
		return
	}

	// close the file once operations are over
	defer file.Close()

	//create a new folder for storing the file
	fileUuid := uuid.New()
	fileName := fileUuid.String()
	folderPath := "pkg/" + storageLocation + "/" + fileName
	filePath := "pkg/" + storageLocation + "/" + fileName + "/" + header.Filename
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		http.Error(w, "failed to create new folder", http.StatusInternalServerError)
	}
	// create a new file in the storage/fileName folder
	newFile, err := os.Create(filePath)

	// handle the error that may occur while creating new file
	if err != nil {
		http.Error(w, "failed to create new file", http.StatusInternalServerError)
		return
	}

	// close the newly created file after operation
	defer newFile.Close()

	//	copy uploaded file to new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "failed to copy uploaded file to storage", http.StatusInternalServerError)
		return
	}

	//wg.Add(1)
	//go
	CreatePlaylistAndSegments(filePath, folderPath)
	//wg.Wait()
	response := Response{
		Message: "Success",
		Data:    fileName,
	}

	//convert response struct to json
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.Write(jsonResponse)
}

func CreatePlaylistAndSegments(filePath string, folderPath string) error {
	//defer wg.Done()
	//TODO : calculate segment duration depending on video length
	segmentDuration := 3
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", filePath,
		"-profile:v", "baseline", // baseline profile is compatible with most devices
		"-level", "3.0",
		"-start_number", "0", // start number segments from 0
		"-hls_time", strconv.Itoa(segmentDuration), //duration of each segment in second
		"-hls_list_size", "0", // keep all segments in the playlist
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", folderPath),
	)
	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to created HLS: %v \nOutput: %s ", err, string(output))
	}
	return nil
}
