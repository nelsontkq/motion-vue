package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

type Video struct {
	ID   int    `json:"id,omitempty"`
	Path string `json:"path,omitempty`
}

var videos []Video

// ListVideos returns the video array
func ListVideos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(videos)
}

// GetVideo gets the selected video
func GetVideo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range videos {
		var i, _ = strconv.Atoi(params["id"])
		if item.ID == i {
			http.ServeFile(w, r, item.Path)
			return
		}
	}
	json.NewEncoder(w).Encode(&Video{})
}

// DeleteVideo deletes a video
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range videos {
		var i, _ = strconv.Atoi(params["id"])
		if item.ID == i {
			videos = append(videos[:index], videos[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(videos)
	}
}

func FilePathWalkDir(root string) ([]Video, error) {
	var videos []Video
	var i int
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		{
			if info.IsDir() {
				return nil
			}
			videos = append(videos, Video{ID: i, Path: path})
			i++
		}
		return nil
	})
	return videos, err
}

// our main function
func main() {
	var (
		root string
		err  error
	)
	root = "/home/nelsontk/Documents/dev/motion-vue/motion/"
	videos, err = FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/video", ListVideos).Methods("GET")
	router.HandleFunc("/video/{id}", GetVideo).Methods("GET")
	router.HandleFunc("/video/{id}", DeleteVideo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
