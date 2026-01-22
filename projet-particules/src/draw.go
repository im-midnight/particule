package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"
	"project-particles/texte"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particle.Particle)

		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
			screen.DrawImage(assets.ParticleImage, &options)

		}
	}

	// dessiner sur l'ecran ses fonctions dessous
	texte.Fleche_draw(screen)
	gravity := texte.Texte{Screen: screen, Texte: "Gravity", Y: 20}
	gravity.Draw()

	gravity_value := texte.Texte{Screen: screen, Texte: "Gravity_value", Y: 40}
	gravity_value.Draw()

	collision := texte.Texte{Screen: screen, Texte: "Collision", Y: 60}
	collision.Draw()

	souris := texte.Texte{Screen: screen, Texte: "Souris", Y: 80}
	souris.Draw()

	spawnRate := texte.Texte{Screen: screen, Texte: "Spawn Rate", Y: 100}
	spawnRate.Draw()

	RandomSpawn := texte.Texte{Screen: screen, Texte: "Random Spawn", Y: 120}
	RandomSpawn.Draw()

	circle := texte.Texte{Screen: screen, Texte: "Circle", Y: 140}
	circle.Draw()

	radius := texte.Texte{Screen: screen, Texte: "Circle_Radius", Y: 160}
	radius.Draw()

	if config.General.Debug { // FPS
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("FPS :", ebiten.CurrentTPS()), config.General.WindowSizeX-160, 0)
	}
	if config.General.Afficher_particules { // nombres de particules
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombres particules :", g.system.Content.Len()), config.General.WindowSizeX-160, 20)
	}

}
