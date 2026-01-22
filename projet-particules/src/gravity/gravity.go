package gravity

import (
	"project-particles/particle"
	"project-particles/config"
)

func Applique_gravite(p *particle.Particle, Dy float64, WindowSizeY float64) {
	// la fonction prend a parametre un particule,la taille en Y et la taille de l'ecran en Y
	// verif si le particule ne bouge plus
	sol := float64(WindowSizeY) - Dy*p.ScaleY

	p.VitesseY += config.General.Gravity_value // gravite

	p.PositionX += p.VitesseX
	p.PositionY += p.VitesseY

	if p.PositionY >= sol { // si il traverse ou touche le sol
		p.PositionY = sol
		//coeff de perte energie
		p.VitesseY = -p.VitesseY * 0.7
		// ralentissement horizontal
		p.VitesseX *= 0.98
	}

}
