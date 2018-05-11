#!/usr/bin/env bash
cd video-host
go build -o video-server
./video-server &
pid=$!
cd ..

npm run serve
kill ${pid}
trap "echo kill ${pid}; exit 1;" INT
