package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
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
			// Return cached imag dta
			return assets[url].data
		}
		// Return temorary imag
		return tmpImg
	}

	// New request, creae cntainer and go get it
	fmt.Println("LoaImage")
	assets[url] = &img{}
	go GetImage(url)
	return tmpImg
}

func GetImage(url string) {
	// Downloads mage, savs it in the container, marks container as downloaded=true
	fmt.Println("getImage")

	//i, err = ebitenutil.NewImageFromURL(url)

	//if err != nil {
	//	log.Fatal(err)
	//}

	assets[url].data = load(url)
	assets[url].downloaded = true

}

//go:embed godragon_32.png
var godragon []byte

//go:embed map_2056.png
var gomap []byte

func load(url string) *ebiten.Image {
	if url == "./img/godragon_32.png" {
		d, _, _ := image.Decode(bytes.NewReader(godragon))
		img := ebiten.NewImageFromImage(d)
		return img
	} else {
		dd, _, _ := image.Decode(bytes.NewReader(gomap))
		img := ebiten.NewImageFromImage(dd)
		return img
	}

}
