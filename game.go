package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	pl             PLAYER
	en             []ENM
	eXL, eXR, eSPD float32
	eLR            bool

	pPRJ, ePRJ []PRJ

	enIM []IM
)

type ENM struct {
	r           z.Rectangle
	hp, numType int
	off         bool
	im          IM
	c           z.Color
}
type PRJ struct {
	x, y, siz float32
	r         z.Rectangle
	dm        int
	off       bool
}
type PLAYER struct {
	r        z.Rectangle
	c        z.Color
	spd, siz float32
	projT    int32
}

func INITGAME() { //MARK: INIT GAME

	//z.SetExitKey(z.KeyEnd) //DELETE

	invSPC = UNIT / 5
	eSPD = UNIT / 10

	mPL()
}

// MARK: MOVE
func movePL(dir bool) {

	if dir { //RIGHT
		if pl.r.X < inREC.X+inREC.Width {
			pl.r.X += pl.spd
		}
		if pl.r.X+pl.r.Width > inREC.X+inREC.Width {
			pl.r.X = inREC.X + inREC.Width - +pl.r.Width
		}
	} else { //LEFT
		if pl.r.X > inREC.X {
			pl.r.X -= pl.spd
		}
		if pl.r.X < inREC.X {
			pl.r.X = inREC.X
		}
	}

}

// MARK: MAKE
func mPLPROJ() {

	p := PRJ{}
	p.siz = UNIT / 4
	p.r = RECCNT(qRecCNT(pl.r), p.siz, p.siz)
	p.r.Y = pl.r.Y - p.r.Height
	p.y = UNIT / 5
	p.dm = 1

	pPRJ = append(pPRJ, p)

}
func mENM() {

	//MAKE RECS
	y := inREC.Y + UNIT*2
	w := float32(invNUMW) * (invW + invSPC)
	w -= invSPC
	x := CNT.X - w/2
	eXL = x
	eXR = x + w
	ox := x
	a := invNUMW * invNUMH
	c := 0
	row := 0
	for a > 0 {
		e := ENM{}
		e.numType = row
		e.r = R(x, y, invW, invW)
		en = append(en, e)
		x += invW + invSPC
		c++
		a--
		if c == invNUMW {
			row++
			c = 0
			x = ox
			y += invW + invSPC
		}
	}

	//HP
	for i := range en {
		switch en[i].numType {
		case 0:
			en[i].im = enIM[0]
			en[i].c = z.Maroon
		case 1:
			en[i].im = enIM[1]
			en[i].c = z.Magenta
		case 2:
			en[i].im = enIM[2]
			en[i].c = z.DarkGreen
		case 3:
			en[i].im = enIM[3]
			en[i].c = z.DarkBlue
		case 4:
			en[i].im = enIM[0]
			en[i].c = z.DarkPurple
		}

		en[i].hp = 1

	}

}
func mPL() {
	pl.siz = UNIT
	pl.c = z.Orange
	//pl.r = z.NewRectangle(CNT.X-pl.siz/2, inREC.Y+inREC.Height-pl.siz, pl.siz, pl.siz)
	pl.r = RECCNT(CNT, pl.siz, pl.siz)
	pl.spd = UNIT / 8

}

// UTILS
func REMPROJ(p []PRJ) []PRJ {
	var p2 []PRJ
	for i := range p {
		if !p[i].off {
			p2 = append(p2, p[i])
		}
	}

	return p2
}
