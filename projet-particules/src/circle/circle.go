package circle

import (
	"math"
	"math/rand/v2"
	"project-particles/config"
	"project-particles/particle"
	"github.com/hajimehoshi/ebiten/v2"
)

// Position_cercle attire les particules vers une formation circulaire autour du centre de l'écran
func Position_cercle(p *particle.Particle, position int, total int) {
	var centerX = float64(config.General.WindowSizeX) / 2
	var centerY = float64(config.General.WindowSizeY) / 2
	if config.General.Spawn_mouse {
		posXs, posYs := ebiten.CursorPosition()
		// Centre de l'écran
		centerX = float64(posXs)
		centerY = float64(posYs)
	}

	radius := config.General.Radius_circle
	// Angle pour cette particule (basé sur son position)
	angle := 2.0 * math.Pi * float64(position) / float64(total)
	//La formule générale pour un point sur un cercle de centre (cx, cy) et rayon r est :

	//x=cx+r⋅cos⁡(angle)
	//y=cy+r⋅sin(angle)

	// Position cible sur le cercle
	targetX := centerX + radius*math.Cos(angle)
	targetY := centerY + radius*math.Sin(angle)
	// Vecteur de direction vers la cible
	dx := targetX - p.PositionX
	dy := targetY - p.PositionY
	// Distance à la cible
	distance := math.Sqrt(dx*dx + dy*dy)

	if distance > 1 {
		// Normaliser le vecteur pour avoir le vecteur unitaire et appliquer une force de 5
		velocityX := (dx / distance) * 5
		velocityY := (dy / distance) * 5

		p.VitesseX = velocityX
		p.VitesseY = velocityY
	} else {
		p.VitesseX = rand.Float64()*2 - 1
		p.VitesseY = rand.Float64()*2 - 1
	}
}
