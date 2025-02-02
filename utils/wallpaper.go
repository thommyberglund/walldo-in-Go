package utils

import (
	"github.com/elias-gill/walldo-in-go/globals"
	"github.com/elias-gill/wallpaper"
)

func WallpaperFitMode() wallpaper.Mode {
	// TODO  poner este switch con un map (diccionarios)
	switch globals.FillStrategy {
	case "Zoom Fill":
		return wallpaper.Fit
	case "Scale":
		return wallpaper.Crop
	case "Center":
		return wallpaper.Center
	case "Original":
		return wallpaper.Span
	case "Tile":
		return wallpaper.Tile
	}
	return wallpaper.Fit
}

// retorna el wallpaper actual
func GetCurrentWallpaper() (string, error) {
	return wallpaper.Get()
}
