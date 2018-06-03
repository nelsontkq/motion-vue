#!/usr/bin/env bash

declare -a requiredPackages=("golang" "nodejs")
declare -a notInstalled=()
for i in $(requiredPackages[@]); do
    if ! [eopkg list-installed | grep -q "^$i"]; then
        echo $i was not found...
        notInstalled+=($i)
    fi
done
read -r -p "Would you like to install the packages? [y/N] " response
if [[ "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]
then
    installLine=$(printf "%s " "$notInstalled[@]")
    eopkg install $installLine
else
    echo Required packages not installed. Aborting operation...
    exit 1
fi


go get "github.com/gorilla/mux"
go get "github.com/rs/cors"
cd src/video-host
go build -o video-server

if ! [ -d $HOME/.motion/ ]; then 
    mkdir $HOME/.motion/
fi
ln -s $HOME/.motion/ $PWD/motion
