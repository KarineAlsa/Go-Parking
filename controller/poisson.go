package poisson

import (
	"parking/models"
	"math/rand"
	"time"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func CarsGeneration(cant int, parking *models.Parking) {
	parking.Slots <- true
	
	for i := 1; i < cant; i++ {
		nCar:= rand.Intn(8) + 1
		nameCar:= fmt.Sprintf("./assets/car%d.png", nCar)
		carImage := canvas.NewImageFromURI( storage.NewFileURI(nameCar) )
		carImage.Resize(fyne.NewSize(70,130))
		carImage.Move( fyne.NewPos(50,700 ) )

		newCar := models.CreateCar(parking, carImage)
		newCar.Id = i 

		parking.CarsPosition <- carImage
		time.Sleep(time.Millisecond*200)
		go newCar.Run()

		
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
	}
}
