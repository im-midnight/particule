package particle

import (
	"project-particles/config"
	"testing"
)

func TestNewParticle_General(t *testing.T) {
	originalRandomSpawn := config.General.RandomSpawn // sauvegarde de l'ancienne valeur

	config.General.RandomSpawn = false // desactive le spawn aleatoire
	for i := 0; i < 10000; i++ {
		result := NewParticle() // creation de la particule

		expectedX := float64(config.General.WindowSizeX) / 2
		expectedY := float64(config.General.WindowSizeY) / 2

		if result.PositionX != expectedX {
			t.Errorf("PositionX attendu: %f, obtenu: %f", expectedX, result.PositionX)
		}
		if result.PositionY != expectedY {
			t.Errorf("PositionY attendu: %f, obtenu: %f", expectedY, result.PositionY)
		}
		if result.PositionX < 0 || result.PositionX > float64(config.General.WindowSizeX) {
			t.Errorf("Position X hors ecran: %f", result.PositionX)
		}
		if result.PositionY < 0 || result.PositionY > float64(config.General.WindowSizeY) {
			t.Errorf("Positione Y hors ecran: %f", result.PositionY)
		}
		if result.Rotation != 0 {
			t.Errorf("Rotation attendu: 0, obtenu: %f", result.Rotation)
		}
		if result.ScaleX < 1 || result.ScaleX > 2.2 {
			t.Errorf("ScaleX attendu: entre 1 et 2.2 , obtenu: %f", result.ScaleX)
		}
		if result.ScaleY < 1 || result.ScaleY > 2.2 {
			t.Errorf("ScaleY attendu: entre 1 et 2.2, obtenu: %f", result.ScaleY)
		}
		if result.Opacity != 1 {
			t.Errorf("Opacity attendu: 1, obtenu: %f", result.Opacity)
		}

		if result.ColorRed < 0 || result.ColorRed > 1 {
			t.Errorf("ColorRed doit etre entre 0 et 1 mais il est de : %f", result.ColorRed)
		}
		if result.ColorGreen < 0 || result.ColorGreen > 1 {
			t.Errorf("ColorGreen doit etre entre 0 et 1 mais il est de: %f", result.ColorGreen)
		}
		if result.ColorBlue < 0 || result.ColorBlue > 1 {
			t.Errorf("ColorBlue doit etre entre 0 et 1 mais il est de: %f", result.ColorBlue)
		}

		if result.VitesseX < -1 || result.VitesseX > 1 {
			t.Errorf("VitesseX doit etre entre -1 et 1.5 mais il est de : %f", result.VitesseX)
		}
		if result.VitesseY < -1 || result.VitesseY > 1 {
			t.Errorf("VitesseY doit etre entre -1 et 1.5 mais il est de : %f", result.VitesseY)
		}

	}

	config.General.RandomSpawn = originalRandomSpawn // restauration de l'ancienne valeur
}

func TestNewParticle_RandomSpawn(t *testing.T) {
	originalRandomSpawn := config.General.RandomSpawn

	config.General.RandomSpawn = true

	result := NewParticle()

	if result.PositionX < 0 || result.PositionX > float64(config.General.WindowSizeX) {
		t.Errorf("PositionX hors limites: %f", result.PositionX)
	}
	if result.PositionY < 0 || result.PositionY > float64(config.General.WindowSizeY) {
		t.Errorf("PositionY hors limites: %f", result.PositionY)
	}

	if result.Rotation != 0 {
		t.Errorf("Rotation attendu: 0, obtenu: %f", result.Rotation)
	}
	if result.ScaleX < 1 || result.ScaleX > 2.2 {
		t.Errorf("ScaleX attendu: entre 1 et 2.2 , obtenu: %f", result.ScaleX)
	}
	if result.ScaleY < 1 || result.ScaleY > 2.2 {
		t.Errorf("ScaleY attendu: entre 1 et 2.2, obtenu: %f", result.ScaleY)
	}
	if result.Opacity != 1 {
		t.Errorf("Opacity attendu: 1, obtenu: %f", result.Opacity)
	}

	config.General.RandomSpawn = originalRandomSpawn
}
func TestNewParticle_Vitessse(t *testing.T) {
	for i := 0; i < 1000; i++ { // test sur 1000 particules
		result := NewParticle()
		// verif de vitesse
		if result.VitesseX < -1 || result.VitesseX > 1 {
			t.Errorf("VitesseX nest pas dans [-1, 1]: %f", result.VitesseX)
		}
		if result.VitesseY < -1 || result.VitesseY > 1 {
			t.Errorf("VitesseY nest pas dans [-1, 1]: %f", result.VitesseY)
		}
	}
}
