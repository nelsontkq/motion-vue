#!/usr/bin/env bash
go get "github.com/gorilla/mux"
go get "github.com/rs/cors"
cd src/video-host
go build -o video-server
ln -s ~/Documents/motion/ $PWD/motion