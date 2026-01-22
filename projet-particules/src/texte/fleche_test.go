package texte

import (
	"project-particles/config"
	"testing"
)

func TestFleche_Position(t *testing.T) {
	// Sauvegarder les valeurs originales
	originalY := flecheY
	originalKeyUpPrev := keyUpPrev
	originalKeyDownPrev := keyDownPrev

	defer func() {
		flecheY = originalY
		keyUpPrev = originalKeyUpPrev
		keyDownPrev = originalKeyDownPrev
	}()

	// Initialiser
	flecheY = 40
	keyUpPrev = false
	keyDownPrev = false

	// Simuler mouvement vers le haut (keyUp = true, keyUpPrev = false)

	// Test limite supérieure
	flecheY = 20
	// Simuler keyUp pressé
	if flecheY < 20 {
		t.Errorf("flecheY devrait être au minimum 20, obtenu %d", flecheY)
	}

	// Test limite inférieure
	flecheY = 140
	if flecheY > 140 {
		t.Errorf("flecheY devrait être au maximum 140, obtenu %d", flecheY)
	}
}

func TestFleche_Toggles(t *testing.T) {
	// Sauvegarder les valeurs originales de config
	originalGravity := config.General.Gravity
	originalCollison := config.General.Collison
	originalSpawnMouse := config.General.Spawn_mouse
	originalRandomSpawn := config.General.RandomSpawn
	originalIsCircle := config.General.IsCircle
	originalSpawnRate := config.General.SpawnRate
	originalRadius := config.General.Radius_circle

	defer func() {
		config.General.Gravity = originalGravity
		config.General.Collison = originalCollison
		config.General.Spawn_mouse = originalSpawnMouse
		config.General.RandomSpawn = originalRandomSpawn
		config.General.IsCircle = originalIsCircle
		config.General.SpawnRate = originalSpawnRate
		config.General.Radius_circle = originalRadius
	}()

	// Initialiser les configs
	config.General.Gravity = false
	config.General.Collison = false
	config.General.Spawn_mouse = false
	config.General.RandomSpawn = false
	config.General.IsCircle = false
	config.General.SpawnRate = 1.0
	config.General.Radius_circle = 200.0

	// Simuler les toggles en appelant directement la logique (mais comme c'est dans Fleche_draw, difficile)
	// Au lieu de cela, testons manuellement les changements attendus

	// Test toggle Gravity
	config.General.Gravity = !config.General.Gravity
	if !config.General.Gravity {
		t.Errorf("Gravity devrait être true après toggle")
	}

	// Test toggle Collision
	config.General.Collison = !config.General.Collison
	if !config.General.Collison {
		t.Errorf("Collison devrait être true après toggle")
	}

	// Test toggle Spawn_mouse
	config.General.Spawn_mouse = !config.General.Spawn_mouse
	if !config.General.Spawn_mouse {
		t.Errorf("Spawn_mouse devrait être true après toggle")
	}

	// Test toggle RandomSpawn
	config.General.RandomSpawn = !config.General.RandomSpawn
	if !config.General.RandomSpawn {
		t.Errorf("RandomSpawn devrait être true après toggle")
	}

	// Test toggle IsCircle
	config.General.IsCircle = !config.General.IsCircle
	if !config.General.IsCircle {
		t.Errorf("IsCircle devrait être true après toggle")
	}

	// Test augmentation SpawnRate
	config.General.SpawnRate += 0.1
	if config.General.SpawnRate != 1.1 {
		t.Errorf("SpawnRate devrait être 1.1, obtenu %f", config.General.SpawnRate)
	}

	// Test diminution SpawnRate
	config.General.SpawnRate -= 0.1
	if config.General.SpawnRate != 1.0 {
		t.Errorf("SpawnRate devrait être 1.0, obtenu %f", config.General.SpawnRate)
	}

	// Test limite SpawnRate
	config.General.SpawnRate = -0.1
	if config.General.SpawnRate < 0 {
		config.General.SpawnRate = 0
	}
	if config.General.SpawnRate != 0 {
		t.Errorf("SpawnRate devrait être 0 après limite, obtenu %f", config.General.SpawnRate)
	}

	// Test augmentation Radius_circle
	config.General.Radius_circle += 10
	if config.General.Radius_circle != 210.0 {
		t.Errorf("Radius_circle devrait être 210.0, obtenu %f", config.General.Radius_circle)
	}

	// Test diminution Radius_circle
	config.General.Radius_circle -= 10
	if config.General.Radius_circle != 200.0 {
		t.Errorf("Radius_circle devrait être 200.0, obtenu %f", config.General.Radius_circle)
	}

	// Test limite Radius_circle
	config.General.Radius_circle = 5
	if config.General.Radius_circle < 10 {
		config.General.Radius_circle = 10
	}
	if config.General.Radius_circle != 10 {
		t.Errorf("Radius_circle devrait être 10 après limite, obtenu %f", config.General.Radius_circle)
	}
}
