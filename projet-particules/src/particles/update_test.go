package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
)

func TestSpawnParticles(t *testing.T) {
	// Test 1: Spawn 1 particule par update
	spawndecimale = 0
	config.General.SpawnRate = 1.0
	config.General.Gravity = false
	config.General.Collison = false
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	s := &System{Content: list.New()}
	s.Update()
	s.Update()
	s.Update()
	if s.Content.Len() < 3 {
		t.Errorf("Test 1: Attendu au moins 3 particules, obtenu %d", s.Content.Len())
	}
}

func TestSpawnParticles2(t *testing.T) {
	// Test 2: Spawn 2 particules par update
	spawndecimale = 0
	config.General.SpawnRate = 2.0
	config.General.Gravity = false
	config.General.Collison = false
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	s := &System{Content: list.New()}
	s.Update()
	s.Update()

	if s.Content.Len() < 4 {
		t.Errorf("Test 2: Attendu au moins 4 particules, obtenu %d", s.Content.Len())
	}
}

func TestSpawnParticles3(t *testing.T) {
	// Test 3: Spawn 0.5 particule par update
	spawndecimale = 0
	config.General.SpawnRate = 0.5
	config.General.Gravity = false
	config.General.Collison = false
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	s := &System{Content: list.New()}
	s.Update()
	s.Update()
	s.Update()
	s.Update()
	if s.Content.Len() < 2 {
		t.Errorf("Test 3: Attendu au moins 2 particules, obtenu %d", s.Content.Len())
	}
}
