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
	for i := 1; i < cant; i++ {
		carImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/car.png") )
		carImage.Resize(fyne.NewSize(70,130))
		carImage.Move( fyne.NewPos(50,700 ) )

		newCar := models.CreateCar(parking, carImage)
		newCar.Id = i 

		parking.CarsPosition <- carImage
		go newCar.Run()

		waiting := rand.Intn(700-100+1) + 1
		time.Sleep(time.Duration(waiting) * time.Millisecond)
	}
}
