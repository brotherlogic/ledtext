package lib

import (
	"flag"
	"time"
	"image"

"golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
        "golang.org/x/image/math/fixed"
	        "image/color"

	"github.com/mcuadros/go-rpi-rgb-led-matrix"
)

var (
	rows                     = flag.Int("led-rows", 32, "number of rows supported")
	cols                     = flag.Int("led-cols", 64, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 1, "number of daisy-chained panels")
	chain                    = flag.Int("led-chain", 2, "number of displays daisy-chained")
	brightness               = flag.Int("brightness", 100, "brightness (0-100)")
	hardware_mapping         = flag.String("led-gpio-mapping", "adafruit-hat", "Name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "Don't use hardware pin-pulse generation.")
)

func topline(str string, d time.Duration) {
	config := &rgbmatrix.DefaultConfig
	config.Rows = *rows
	config.Cols = *cols
	config.Parallel = *parallel
	config.ChainLength = *chain
	config.Brightness = *brightness
	config.HardwareMapping = *hardware_mapping
	config.ShowRefreshRate = *show_refresh
	config.InverseColors = *inverse_colors
	config.DisableHardwarePulsing = *disable_hardware_pulsing
	config.PWMLSBNanoseconds = 200

	m, err := rgbmatrix.NewRGBLedMatrix(config)
	fatal(err)

	c := rgbmatrix.NewCanvas(m)
	defer c.Close()
	addLabel(c, 0,13, str)
	c.Render()
	time.Sleep(d)
}

func addLabel(img *rgbmatrix.Canvas, x, y int, label string) {
    col := color.RGBA{255, 0, 0, 255}
        point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

    d := &font.Drawer{
            Dst:  img,
	            Src:  image.NewUniform(col),
		            Face: basicfont.Face7x13,
			            Dot:  point,
				        }
					    d.DrawString(label)
					    }

func init() {
	flag.Parse()
}


func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
