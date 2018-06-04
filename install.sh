#!/usr/bin/env bash
npm install
go get "github.com/gorilla/mux"
go get "github.com/rs/cors"
cd src/video-host
go build -o video-server
# This is sthe standard directory 
ln -s $HOME/.motion/ $PWD/motion