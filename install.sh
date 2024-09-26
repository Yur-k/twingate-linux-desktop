#!/bin/bash

# Install dependencies
sudo apt-get install -y gir1.2-ayatanaappindicator3-0.1 xclip libdbusmenu-glib-dev

APP_NAME="twingate-desktop"
AUTOAPPROVE=false

# Flags were provided, process them
while getopts ":y" opt; do
    case ${opt} in
        y )
            AUTOAPPROVE=true
            ;;
        \? ) echo "Usage: cmd [-y]"
            ;;
    esac
done

sudo mkdir -p /opt/${APP_NAME}
sudo cp ${APP_NAME} /opt/${APP_NAME}/${APP_NAME}

# If AUTOAPPROVE flag is not true, make the script interactive
if [ "$AUTOAPPROVE" = false ] ; then
    read -p "Do you want to add application icon to main menu? (y/n) " response
    if [[ $response =~ ^[Yy]$ ]]
    then
        sudo cp ${APP_NAME}.desktop /usr/share/applications
        sudo cp icons/twingate.png /opt/${APP_NAME}/twingate.png
    fi

    read -p "Do you want start application on startup? (y/n) " response
    if [[ $response =~ ^[Yy]$ ]]
    then
        cp ${APP_NAME}.desktop ~/.config/autostart/${APP_NAME}.desktop
    fi
else
    sudo cp ${APP_NAME}.desktop /usr/share/applications
    sudo cp icons/twingate.png /opt/${APP_NAME}/twingate.png
    cp ${APP_NAME}.desktop ~/.config/autostart/${APP_NAME}.desktop
fi