package scenes

import (
	_ "fmt"
	//"image/color"
	"jueguito/models"
	"time"
	"math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"sync"
)

type MainMenuScene struct {
	window  fyne.Window
	startCh chan struct{} // Canal para iniciar las goroutines de Dog y Rabbit
	wg sync.WaitGroup
}

var dogs []*models.Dog
var r *models.Rabbit
var p *models.People

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window}
}

func (s *MainMenuScene) Show() {
// Crear un canal compartido para iniciar las goroutines de Dog y Rabbit
startCh := make(chan struct{})
	//crear la imagen de background

	

	dog1Image := canvas.NewImageFromURI(storage.NewFileURI("./assets/perro.png"))
    dog1Image.Resize(fyne.NewSize(50, 50))
    dog1Image.Move(fyne.NewPos(100, 100))
    dog1 := models.NewDog(100, 100, dog1Image, startCh)
    dogs = append(dogs, dog1)

    dog2Image := canvas.NewImageFromURI(storage.NewFileURI("./assets/perro2.png"))
    dog2Image.Resize(fyne.NewSize(50, 50))
    dog2Image.Move(fyne.NewPos(100, 200))
    dog2 := models.NewDog(100, 200, dog2Image, startCh)
    dogs = append(dogs, dog2)

    dog3Image := canvas.NewImageFromURI(storage.NewFileURI("./assets/perro3.png"))
    dog3Image.Resize(fyne.NewSize(50, 50))
    dog3Image.Move(fyne.NewPos(100, 300))
    dog3 := models.NewDog(100, 300, dog3Image, startCh)
    dogs = append(dogs, dog3)

	rabbit := canvas.NewImageFromURI(storage.NewFileURI("./assets/conejo1.png"))
	rabbit.Resize(fyne.NewSize(150, 100))
	rabbit.Move(fyne.NewPos(235, 200))	
	r = models.NewRabbit(235, 200, rabbit, startCh)

	people := canvas.NewImageFromURI(storage.NewFileURI("./assets/arbitro.png"))
	people.Resize(fyne.NewSize(150, 100))
	people.Move(fyne.NewPos(980, 30))
	//Creamos el modelo
	p = models.NewPeople(980,95,people)


	//green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	// Botón "Play" con un color de fondo personalizado
	botonIniciar := widget.NewButton("Play", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(150, 30))
	botonIniciar.Move(fyne.NewPos(520, 8))
	botonIniciar.SetIcon(theme.MediaPlayIcon())

	// Botón "Stop" con un color de fondo personalizado
	botonDetener := widget.NewButton("Stop", s.StopGame)
	botonDetener.Resize(fyne.NewSize(150, 30))
	botonDetener.Move(fyne.NewPos(520, 40))
	botonDetener.SetIcon(theme.MediaStopIcon())

	fondo := canvas.NewImageFromURI(storage.NewFileURI("./assets/campo.jpg"))
	fondo.FillMode = canvas.ImageFillOriginal
	fondo.Resize(fyne.NewSize(1200, 600))

	content := fyne.NewContainerWithoutLayout(fondo, dog1Image, dog2Image,dog3Image, rabbit, people, botonDetener, botonIniciar)
	s.window.SetContent(container.NewWithoutLayout(content))
}




func (s *MainMenuScene) StartGame() {
	go p.Run()

 
	go r.Run()

	
    // Crea un canal para señalar a los perros que deben comenzar
    startDogsCh := make(chan struct{})

    // Inicia todas las goroutines de los perros
    for _, dog := range dogs {
        go func(dog *models.Dog) {
            <-startDogsCh // Espera la señal de inicio
				    // Esperar un tiempo aleatorio antes de enviar la señal para iniciar la carrera
		randomDelay := time.Duration(rand.Intn(5000)) * time.Millisecond // Tiempo aleatorio entre 0 y 5 segundos
		time.Sleep(randomDelay)

            dog.Run()
        }(dog)
    }
	    // Esperar un tiempo aleatorio antes de enviar la señal para iniciar la carrera
		

    // Cierra el canal startDogsCh para señalar que todos los perros deben comenzar simultáneamente
    close(startDogsCh)


}

func (s *MainMenuScene) StopGame() {

p.SetStatus(false)

 // Detener todos los perros estableciendo su estado en falso
 for _, dog := range dogs {
	dog.SetStatus(false)
}

// Esperar a que todas las goroutines de los perros terminen
s.wg.Wait()
	r.SetStatus(false)

}
