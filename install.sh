#!/bin/bash

# Install dependencies
sudo apt-get install -y gir1.2-ayatanaappindicator3-0.1 xclip libdbusmenu-glib-dev

APP_NAME="twingate-desktop"
AUTOAPPROVE=false
START=true
STOP=true
STOP_NOTIFICATIONS=false
RESUME_NOTIFICATIONS=false

# Flags were provided, process them
while getopts ":y-:" opt; do
    case ${opt} in
        y )
          AUTOAPPROVE=true
          ;;
        - )
            case "${OPTARG}" in
                start=*)
                    START=${OPTARG#*=}
                    ;;
                stop=*)
                    STOP=${OPTARG#*=}
                    ;;
                stop-notifications=*)
                    STOP_NOTIFICATIONS=${OPTARG#*=}
                    ;;
                *)
                    echo "Failed. Usage: cmd [-y] [--start=true|false] [--stop=true|false] [--stop-notifications=true|false] resume-notifications=true|false"
                    exit 1
                    ;;
            esac
            ;;
        \? ) echo "Failed. Usage: cmd [-y] [--start=true|false] [--stop=true|false] [--stop-notifications=true|false] resume-notifications=true|false"
            exit 1
            ;;
    esac
done

sudo mkdir -p /opt/${APP_NAME}
sudo cp ${APP_NAME} /opt/${APP_NAME}/${APP_NAME}

# Update .desktop file with the provided arguments
DESKTOP_FILE_CONTENT="[Desktop Entry]
Version=0.0.1
Name=Twingate Desktop
Exec=\"/opt/twingate-desktop/twingate-desktop\" -start=${START} -stop=${STOP} -stop-notifications=${STOP_NOTIFICATIONS} -resume-notifications=${RESUME_NOTIFICATIONS} > /dev/null 2>&1
Icon=/opt/twingate-desktop/twingate.png
Type=Application
StartupNotify=true
X-GNOME-Autostart-enabled=true
Encoding=UTF-8"

echo "${DESKTOP_FILE_CONTENT}" | sudo tee /usr/share/applications/${APP_NAME}.desktop > /dev/null

# If AUTOAPPROVE flag is not true, make the script interactive
if [ "$AUTOAPPROVE" = false ] ; then
    read -p "Do you want to add application icon to main menu? (y/n) " response
    if [[ $response =~ ^[Yy]$ ]]
    then
        sudo cp icons/twingate.png /opt/${APP_NAME}/twingate.png
    fi

    read -p "Do you want start application on startup? (y/n) " response
    if [[ $response =~ ^[Yy]$ ]]
    then
        cp /usr/share/applications/${APP_NAME}.desktop ~/.config/autostart/${APP_NAME}.desktop
    fi
else
    sudo cp icons/twingate.png /opt/${APP_NAME}/twingate.png
    cp /usr/share/applications/${APP_NAME}.desktop ~/.config/autostart/${APP_NAME}.desktop
fi