package models

import (
	
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	
)

type Rabbit struct {
    posX, posY float32
    status     bool
    rab        *canvas.Image
  
}

func NewRabbit(posx float32, posy float32, img *canvas.Image, startCh chan struct{}) *Rabbit {
    return &Rabbit{
        posX:    posx,
        posY:    posy,
        status:  true,
        rab:     img,
       
    }
}

func (r *Rabbit) Run() {
    var incX float32 = 5 // Incremento en la dirección X 
    r.status = true
	r.posX = 230



    for r.status {
        if r.posX > 1005 { // Si llega al borde derecho de la ventana
            r.status = false // Detiene la carrera en lugar de reiniciar
        }
		if !r.status {
            // Si la carrera se ha detenido, restablece la posición del perro
            r.posX = 235 // Ajusta la posición inicial 
            r.rab.Move(fyne.NewPos(r.posX, r.posY))
        }

        r.posX += incX

        r.rab.Move(fyne.NewPos(r.posX, r.posY))
        time.Sleep(50 * time.Millisecond) // Controla la velocidad 
    }
}






func (r *Rabbit) SetStatus(status bool) {
	r.status = status
}