package particle

import (
	"math/rand"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
Renvoie une particule définie par

	PositionX, PositionY            float64 valeur par défaut WindowSizeX, WindowSizeY
	Rotation                        float64 valeur par défaut 0
	ScaleX, ScaleY                  float64 valeur par défaut 1,1
	ColorRed, ColorGreen, ColorBlue float64 valeur par défaut 1,1,1
	Opacity 						float64 valeur par défaut 1

La particule peut ne pas être à l'écran
*/

func NewParticle() Particle {
	// position de spawn
	var posX, posY float64 = float64(config.General.WindowSizeX) / 2, float64(config.General.WindowSizeY) / 2

	spawn_mouse := config.General.Spawn_mouse // si le spawn suit la souris
	random := config.General.RandomSpawn      // si le spawn est aleatoire

	if spawn_mouse {
		posXs, posYs := ebiten.CursorPosition() // position de la souris
		// verif si la position est hors ecran
		if float64(posXs) > float64(config.General.WindowSizeX)-10 ||
			posXs <= 0 ||
			float64(posYs) > float64(config.General.WindowSizeY) ||
			posYs < 0 {
			if random { // spawn aleatoire si hors ecran et random active
				posX = rand.Float64() * float64(config.General.WindowSizeX)
				posY = rand.Float64() * float64(config.General.WindowSizeY)
			}

		} else { // spawn a la position de la souris
			posX = float64(posXs)
			posY = float64(posYs)
		}

	} else if random { // spawn aleatoire
		posX = rand.Float64() * float64(config.General.WindowSizeX)
		posY = rand.Float64() * float64(config.General.WindowSizeY)
	}

	scale := rand.Float64()*1.2 + 1

	return Particle{ // creation de la particule avec les valeurs
		PositionX:  posX,
		PositionY:  posY,
		Rotation:   0,
		ScaleX:     scale,
		ScaleY:     scale,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		VitesseX:   rand.Float64()*2 - 1,
		VitesseY:   rand.Float64()*2 - 1,
	}
}
