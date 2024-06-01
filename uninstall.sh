#!/bin/bash

APP_NAME="twingate-desktop"

# Remove application and icon
sudo rm -f /opt/${APP_NAME}/${APP_NAME}
sudo rm -f /opt/${APP_NAME}/${APP_NAME}.png
sudo rm -f /opt/${APP_NAME}/twingate.png

# Remove application directory if it is empty
if [ -z "$(ls -A /opt/${APP_NAME})" ]; then
   sudo rmdir /opt/${APP_NAME}
fi

# Remove .desktop file from applications directory
sudo rm -f /usr/share/applications/${APP_NAME}.desktop

# Remove .desktop file from autostart directory
rm -f ~/.config/autostart/${APP_NAME}.desktop
