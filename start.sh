#!/usr/bin/env bash
cd src/Server
dotnet build
if [ -d $PWD/motion ]; then 
  echo 'motion symlink found.';
else 
  ln -s $HOME/.motion/ $PWD/motion
fi
dotnet run &
pid=$!
cd ../../

npm run serve
kill ${pid}
trap "echo kill ${pid}; exit 1;" INT