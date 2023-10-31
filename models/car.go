package models

import (
	"math/rand"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Car struct {
	Parking *Parking
	Id int
	position int
	image *canvas.Image
}

func CreateCar(p *Parking, c *canvas.Image) *Car {
	return &Car{
		Parking: p,
		image: c,
	}
}

func (s *Car) Run() {
	s.Parking.Slots <- true
	s.Parking.M.Lock()
	for i := 0; i < len(s.Parking.InSlot); i++ {
		if s.Parking.InSlot[i].Availability {
			s.image.Move(fyne.NewPos( s.Parking.InSlot[i].X , s.Parking.InSlot[i].Y ))
			s.image.Refresh()
			s.position = i
			s.Parking.InSlot[i].Availability = false
			break
		}
	}
	
	time.Sleep(300 *time.Millisecond)
	s.Parking.M.Unlock()

	waiting := rand.Intn(5) 
	time.Sleep(time.Duration(waiting) * time.Second)

	s.Parking.M.Lock()
	<- s.Parking.Slots
	s.Parking.InSlot[s.position].Availability = true
	s.image.Move(fyne.NewPos( 920, 237))
	s.image.Refresh()
	time.Sleep(300 * time.Millisecond)
	s.Parking.M.Unlock()
}
