package gravity

import (
	"project-particles/particle"
	"testing"
)

func TestApplique_gravite_acceleration(t *testing.T) {
	// Initialise la config pour le test
	p := &particle.Particle{
		PositionX: 100,
		PositionY: 100,
		VitesseX:  0,
		VitesseY:  0,
		ScaleX:    1,
		ScaleY:    1,
	}

	oldV := p.VitesseY
	Applique_gravite(p, 5, 600)

	// Vérifier que la vitesse a augmenté de 0.2 (la gravité)
	if p.VitesseY != oldV+0.2 {
		t.Errorf("VitesseY devrait augmenter de 0.2, ancien=%f, nouveau=%f", oldV, p.VitesseY)
	}
}

func TestApplique_gravite_accumulation(t *testing.T) {
	p := &particle.Particle{
		VitesseY: 0,
		ScaleX:   1,
		ScaleY:   1}

	Applique_gravite(p, 5.0, 600.0)
	Applique_gravite(p, 5.0, 600.0)
	expected := 0.4
	// Vérifier que la vitesse a augmenté de 0.4 (2fois)
	if p.VitesseY != expected {
		t.Errorf("Après 2 frames, VitesseY devrait être 0.4, obtenu %f", p.VitesseY)
	}
}

func TestApplique_gravite_rebond(t *testing.T) {
	p := &particle.Particle{
		PositionY: 599,
		VitesseY:  5,
		VitesseX:  2,
		ScaleX:    1,
		ScaleY:    1}

	Dy := 5.0            // taille particule en Y
	WindowSizeY := 600.0 // taille ecran
	Applique_gravite(p, Dy, WindowSizeY)
	sol := WindowSizeY - Dy*p.ScaleY // le sol pour les particules rebondissent dessus

	if p.PositionY > sol {
		t.Errorf("La particule dépasse le sol: PositionY=%f", p.PositionY)
	}
	if p.VitesseY >= 0 {
		t.Errorf("VitesseY devrait être inversée après le rebond, obtenu %f", p.VitesseY)
	}
}
