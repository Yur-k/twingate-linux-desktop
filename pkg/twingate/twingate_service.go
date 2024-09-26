package twingate

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type TwingateService struct{}

func NewTwingateService() *TwingateService {
	return &TwingateService{}
}

func (t TwingateService) Start() {
	startCmd := exec.Command("/usr/bin/twingate", "start")
	startCmd.Stdout = os.Stdout
	startCmd.Stderr = os.Stderr
	err := startCmd.Run()
	if err != nil {
		log.Printf("Failed to start app: %v\n", err)
	} else {
		log.Println("App started")
	}
}

func (t TwingateService) Stop() {
	stopCmd := exec.Command("/usr/bin/twingate", "stop")
	stopCmd.Stdout = os.Stdout
	stopCmd.Stderr = os.Stderr
	err := stopCmd.Run()
	if err != nil {
		log.Printf("Failed to stop app: %v\n", err)
	} else {
		log.Println("App stopped")
	}
}

func (t TwingateService) GetStatus() (string, error) {
	statusCmd := exec.Command("/usr/bin/twingate", "status")
	output, err := statusCmd.Output()
	if err != nil {
		log.Printf("Failed to get status: %v\n", err)
		return "", err
	}

	status := string(output)
	status = strings.TrimSpace(status)
	return status, nil
}

func (t TwingateService) StopNotifications() {
	stopNotificationsCmd := exec.Command("/usr/bin/twingate", "desktop-stop")
	err := stopNotificationsCmd.Run()
	if err != nil {
		log.Printf("Failed to stop notifications: %v\n", err)
	} else {
		log.Println("Notifications stopped")
	}
}

func (t TwingateService) ResumeNotifications() {
	startNotificationsCmd := exec.Command("/usr/bin/twingate", "desktop-start")
	err := startNotificationsCmd.Run()
	if err != nil {
		log.Printf("Failed to start notifications: %v\n", err)
	} else {
		log.Println("Notifications started")
	}
}

func (t TwingateService) IsInstalled() bool {
	_, err := exec.LookPath("/usr/bin/twingate")
	return err == nil
}
