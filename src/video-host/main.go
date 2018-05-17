package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Video struct holding relevant videos
type Video struct {
	ID   int    `json:"id,omitempty"`
	Path string `json:"path,omitempty`
}

var videos []Video
var root = "/home/nelsontk/Documents/dev/motion-vue/motion/"

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

// DeleteVideo renames video to .trash and removes it from the array.
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range videos {
		var i, _ = strconv.Atoi(params["id"])
		if item.ID == i {
			videos = append(videos[:index], videos[index+1:]...)
			err := os.Rename(path.Join(root, item.Path), path.Join(root, ".trash", item.Path))
			if err != nil {
				panic(err)
			}
			fmt.Println("moved file to trash.")
			break
		}
		json.NewEncoder(w).Encode(videos)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func filePathWalkDir(root string) ([]Video, error) {
	var videos []Video
	var i int
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		{
			if info.Name() == "motion" {
				return nil
			} else if info.IsDir() {
				return filepath.SkipDir
			}
			i++
			videos = append(videos, Video{ID: i, Path: filepath.Base(path)})
		}
		return nil
	})
	return videos, err
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func heartBeat(root string) {
	for range time.Tick(time.Hour * 24) {
		fmt.Println("Clearing trash...")
		removeContents(root)
	}
}

// our main function
func main() {
	var (
		err         error
		pathErr     error
		trashExists bool
	)
	var trashDir = path.Join(root, ".trash")
	videos, err = filePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	trashExists, pathErr = exists(trashDir)
	if pathErr != nil {
		panic(pathErr)
	}
	if !trashExists {
		os.Mkdir(trashDir, os.ModePerm)
	}
	go heartBeat(trashDir)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "DELETE"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/video", ListVideos).Methods("GET")
	router.HandleFunc("/video/{id}", GetVideo).Methods("GET")
	router.HandleFunc("/video/{id}", DeleteVideo).Methods("DELETE")
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
