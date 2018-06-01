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
	Name string `json:"name,omitempty"`
}

var videos []Video
var root string

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
			fmt.Println("serving " + item.Name)
			http.ServeFile(w, r, filepath.Join(root, item.Name))
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
			err := os.Rename(path.Join(root, item.Name), path.Join(root, ".trash", item.Name))
			if err != nil {
				panic(err)
			}
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
			videos = append(videos, Video{ID: i, Name: filepath.Base(path)})
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
func dailyCleanUp(root string) {
	for range time.Tick(time.Hour * 24) {
		removeContents(root)
	}
}

func folderWatch(ms int64) {
	for range time.Tick(time.Duration(ms) * time.Millisecond) {
		fmt.Println("Scanning for new files...")
	}
}

// our main function
func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	root, err = filepath.EvalSymlinks(filepath.Join(filepath.Dir(ex), "motion"))
	if err != nil {
		panic(err)
	}
	trashDir := path.Join(root, ".trash")
	fmt.Println("root " + root)
	videos, err = filePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	fmt.Println("videos found: ", videos)
	trashExists, err := exists(trashDir)
	if err != nil {
		panic(err)
	}
	if !trashExists {
		os.Mkdir(trashDir, os.ModePerm)
	}
	go folderWatch(1000) // ms
	go dailyCleanUp(trashDir)

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
