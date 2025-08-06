package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	debug, debug2, debugBOOL bool
	debugNUM                 int
)

func DEBUG() {

	debugBOOL = FULLSCREEN

	if !debug2 {
		dREC(z.NewRectangle(0, 0, 300, float32(SCRH)), CA(z.Maroon, 150))
	}
	var x, y, sizTXT float32 = 4, 4, 0.5
	h := qTXTheight(FONT1, sizTXT)
	dTXTfont1XY("mouseX "+fmt.Sprintf("%.0f", MS.X)+" mouseY "+fmt.Sprintf("%.0f", MS.Y), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("FRAMES "+fmt.Sprint(FRAMES)+" debugBOOL "+fmt.Sprint(debugBOOL), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("debugNUM "+fmt.Sprint(debugNUM)+" debugBOOL "+fmt.Sprint(debugBOOL), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("inREC.X "+fmt.Sprint(inREC.X)+" en[0].r.X "+fmt.Sprint(en[0].r.X), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("eXL "+fmt.Sprint(eXL)+" eXR "+fmt.Sprint(eXR), x, y, sizTXT, z.White)
	y += h

}
