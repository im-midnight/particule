package particles

import (
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

func TestNewSystem(t *testing.T) {
	// Test 1: 0 particules
	config.General.InitNumParticles = 0
	s := NewSystem()
	if s.Content.Len() != 0 {
		t.Errorf("Test 1: Attendu 0 particules, obtenu %d", s.Content.Len())
	}

	// Test 2: 10 particules
	config.General.InitNumParticles = 10
	s = NewSystem()
	if s.Content.Len() != 10 {
		t.Errorf("Test 2: Attendu 10 particules, obtenu %d", s.Content.Len())
	}

	// Test 3: 1 particule
	config.General.InitNumParticles = 1
	s = NewSystem()
	if s.Content.Len() != 1 {
		t.Errorf("Test 3: Attendu 1 particule, obtenu %d", s.Content.Len())
	}

	// Test 4: 10000 particules
	config.General.InitNumParticles = 10000
	s = NewSystem()
	if s.Content.Len() != 10000 {
		t.Errorf("Test 4: Attendu 10000 particules, obtenu %d", s.Content.Len())
	}

	// Test 5: Verifier que le Content n'est pas nil
	config.General.InitNumParticles = 5
	s = NewSystem()
	if s.Content == nil {
		t.Errorf("Test 5: Content ne devrait pas être nil")
	}
}

func TestParticle_general(t *testing.T) {
	config.General.InitNumParticles = 10000
	s := NewSystem()

	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)

		if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX) {
			t.Errorf("Position X hors ecran: %f", p.PositionX)
		}
		if p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY) {
			t.Errorf("Positione Y hors ecran: %f", p.PositionY)
		}
		if p.VitesseX < -1 || p.VitesseX > 1 {
			t.Errorf("Vitesse X invalide: %f", p.VitesseX)
		}
		if p.VitesseY < -1 || p.VitesseY > 1 {
			t.Errorf("Vitesse Y invalide: %f", p.VitesseY)
		}
		if p.ScaleX < 1 || p.ScaleX > 2.2 {
			t.Errorf("Scale X invalide: %f", p.ScaleX)
		}
		if p.ScaleY < 1 || p.ScaleY > 2.2 {
			t.Errorf("Scale Y invalide: %f", p.ScaleY)
		}
	}
}
