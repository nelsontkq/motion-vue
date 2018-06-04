#!/usr/bin/env bash
cd src/video-host
go build -o video-server
if [ -d $PWD/motion ]; then 
  echo 'motion symlink found.';
else 
  ln -s $HOME/.motion/ $PWD/motion
fi
./video-server &
pid=$!
cd ../../

npm run serve
kill ${pid}
trap "echo kill ${pid}; exit 1;" INT
