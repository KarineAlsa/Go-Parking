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
	Gate chan bool
	
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
		Gate:        make(chan bool,20),
		InSlot: []Slot{
			{	X: 20, 
				Y: 79, 
				Availability: true},

			{	X: 145, 
				Y: 79, 
				Availability: true},
			
			{	X: 270, 
				Y: 79, 
				Availability: true},

			{	X: 395, 
				Y: 79, 
				Availability: true},

			{	X: 520, 
				Y: 79, 
				Availability: true},

			{	X: 645, 
				Y: 79, 
				Availability: true},
			
			{	X: 770, 
				Y: 79, 
				Availability: true},

			{	X: 895, 
				Y: 79, 
				Availability: true},
			
			{	X: 270, 
				Y: 400, 
				Availability: true},

			{	X: 395, 
				Y: 400, 
				Availability: true},

			{	X: 520, 
				Y: 400, 
				Availability: true},

			{	X: 645, 
				Y: 400, 
				Availability: true},
			
			{	X: 770, 
				Y: 400, 
				Availability: true},

			{	X: 895, 
				Y: 400, 
				Availability: true},

			{	X: 270, 
				Y: 570, 
				Availability: true},

			{	X: 395, 
				Y: 570, 
				Availability: true},
			
			{	X: 520, 
				Y: 570, 
				Availability: true},

			{	X: 645, 
				Y: 570, 
				Availability: true},
			
			{	X: 770, 
				Y: 570, 
				Availability: true},

			{	X: 895, 
				Y: 570, 
				Availability: true},
			


		},
	}
}