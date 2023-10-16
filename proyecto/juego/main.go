package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"jueguito/scenes"
)

func main(){
	juego := app.New()
	ventana := juego.NewWindow("Carrera de perros")
	ventana.CenterOnScreen()
	ventana.SetFixedSize(true)
	ventana.Resize(fyne.NewSize(1200, 625))

	//Cargar y mostrar la escena principal
	mainMenuScene := scenes.NewMainMenuScene(ventana)
	mainMenuScene.Show()
	ventana.ShowAndRun()
}
