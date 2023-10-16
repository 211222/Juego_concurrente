package models

import (
	
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Dog struct {
    posX, posY float32
    status     bool
    dog        *canvas.Image
    startCh    chan struct{} // Canal para iniciar la carrera del perro
}

func NewDog(posx float32, posy float32, img *canvas.Image, startCh chan struct{}) *Dog {
    return &Dog{
        posX:    posx,
        posY:    posy,
        status:  true,
        startCh: startCh, // canal compartido
		dog:     img,
    }
}



func (d *Dog) Run() {
    var incX float32 = 5 // Incremento en la dirección X 
    d.status = true

    

    for d.status {
        if d.posX > 950 { // Si llega al borde derecho de la ventana
            d.status = false // Detiene la carrera en lugar de reiniciar
        }
        if !d.status {
            // Si la carrera se ha detenido, restablece la posición del perro
            d.posX = 185 // Ajusta la posición inicial 
            d.dog.Move(fyne.NewPos(d.posX, d.posY))
        }

        d.posX += incX

        d.dog.Move(fyne.NewPos(d.posX, d.posY))
        time.Sleep(50 * time.Millisecond) // Controla la velocidad 
    }
}


func (d *Dog) SetStatus(status bool) {
    d.status = status
}