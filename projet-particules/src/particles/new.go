package particles

import (
	"container/list"
	"project-particles/config"
	"project-particles/particle"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

func NewSystem() System {
	l := list.New() // creation de la liste de particules

	var nbr int = config.General.InitNumParticles // nombre de particules a creer
	if nbr < 0 {
		nbr = 0 // evite les valeurs negatives
	}
	for i := 0; i < nbr; i++ { // creation des particules
		p := particle.NewParticle()
		l.PushFront(&p) // ajout de la particule a la liste
	}

	return System{Content: l} // retourne le systeme de particules
}
