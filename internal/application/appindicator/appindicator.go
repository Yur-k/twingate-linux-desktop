package appindicator

import (
	"log"
	"twingate-linux-desktop/icons"
	"twingate-linux-desktop/internal/tray"
	"twingate-linux-desktop/pkg/icongallery"
	twingateWrapper "twingate-linux-desktop/pkg/twingate"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a App) Start() {
	var (
		activeIconName   = "active"
		inactiveIconName = "inactive"
	)

	twingate := twingateWrapper.NewTwingateService()
	if !twingate.IsInstalled() {
		log.Fatalf("Twingate is not installed")
	}
	iconGallery := icongallery.NewIconGallery()
	iconGallery.AddIconByByte(activeIconName, icons.Active)
	iconGallery.AddIconByByte(inactiveIconName, icons.Inactive)

	menuConfig := tray.ParseCLIArgs()
	indicator := tray.NewSystemTrayIndicator(iconGallery, twingate, activeIconName, inactiveIconName)
	indicator.Run(menuConfig)
}
