package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Car struct {
	Estacionamiento *Estacionamiento
	I               int
	skin			*canvas.Image
}

func CreateCarro(e *Estacionamiento, s *canvas.Image) *Car {
	return &Car{
		Estacionamiento: e,
		skin: s,
	}
}

func (c *Car) RunCarro() {

	c.Estacionamiento.SlotsEstacionamiento <- true
	c.Estacionamiento.M.Lock()
		x := float32( rand.Intn(650-150+1) )
		y := float32( rand.Intn(300-50+1) )
		c.skin.Move(fyne.NewPos( x, y ))
		fmt.Println("Carro ", c.I, " Entra")
		time.Sleep(200 *time.Millisecond)
	c.Estacionamiento.M.Unlock()

	TiempoEsperar := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(TiempoEsperar) * time.Second)

	c.Estacionamiento.M.Lock()
		<- c.Estacionamiento.SlotsEstacionamiento
		c.skin.Move(fyne.NewPos( 0,0 ))
		fmt.Println("Carro ", c.I, " Sale")
		time.Sleep(200 *time.Millisecond)
	c.Estacionamiento.M.Unlock()
}
