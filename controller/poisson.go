package poisson

import (
	"parking/models"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func CarsGeneration(cant int, parking *models.Parking) {
	parking.Slots <- true
	for i := 0; i < cant; i++ {
		carImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/car.png") )
		carImage.Resize(fyne.NewSize(70,130))
		x := rand.Intn(700-100+1) + 1
		carImage.Move( fyne.NewPos(float32(x), 500) )

		newCar := models.CreateCar(parking, carImage)
		newCar.Id = i + 1

		parking.CarsPosition <- carImage
		go newCar.Run()
		
		waiting := rand.Intn(700-100+1) + 1
		time.Sleep(time.Duration(waiting) * time.Millisecond)
	}
}
