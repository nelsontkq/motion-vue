#!/usr/bin/env bash

declare listInstalledPackages;
declare installPackage;

################################# Get distro ##################################
. /etc/os-release
while :
do
    case $NAME in
        Solus)
            listInstalledPackages="eopkg list-installed"
            installPackage="eopkg install"
            ;;
        Debian)
            listInstalledPackages="apt-get --installed"
            installPackage="apt-get install"
            ;;
        Ubuntu)
            listInstalledPackages="apt-get --installed"
            installPackage="apt-get install"
            ;;
        *)
            echo "$NAME is an unrecognised/unsupported distro."
    esac
done
# Possibly unnecessary future-proofing.
unset NAME
unset VERSION
unset ID
unset VERSION_ID
unset PRETTY_NAME
unset ANSI_COLOR
unset HOME_URL
unset SUPPORT_URL
unset BUG_REPORT_URL
###############################################################################

################################ Check for root ###############################
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi
###############################################################################

################## Check required packages are installed. #####################
declare -a requiredPackages=("golang" "nodejs")
declare notInstalled=""
for i in "${requiredPackages[@]}"; do
  $listInstalledPackages | grep -lq "^$i"
  if [ $? -ge 1 ]; then
    notInstalled+=" $i"
  fi
done
if [ ${#notInstalled} -ge 1 ]; then
  echo Following packages not installed: $notInstalled
  read -r -p "Would you like to install the packages? [y/N] " response
  if [[ "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]
  then
    $installPackage $notInstalled
  else
    echo Required packages not installed. Aborting operation...
    exit 1
  fi
fi
###############################################################################

######################### Install and config backend ##########################
>>>>>>> 69105b518799f5c341fd0ee0c61c5ac9fde99fd1
go get "github.com/gorilla/mux"
go get "github.com/rs/cors"
cd src/video-host
go build -o video-server