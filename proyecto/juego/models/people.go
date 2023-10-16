package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type People struct {
	posX, posY float32
	status     bool
	peo        *canvas.Image
}

func NewPeople(posx float32, posy float32, img *canvas.Image) *People {
	return &People{
		posX:   posx,
		posY:   posy,
		status: true,
		peo:    img,
	}
}
func (p *People) Run() {
	randomSpeed := rand.Float32() * 7 // Ajusta el rango de velocidades 

	p.status = true
	p.posX = 980 // Posición inicial en la parte superior

	for p.status {
		if p.posY > 450 { // Si llega al borde inferior de la ventana
			p.status = false // Detiene la carrera en lugar de reiniciar
		}
		if !p.status {
			// Si la carrera se ha detenido, restablece la posición del personaje
			p.posY = 30 // Ajusta la posición
			p.peo.Move(fyne.NewPos(p.posX, p.posY))
		}

		
		p.posY += randomSpeed

		p.peo.Move(fyne.NewPos(p.posX, p.posY))
		time.Sleep(50 * time.Millisecond) // Controla la velocidad 
	}
}
func (p *People) SetStatus(status bool) {
	p.status = status
}
