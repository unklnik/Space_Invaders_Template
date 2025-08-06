package main

import (
	"time"
)

func UP() {
	INP()
	UPGAME()
	TIMERS()
}

var (
	FRAMES      int32
	FRAMESTIMER time.Time
	SECONDS     int

	TOGfpsHALF, TOGfpsQUART bool
)

func TIMERS() {

	//TOGGLES
	if FRAMES == FPS/2 {
		TOGfpsHALF = !TOGfpsHALF
	}
	if FRAMES != 0 {
		if FRAMES%FPS/4 == 0 {
			TOGfpsQUART = !TOGfpsQUART
		}
	}

	//FRAMES
	if FRAMESTIMER.IsZero() {
		FRAMESTIMER = time.Now()
	}
	if time.Since(FRAMESTIMER) >= time.Second/time.Duration(FPS) {
		FRAMES++
		if FRAMES == FPS {
			SECONDS++
			FRAMES = 0
			FRAMESTIMER = time.Time{}
		}
	}
}

func UPGAME() {

	if START {
		uPL()
		uENM()
		if len(pPRJ) > 0 {
			uPPRJ()
			pPRJ = REMPROJ(pPRJ)
		}

	}
	//MOUSE
	if MSL && clickT == 0 {
		clickT = clickP
	}

}
func uPPRJ() {
	for i := range pPRJ {
		if !pPRJ[i].off {
			collis, num := cPPRJ(pPRJ[i])
			if collis {
				en[num].hp -= pPRJ[i].dm
				if en[num].hp <= 0 {
					en[num].off = true
				}
				pPRJ[i].off = true
			} else {
				if pPRJ[i].r.Y-pPRJ[i].y <= inREC.Y {
					pPRJ[i].off = true
				} else {
					pPRJ[i].r.Y -= pPRJ[i].y
				}
			}
		}
	}
}
func cPPRJ(p PRJ) (bool, int) {
	collides := false
	num := 0
	for i := range en {
		if !en[i].off {
			if cRR(p.r, en[i].r) {
				collides = true
				num = i
				break
			}
		}
	}
	return collides, num
}
func uENM() {

	//MOVE
	if eLR {
		if eXL-eSPD <= inREC.X {
			for i := range en {
				en[i].r.Y += eSPD * 2
			}
			eLR = false
		} else {
			eXL -= eSPD
			eXR -= eSPD
			for i := range en {
				en[i].r.X -= eSPD
			}
		}
	} else {
		if eXR+eSPD >= inREC.X+inREC.Width {
			for i := range en {
				en[i].r.Y += eSPD * 2
			}
			eLR = true
		} else {
			eXR += eSPD
			eXL += eSPD
			for i := range en {
				en[i].r.X += eSPD
			}
		}
	}

}
func uPL() {

	//FIRE
	if kSPC && pl.projT == 0 {
		mPLPROJ()
	}

	//MOVE
	if kL {
		movePL(false)
	}
	if kR {
		movePL(true)
	}
}
