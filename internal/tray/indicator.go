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
}

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

func NewSystemTrayIndicator(iconGallery IconGalleryI, app ThirdPartyAppI, active string, inactive string) *SystemTrayIndicator {
	return &SystemTrayIndicator{
		iconGallery:      iconGallery,
		app:              app,
		activeIconName:   active,
		inactiveIconName: inactive,
	}
}

func (s SystemTrayIndicator) Run() {
	systray.Run(s.onReady, s.onExit)
}

func (s SystemTrayIndicator) onReady() {

	mStart := systray.AddMenuItem("Start Twingate", "Start")
	mStopNotifications := systray.AddMenuItem("Stop Notifications", "Stop")
	mResumeNotifications := systray.AddMenuItem("Resume Notifications", "Resume")
	mStop := systray.AddMenuItem("Stop Twingate", "Stop")
	mQuit := systray.AddMenuItem("Quit", "Quit")

	s.updateIconByStatus()

	go func() {
		for {
			select {
			case <-mStart.ClickedCh:
				s.app.Start()
			case <-mStop.ClickedCh:
				s.app.Stop()
				s.setInactiveIcon()
			case <-mStopNotifications.ClickedCh:
				s.app.StopNotifications()
			case <-mResumeNotifications.ClickedCh:
				s.app.ResumeNotifications()
			case <-mQuit.ClickedCh:
				s.app.Stop()
				systray.Quit()
			}
		}
	}()
	// Periodically check the status
	go func() {
		for {
			s.updateIconByStatus()
			time.Sleep(1 * time.Second)
		}
	}()
}

func (s SystemTrayIndicator) onExit() {
	log.Println("Exiting wrapper application")
}

func (s SystemTrayIndicator) updateIconByStatus() {
	status, err := s.app.GetStatus()
	if err != nil {
		log.Println("Failed to get status")
		return
	}
	if status == "online" {
		s.setActiveIcon()
	} else {
		s.setInactiveIcon()
	}
}

func (s SystemTrayIndicator) setActiveIcon() {
	systray.SetTemplateIcon(s.iconGallery.GetIcon(s.activeIconName), s.iconGallery.GetIcon(s.activeIconName))
}

func (s SystemTrayIndicator) setInactiveIcon() {
	systray.SetTemplateIcon(s.iconGallery.GetIcon(s.inactiveIconName), s.iconGallery.GetIcon(s.inactiveIconName))
}
