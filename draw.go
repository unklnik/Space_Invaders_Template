package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	inREC z.Rectangle

	START        bool
	invW, invSPC float32

	invNUMW, invNUMH = 11, 5
)

func DRAWCAM() {

	if !START { // GAME
		_, inREC = dRECHollowCenterCNT(CNT, float32(SCRH)-UNIT, float32(SCRH)-UNIT, UNIT/2, z.Maroon) //BORDER
		invW = inREC.Width / float32(invNUMW+2)
		invW -= invSPC
		mENM()
		pl.r.Y = inREC.Y + inREC.Height - pl.siz
		START = true
	}
	dRECHollowCenterCNT(CNT, float32(SCRH)-UNIT, float32(SCRH)-UNIT, UNIT/2, z.Maroon) //BORDER

	//PROJ
	if len(pPRJ) > 0 {
		for i := range pPRJ {
			dREC(pPRJ[i].r, z.Magenta)
		}
	}

	//ENM
	for i := range en {
		if !en[i].off {
			if eLR {
				dIMflip(en[i].im, en[i].r)
			} else {
				dIM(en[i].im, en[i].r)
			}

			if debug {
				dRECLINE(en[i].r, 2, z.Magenta)
			}
		}
	}

	//PL
	dIM(ETC[0], pl.r)

	//DEBUG
	if debug {
		dRECLINE(inREC, 2, z.Magenta) //INREC

	}

}
func DRAWNOCAMSHADER() { //MARK: NOCAM SHADER
	//MOUSE CURSOR
	/*
		siz := UNIT + UNIT/2
		cursorREC = z.NewRectangle(MS.X, MS.Y, siz, siz)
		if clickT > 0 {
			dIMcolorSHADOW(ETC[5], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
		} else {
			dIMcolorSHADOW(ETC[4], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
		}
	*/

}
func DRAWNOCAMNOSHADER() { //MARK: NOCAM NOSHADER
	SCAN(1, 2, z.Fade(z.Black, 0.3))

	if debug {
		DEBUG()

	}
	if scrollONOFF {
		SCROLL()
	}
}
