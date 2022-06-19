package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var assets map[string]*img = make(map[string]*img)
var tmpImg *ebiten.Image = ebiten.NewImage(64, 64)

type img struct {
	downloaded bool
	data       *ebiten.Image
}

func LoadImage(url string) *ebiten.Image {
	if assets[url] != nil {
		if assets[url].downloaded {
			// Return cached image data
			return assets[url].data
		}
		// Return temporary image
		return tmpImg
	}

	// New request, create container and go get it
	fmt.Println("LoadImage")
	assets[url] = &img{}
	go GetImage(url)
	return tmpImg
}

func GetImage(url string) {
	// Downloads image, saves it in the container, marks container as downloaded=true
	fmt.Println("GetImage")
	var err error
	var i *ebiten.Image

	i, err = ebitenutil.NewImageFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	assets[url].data = i
	assets[url].downloaded = true
}
