package particles

import (
	"math"
	"project-particles/assets"
	"project-particles/circle"
	"project-particles/config"
	"project-particles/couleur"
	"project-particles/gravity"
	"project-particles/particle"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var spawndecimale float64

func (s *System) Update() {

	// Gestion du spawn des particules
	for spawndecimale += config.General.SpawnRate; spawndecimale >= 1; spawndecimale-- {
		g := particle.NewParticle() // création d'une nouvelle particule
		s.Content.PushBack(&g)      // ajout de la particule au systeme
	}

	// Appliquer la formation circulaire si activée
	if config.General.IsCircle {
		position := 1
		total := s.Content.Len()
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*particle.Particle)
			circle.Position_cercle(p, position, total)
			// Appliquer le mouvement vers le cercle
			p.PositionX += p.VitesseX
			p.PositionY += p.VitesseY
			position++
		}
	}

	for e := s.Content.Front(); e != nil; { // parcours de toutes les particules dans le systeme
		p := e.Value.(*particle.Particle)

		// recuperation de la taile du particule(de l'image)
		config.General.Dx = float64(assets.ParticleImage.Bounds().Dx())
		config.General.Dy = float64(assets.ParticleImage.Bounds().Dy())

		sol := float64(config.General.WindowSizeY) - config.General.Dy*p.ScaleY // position du sol
		doit_suppr := false
		next := e.Next() // sauvegarde du suivant car on peut supprimer e
		couleur.Color(p) // application de la couleur

		// application de la gravite si active (sauf en mode cercle)
		if config.General.Gravity && !config.General.IsCircle {
			gravity.Applique_gravite(p, config.General.Dy, float64(config.General.WindowSizeY))
		} else if !config.General.IsCircle { // deplacement simple sans gravite
			p.PositionX += p.VitesseX
			p.PositionY += p.VitesseY
		}

		// verif si le particule ne bouge plus et est au sol
		if p.PositionY >= sol && math.Abs(p.VitesseY) < 0.5 && math.Abs(p.VitesseX) < 0.8 {
			doit_suppr = true // marque pour suppression
		}
		if p.PositionX > float64(config.General.WindowSizeX)-config.General.Dx*p.ScaleX || // verif si hors ecran
			p.PositionX <= 0 ||
			p.PositionY > sol ||
			p.PositionY <= 0 {
			doit_suppr = true // marque pour suppression
		}

		// suppression du particule si necessaire
		if doit_suppr {
			s.Content.Remove(e) // pas couteux en liste chainee
		}
		e = next // avance au suivant

	}
	if config.General.Collison { // gestion des collisions
		Update_col(s)
	}
}
