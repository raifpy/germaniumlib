package germaniumlib

import (
	"fmt"
	"os"
	"path/filepath"

	findfont "github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func LoadFont(opts Options) (font.Face, error) {
	fontData := ttf_raw
	if opts.Font != "Hack-Regular" {
		fontPath, err := findfont.Find(opts.Font)
		if err != nil {
			return nil, err
		}

		fontData, err = os.ReadFile(fontPath)
		if err != nil {
			return nil, err
		}
	}

	ft, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ft, &truetype.Options{Size: float64(opts.FontSize)}), nil
}

func ListFonts() {
	for _, path := range findfont.List() {
		base := filepath.Base(path)
		ext := filepath.Ext(path)
		if ext == ".ttf" {
			fmt.Println(base[0 : len(base)-len(ext)])
		}
	}
}
