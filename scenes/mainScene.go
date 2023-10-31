package scenes

import (
	"parking/models"
	Poisson "parking/controller"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type ParkingScene struct {
	window fyne.Window
	content *fyne.Container
}

func NewParkingScene(window fyne.Window) *ParkingScene {
	scene := &ParkingScene{window: window}
    scene.Generate()
    return scene
}

func (s *ParkingScene) Generate() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/bg.png") )
    backgroundImage.Resize(fyne.NewSize(1000,800))
	backgroundImage.Move( fyne.NewPos(0,0) )

	s.content = container.NewWithoutLayout(
        backgroundImage,
    )
    s.window.SetContent(s.content) 
    s.Starting()
}

func (s *ParkingScene) Starting() {
	e := models.CreateParking(20)
	go Poisson.CarsGeneration(100, e)
	go s.CarsPosition(e)
}

func (s *ParkingScene) CarsPosition(e *models.Parking) {
	for {
		imagen := <- e.CarsPosition
		s.content.Add(imagen)
        s.window.Canvas().Refresh(s.content)
	}
}

func Skin( fileUri string, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(50,50))
	image.Move(fyne.NewPos(posX,posY))
	return image
}

