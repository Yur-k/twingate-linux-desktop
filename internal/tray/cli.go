package tray

import (
	"flag"
)

// ParseCLIArgs parses the CLI arguments and configures which tray menu items to display
func ParseCLIArgs() *MenuConfig {
	startPtr := flag.Bool("start", true, "Show 'Start' menu item")
	stopPtr := flag.Bool("stop", true, "Show 'Stop' menu item")
	stopNotificationsPtr := flag.Bool("stop-notifications", false, "Show 'Stop Notifications' menu item")
	resumeNotificationsPtr := flag.Bool("resume-notifications", false, "Show 'Resume Notifications' menu item")
	quitPtr := flag.Bool("quit", true, "Show 'Quit' menu item")

	flag.Parse()

	return &MenuConfig{
		Items: map[string]bool{
			"Start Twingate":       *startPtr,
			"Stop Twingate":        *stopPtr,
			"Stop Notifications":   *stopNotificationsPtr,
			"Resume Notifications": *resumeNotificationsPtr,
			"Quit":                 *quitPtr,
		},
	}
}
