package texte

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var flecheX = 23
var flecheY = 20
var step = 20
var keyUpPrev bool
var keyDownPrev bool
var keyEnterPrev bool
var keyShiftPrev bool

func Fleche_draw(Screen *ebiten.Image) {
	keyEnter := ebiten.IsKeyPressed(ebiten.KeyEnter)
	keyShift := ebiten.IsKeyPressed(ebiten.KeyShift)
	keyUp := ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)
	keyDown := ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)

	// Déplacer uniquement si la touche est juste pressée
	if keyUp && !keyUpPrev {
		flecheY -= step
	}
	if keyDown && !keyDownPrev {
		flecheY += step
	}
	// Limites
	if flecheY < 20 {
		flecheY = 160
	}
	if flecheY > 160 {
		flecheY = 20
	}

	// Dessiner la flèche
	ebitenutil.DebugPrintAt(Screen, ">>>", flecheX, flecheY)
	ebitenutil.DebugPrintAt(Screen, ">>>", flecheX+1, flecheY)

	// Appuie de Enter, une seule fois par appui
	if keyEnter && !keyEnterPrev {
		switch flecheY {
		case 20:
			config.General.Gravity = !config.General.Gravity
		case 40:
			config.General.Gravity_value += 0.1
		case 60:
			config.General.Collison = !config.General.Collison
		case 80:
			config.General.Spawn_mouse = !config.General.Spawn_mouse
		case 100:
			config.General.SpawnRate += 0.1
		case 120:
			config.General.RandomSpawn = !config.General.RandomSpawn
		case 140:
			config.General.IsCircle = !config.General.IsCircle
		case 160:
			config.General.Radius_circle += 10
		}
	}

	// Shift pour diminuer SpawnRate et Rayon du Cercle
	if keyShift && !keyShiftPrev {
		switch flecheY {
		case 40:
			config.General.Gravity_value -= 0.1
			if config.General.Gravity_value < 0 {
				config.General.Gravity_value = 0
			}

		case 100:
			config.General.SpawnRate -= 0.1
			if config.General.SpawnRate < 0 {
				config.General.SpawnRate = 0
			}
		case 160:
			config.General.Radius_circle -= 10
			if config.General.Radius_circle < 10 {
				config.General.Radius_circle = 10

			}
		}

	}

	// Mettre à jour l'état précédent
	keyUpPrev = keyUp
	keyDownPrev = keyDown
	keyEnterPrev = keyEnter
	keyShiftPrev = keyShift
}
