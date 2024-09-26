package tray

import (
	"github.com/getlantern/systray"
	"log"
	"time"
)

type SystemTrayIndicator struct {
	iconGallery      IconGalleryI
	app              ThirdPartyAppI
	activeIconName   string
	inactiveIconName string
	menuItems        map[string]*systray.MenuItem
	menuActions      map[string]Action
}

type Action func()

type IconGalleryI interface {
	GetIcon(name string) []byte
}

type ThirdPartyAppI interface {
	Start()
	Stop()
	StopNotifications()
	ResumeNotifications()
	GetStatus() (string, error)
}

type MenuConfig struct {
	Items map[string]bool // Maps menu labels to their visibility (true/false)
}

func NewSystemTrayIndicator(iconGallery IconGalleryI, app ThirdPartyAppI, active string, inactive string) *SystemTrayIndicator {
	return &SystemTrayIndicator{
		iconGallery:      iconGallery,
		app:              app,
		activeIconName:   active,
		inactiveIconName: inactive,
		menuItems:        make(map[string]*systray.MenuItem),
		menuActions:      make(map[string]Action),
	}
}

func (s *SystemTrayIndicator) Run(config *MenuConfig) {
	systray.Run(func() { s.onReady(config) }, s.onExit)
}

func (s *SystemTrayIndicator) onReady(config *MenuConfig) {
	menuOptions := []struct {
		label   string
		tooltip string
		action  Action
	}{
		{"Start Twingate", "Start", s.app.Start},
		{"Stop Twingate", "Stop", func() { s.app.Stop(); s.setInactiveIcon() }},
		{"Stop Notifications", "Stop Notifications", s.app.StopNotifications},
		{"Resume Notifications", "Resume Notifications", s.app.ResumeNotifications},
		{"Quit", "Quit the application", func() { s.app.Stop(); systray.Quit() }},
	}

	for _, option := range menuOptions {
		if config.Items[option.label] {
			s.AddMenuItem(option.label, option.tooltip, option.action)
		}
	}

	s.updateIconByStatus()

	// Periodically check the status
	go func() {
		for {
			s.updateIconByStatus()
			time.Sleep(1 * time.Second)
		}
	}()
}

// AddMenuItem dynamically adds menu items and their associated actions
func (s *SystemTrayIndicator) AddMenuItem(label string, tooltip string, action Action) {
	menuItem := systray.AddMenuItem(label, tooltip)
	s.menuItems[label] = menuItem
	s.menuActions[label] = action
	go func() {
		for {
			select {
			case <-menuItem.ClickedCh:
				action()
			}
		}
	}()
}

func (s *SystemTrayIndicator) onExit() {
	log.Println("Exiting wrapper application")
}

func (s *SystemTrayIndicator) updateIconByStatus() {
	status, err := s.app.GetStatus()
	if err != nil {
		return
	}
	if status == "online" {
		s.setActiveIcon()
	} else {
		s.setInactiveIcon()
	}
}

func (s *SystemTrayIndicator) setActiveIcon() {
	systray.SetTemplateIcon(s.iconGallery.GetIcon(s.activeIconName), s.iconGallery.GetIcon(s.activeIconName))
}

func (s *SystemTrayIndicator) setInactiveIcon() {
	systray.SetTemplateIcon(s.iconGallery.GetIcon(s.inactiveIconName), s.iconGallery.GetIcon(s.inactiveIconName))
}
