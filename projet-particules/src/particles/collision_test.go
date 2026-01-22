package particles

import (
	"project-particles/particle"
	"testing"
)

// TestCollision_en_x vérifie que deux particules se percutent correctement
func TestCollision_en_x(t *testing.T) {
	// Crée deux particules
	p1 := &particle.Particle{
		PositionX: 0,
		PositionY: 0,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  5,
		VitesseY:  0,
	}
	p2 := &particle.Particle{
		PositionX: 10,
		PositionY: 0,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  -3,
		VitesseY:  0,
	}

	radius := 5.0 // rayon de chaque particule

	// Vérifie que la collision est détectée
	if !Collision(p1, p2, radius, radius) {
		t.Errorf("Collision non détectée alors que les particules se chevauchent")
	}

	// Applique la collision
	Gerer_Collision(p1, p2, radius, radius)

	// Après collision, elles devraient se séparer
	if distance(p1, p2) < radius*2 {
		t.Errorf("Les particules restent trop proches après collision")
	}

	// Vérifie que la vitesse a été modifiée
	if p1.VitesseX == 5 || p2.VitesseX == -3 {
		t.Errorf("Les vitesses n'ont pas été modifiées après la collision")
	}
}

// TestCollision_en_y vérifie que deux particules se percutent correctement
func TestCollision_en_y(t *testing.T) {
	// Crée deux particules
	p1 := &particle.Particle{
		PositionX: 0,
		PositionY: 0,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  0,
		VitesseY:  2,
	}
	p2 := &particle.Particle{
		PositionX: 0,
		PositionY: 10,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  -3,
		VitesseY:  0,
	}

	radius := 5.0 // rayon de chaque particule

	// Vérifie que la collision est détectée
	if !Collision(p1, p2, radius, radius) {
		t.Errorf("Collision non détectée alors que les particules se chevauchent")
	}

	// Applique la collision
	Gerer_Collision(p1, p2, radius, radius)

	// Après collision, elles devraient se séparer
	if distance(p1, p2) < radius*2 {
		t.Errorf("Les particules restent trop proches après collision")
	}

	// Vérifie que la vitesse a été modifiée
	if p1.VitesseY == 2 || p2.VitesseY == -3 {
		t.Errorf("Les vitesses n'ont pas été modifiées après la collision")
	}
}

func TestCollision_diagonale(t *testing.T) {
	p1 := &particle.Particle{
		PositionX: 0,
		PositionY: 0,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  2,
		VitesseY:  2,
	}
	p2 := &particle.Particle{
		PositionX: 4,
		PositionY: 4,
		ScaleX:    1,
		ScaleY:    1,
		VitesseX:  -2,
		VitesseY:  -2,
	}

	radius := 5.0

	// Vérifie que la collision est détectée
	if !Collision(p1, p2, radius, radius) {
		t.Errorf("Collision non détectée alors que les particules se chevauchent")
	}

	// Applique la collision
	Gerer_Collision(p1, p2, radius, radius)

	// Vérifie qu'elles ne se chevauchent plus
	if distance(p1, p2) < radius*2 { // tolérance pour les collisions diagonales
		t.Errorf("Les particules se chevauchent encore après collision")
	}

	// Vérifie que la vitesse a été modifiée
	if p1.VitesseX == 10 && p1.VitesseY == 5 {
		t.Errorf("Les vitesses de p1 n'ont pas été modifiées après la collision")
	}
	if p2.VitesseX == -10 && p2.VitesseY == -10 {
		t.Errorf("Les vitesses de p2 n'ont pas été modifiées après la collision")
	}
}

// TestPasDeCollision vérifie que deux particules éloignées ne rentrent pas en collision
func Test_pas_de_Collision(t *testing.T) {
	p1 := &particle.Particle{
		PositionX: 0,
		PositionY: 0,
		VitesseX:  0,
		VitesseY:  0,
	}
	p2 := &particle.Particle{
		PositionX: 20,
		PositionY: 20,
		VitesseX:  0,
		VitesseY:  0,
	}

	radius := 5.0

	if Collision(p1, p2, radius, radius) {
		t.Errorf("Collision détectée alors qu'il n'y en a pas")
	}
}

// TestCollisionAvecParticuleImmobile vérifie qu'une particule immobile est correctement affectée
func TestCollision_Immobile(t *testing.T) {
	p1 := &particle.Particle{
		PositionX: 0,
		PositionY: 0,
		VitesseX:  2,
		VitesseY:  0,
	}
	p2 := &particle.Particle{
		PositionX: 10,
		PositionY: 0,
		VitesseX:  0,
		VitesseY:  0,
	}

	radius := 5.0

	if !Collision(p1, p2, radius, radius) {
		t.Errorf("Collision avec particule immobile non détectée")
	}

	Gerer_Collision(p1, p2, radius, radius)

	if distance(p1, p2) < radius*2 {
		t.Errorf("Les particules restent trop proches après collision avec particule immobile")
	}

	if p2.VitesseX == 0 && p2.VitesseY == 0 {
		t.Errorf("La particule immobile n'a pas acquis de vitesse après collision")
	}
}
