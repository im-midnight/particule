package couleur

import (
	"project-particles/particle"
	"testing"
)

func TestCouleur(t *testing.T) {
	p := particle.NewParticle() // creation d'une particule
	Color(&p)                   // application de la fonction couleur

	if p.ColorRed < 0 || p.ColorRed > 1 {
		t.Errorf("ColorRed n'est pas entre 0 et 1: obtenu %f", p.ColorRed)
	}
	if p.ColorGreen < 0 || p.ColorGreen > 1 {
		t.Errorf("ColorGreen n'est pas entre 0 et 1: obtenu %f", p.ColorGreen)
	}
	if p.ColorBlue < 0 || p.ColorBlue > 1 {
		t.Errorf("ColorBlue n'est pas entre 0 et 1: obtenu %f", p.ColorBlue)
	}
}

func TestCouleur10000s(t *testing.T) {
	p := particle.NewParticle()

	// Appliquer la couleur plusieurs fois
	for i := 0; i < 10000; i++ {
		Color(&p)

		if p.ColorRed < 0 || p.ColorRed > 1 {
			t.Errorf("Appel %d: ColorRed hors limites: %f", i, p.ColorRed)
		}
		if p.ColorGreen < 0 || p.ColorGreen > 1 {
			t.Errorf("Appel %d: ColorGreen hors limites: %f", i, p.ColorGreen)
		}
		if p.ColorBlue < 0 || p.ColorBlue > 1 {
			t.Errorf("Appel %d: ColorBlue hors limites: %f", i, p.ColorBlue)
		}
	}
}
