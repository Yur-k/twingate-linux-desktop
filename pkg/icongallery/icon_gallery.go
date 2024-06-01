package icongallery

import (
	"log"
	"os"
)

type IconGallery struct {
	Icons map[string][]byte
}

func NewIconGallery() *IconGallery {
	return &IconGallery{
		Icons: make(map[string][]byte),
	}
}

func (i IconGallery) AddIconFromStorage(name string, path string) {
	i.Icons[name] = i.readIcon(path)
}

func (i IconGallery) AddIconByByte(name string, icon []byte) {
	i.Icons[name] = icon
}

func (i IconGallery) GetIcon(name string) []byte {
	return i.Icons[name]
}

func (i IconGallery) readIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		log.Fatalf("Error reading icon: %s", err)
	}
	return b
}
