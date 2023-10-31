package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Parking struct {
	Slots chan bool
	CarsPosition chan *canvas.Image
	M sync.Mutex
}

func CreateParking(nS int) *Parking {
	return &Parking{
		Slots: make(chan bool, nS+1),
		CarsPosition: make(chan *canvas.Image, 1),
	}
}