package couleur

import (
	"project-particles/config"
	"project-particles/particle"
)

func Color(p *particle.Particle) {
	//couleur

	// Dégradé vertical basé sur la position Y de la particule
	h := p.PositionY / float64(config.General.WindowSizeY)

	// Dégradé : rouge -> jaune -> blanc
	if h < 0.5 {
		// Rouge -> Jaune
		t := h * 2
		p.ColorRed = 1
		p.ColorGreen = t
		p.ColorBlue = 0
	} else {
		// Jaune -> Blanc
		t := (h - 0.5) * 2
		p.ColorRed = 1
		p.ColorGreen = 1
		p.ColorBlue = t
	}

}
