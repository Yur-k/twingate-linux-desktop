# Makefile

APP_NAME = twingate-desktop

MAIN_GO = cmd/appindicator/main.go

build:
	CGO_ENABLED=1 go build -o $(APP_NAME) $(MAIN_GO)

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)
	sudo rm /opt/twingate-desktop/$(APP_NAME)
	sudo rm /usr/share/applications/twingate-desktop.desktop
	sudo rm /opt/twingate-desktop/twingate.png

deploy: build
	sudo mkdir -p /opt/twingate-desktop
	sudo cp $(APP_NAME) /opt/twingate-desktop/$(APP_NAME)
	sudo cp icons/twingate.png /opt/twingate-desktop/twingate.png
	sudo cp twingate-desktop.desktop /usr/share/applications
	cp twingate-desktop.desktop ~/.config/autostart/twingate-desktop.desktop