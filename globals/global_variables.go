package globals

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2/app"
)

var (
	Original_images []string
	Original_paths  []string
	Resized_images  []string
)

const SYS_OS = runtime.GOOS

var (
	MyApp  = app.NewWithID("walldo")
	Window = MyApp.NewWindow("Walldo in go")
)

// configuracion
var GridSize = MyApp.Preferences().StringWithFallback("GridSize", "default")
var GridTitles = MyApp.Preferences().StringWithFallback("GridTitles", "Borderless")

var (
	LayoutStyle  = MyApp.Preferences().StringWithFallback("Layout", "Grid")
	FillStrategy = MyApp.Preferences().StringWithFallback("FillStrategy", "Zoom Fill")
)

// archivos de config
var ConfigDir, _ = os.UserHomeDir() // home del usuario
var ConfigPath = ConfigDir          // path de las configuraciones

// Determinar la direccion del archivo de config
// ~/.config/walldo/config.json (unix)
// ~/AppData/Local/walldo/config.json (windows)
func SetGlobalValues() {
	os.Setenv("FYNE_THEME", "dark")

	switch SYS_OS {
	case "windows":
		ConfigDir += "/AppData/Local/walldo/config.json"
		ConfigPath += "/AppData/Local/walldo/"

	default:
		// sistemas Unix (Mac y Linux)
		ConfigDir += "/.config/walldo/config.json"
		ConfigPath += "/.config/walldo/"
	}
}
