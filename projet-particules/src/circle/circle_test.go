package circle

import (
	"math"
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

func TestPositionCercle_VitesseVersCible(t *testing.T) {

	totalParticles := 4
	// Particule initiale loin de la position cible
	p := &particle.Particle{
		PositionX: 2,
		PositionY: 2,
		VitesseX:  0,
		VitesseY:  0,
	}

	Position_cercle(p, 1, totalParticles)

	// La vitesse ne doit pas être nulle
	if p.VitesseX == 0 && p.VitesseY == 0 {
		t.Errorf("La particule n'a pas bougé ou vitesse non calculée")
	}

	// Récupération du centre et du rayon
	centerX := float64(config.General.WindowSizeX) / 2
	centerY := float64(config.General.WindowSizeY) / 2
	radius := 200.0

	// Calcul de la position cible
	angleAttendu := 2.0 * math.Pi * 1 / float64(totalParticles)
	targetX := centerX + radius*math.Cos(angleAttendu)
	targetY := centerY + radius*math.Sin(angleAttendu)

	// Distance au carré avant le déplacement
	dxBefore := targetX - 2
	dyBefore := targetY - 2
	distBefore2 := dxBefore*dxBefore + dyBefore*dyBefore

	// Distance au carré après le déplacement
	newX := p.PositionX + p.VitesseX
	newY := p.PositionY + p.VitesseY
	dxAfter := targetX - newX
	dyAfter := targetY - newY
	distAfter2 := dxAfter*dxAfter + dyAfter*dyAfter

	if distAfter2 >= distBefore2 {
		t.Errorf("La particule ne se rapproche pas de la cible")
	}
}

func TestPositionCercle_VitesseNulleSiProche(t *testing.T) {
	totalParticles := 4

	centerX := float64(config.General.WindowSizeX) / 2
	centerY := float64(config.General.WindowSizeY) / 2
	radius := 200.0

	// Particule déjà proche de la cible
	angle := 0.0
	targetX := centerX + radius*math.Cos(angle)
	targetY := centerY + radius*math.Sin(angle)

	p := &particle.Particle{
		PositionX: targetX,
		PositionY: targetY,
	}

	Position_cercle(p, 0, totalParticles)

	if p.VitesseX != 0 || p.VitesseY != 0 {
		t.Errorf("La vitesse doit être nulle quand la particule est proche de sa cible")
	}
}
