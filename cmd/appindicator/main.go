package main

import (
	"twingate-linux-desktop/internal/application/appindicator"
)

func main() {

	indicator := appindicator.NewApp()
	indicator.Start()
}
