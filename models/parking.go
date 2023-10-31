package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Parking struct {
	Slots chan bool
	CarsPosition chan *canvas.Image
	M sync.Mutex
	InSlot []Slot
}

type Slot struct {
	X float32
	Y float32
	Availability bool
}

func CreateParking(S int) *Parking {
	return &Parking{
		Slots: make(chan bool, S+1),
		CarsPosition: make(chan *canvas.Image, 100),

		InSlot: []Slot{
			{	X: 0, 
				Y: 69, 
				Availability: true},

			{	X: 135, 
				Y: 69, 
				Availability: true},
			
			{	X: 260, 
				Y: 69, 
				Availability: true},

			{	X: 385, 
				Y: 69, 
				Availability: true},

			{	X: 510, 
				Y: 69, 
				Availability: true},

			{	X: 635, 
				Y: 69, 
				Availability: true},
			
			{	X: 760, 
				Y: 69, 
				Availability: true},

			{	X: 885, 
				Y: 69, 
				Availability: true},
			
			{	X: 260, 
				Y: 390, 
				Availability: true},

			{	X: 385, 
				Y: 390, 
				Availability: true},

			{	X: 510, 
				Y: 390, 
				Availability: true},

			{	X: 635, 
				Y: 390, 
				Availability: true},
			
			{	X: 760, 
				Y: 390, 
				Availability: true},

			{	X: 885, 
				Y: 390, 
				Availability: true},

			{	X: 260, 
				Y: 560, 
				Availability: true},

			{	X: 385, 
				Y: 560, 
				Availability: true},
			
			{	X: 510, 
				Y: 560, 
				Availability: true},

			{	X: 635, 
				Y: 560, 
				Availability: true},
			
			{	X: 760, 
				Y: 560, 
				Availability: true},

			{	X: 885, 
				Y: 560, 
				Availability: true},
			


		},
	}
}