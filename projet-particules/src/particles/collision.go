package particles

import (
	"math"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"
	"sort"
)

func distance(p1, p2 *particle.Particle) float64 {
	dx := p2.PositionX - p1.PositionX    // distance en x entre les deux particules
	dy := p2.PositionY - p1.PositionY    // distance en y entre les deux particules
	distance := math.Sqrt(dx*dx + dy*dy) // distance entre les deux particules
	return distance
}

// Collision verifie si deux particules entrent en collision
func Collision(p1, p2 *particle.Particle, collisionRadiusp1, collisionRadiusp2 float64) bool {
	radiusSum := collisionRadiusp1 + collisionRadiusp2 // somme des rayons des deux particule
	return distance(p1, p2) <= radiusSum               // retourne true si la distance est inferieure a la somme des rayons
}

// Gerer_Collision gere la collision entre deux particules
func Gerer_Collision(p1, p2 *particle.Particle, collisionRadiusp1, collisionRadiusp2 float64) {
	dx := p2.PositionX - p1.PositionX    // distance en x entre les deux particules
	dy := p2.PositionY - p1.PositionY    // distance en y entre les deux particules
	distance := math.Sqrt(dx*dx + dy*dy) // distance entre les deux particules

	const minDistance = 0.01 // pour eviter la division par zero
	if distance < minDistance {
		distance = minDistance
	}
	// Si distance = 0 (particules exactement au meme endroit)
	//On force distance = 0.01 pour eviter le crash

	// Normaliser le vecteur de collision (vecteur unitaire)  = longeur 1
	// Vecteur normal (direction de la collision)
	nx := dx / distance
	ny := dy / distance

	/*
		 ### Visualisation :
				P2 (108, 106)
				●
			   /│
			  / |  dy=6
		     /	│
		    /   │
		P1 ●────┘
		(100,100)
		dx=8

		Vecteur normal n = (0.8, 0.6)
		Cest la fleche qui va de P1 à P2, de longueur 1

	*/
	//On calcule la vitesse relative
	dvx := p2.VitesseX - p1.VitesseX //dvx = "A quelle vitesse P2 se rapproche/eloigne de P1 en X"
	dvy := p2.VitesseY - p1.VitesseY //dvy = "A quelle vitesse P2 se rapproche/eloigne de P1 en Y"
	/*
		Exemples :
		Cas 1 : Collision frontale
		P1 : VitesseX = 5,  VitesseY = 0   (va vers la droite ->)
		P2 : VitesseX = -3, VitesseY = 0   (va vers la gauche <-)

		dvx = -3 - 5 = -8  (P2 sapproche de P1 à -8 px/frame)
		dvy = 0 - 0 = 0    (pas de mouvement vertical)

		Visualisation:
		P1 -->     <-- P2
			5        -3

		Vitesse relative = -8 (elles foncent une vers l'autre)

		Cas 2 : Elles seloignent

		P1 : VitesseX = -5, VitesseY = 0  (va vers la gauche <-)
		P2 : VitesseX = 3,  VitesseY = 0  (va vers la droite ->)

		dvx = 3 - (-5) = 8  (P2 'seloigne de P1 à +8 px/frame)
		dvy = 0

		Visualisation:
		P1 <--     --> P2
			-5         3

		Vitesse relative = +8 (elles seloignent)
	*/
	//(v2x​−v1x​)*nx​+(v2y​−v1y​)*ny​ avec (v2x​−v1x​) = dvx et (v2y​−v1y​) = dvy

	produit_scal := dvx*nx + dvy*ny // produit scalaire
	if produit_scal > 0 {
		return
	}
	/*
		Le produit scalaire : dvx*nx + dvy*ny`

		la vitesse relative sur laxe de collision.
		produit_scal = "A quelle vitesse les particules se rapprochent
					DANS LA DIRECTION de la collision"
		Si produit_scal < 0 -> les particules se rapprochent car on a un angle aigu
		Si produit_scal > 0 -> les particules seloignent car on a un angle obtus

		Exemple
		P1 à (100, 100), Vitesse (5, 0)   -->
		P2 à (110, 100), Vitesse (-3, 0)  <--

		Vecteur normal:
		dx = 110 - 100 = 10
		dy = 100 - 100 = 0

		distance = 10

		nx = 10/10 = 1
		ny = 0/10 = 0

		Vecteur normal = (1, 0)  (pointe vers la droite ->)

		Vitesse relative:
		dvx = -3 - 5 = -8
		dvy = 0 - 0 = 0

		Produit scalaire: dvx*nx + dvy*ny
		= (-8) * 1 + 0 * 0 = -8

		produit_scal < 0 → Elles se rapprochent
	*/

	//ajuste le changement de vitesse,plus est grand, plus la vitesse rebondie est importante
	const coeff_rebond = 0.8
	/*
		Calcul de l'impulsion
		Formule general mais comme on a pas de masse on simplifie
		=> tous les articules ont la meme masse donc m1 = m2 = m

		v1' = v1 + ((1 + CR) × m2) / (m1 + m2) × (v2n - v1n)
		v2' = v2 + ((1 + CR) × m2) / (m1 + m2) × (v2n - v1n)
		=> v1' = v1 + (1 + CR) m / 2m × (v2n - v1n)
		=> v1' = v1 + (1 + CR) / 2 × (v2n - v1n)
		=> v1' = (1 + CR) / 2 × produit_scal
		de meme pour v2 qui aura la valeur d'impulsion dans le sens oppose
	*/

	impulse := (1 + coeff_rebond) / 2 * produit_scal

	// 3e loi de Newton :
	// P1 recoit l'impulsion dans le sens de n
	p1.VitesseX += impulse * nx
	p1.VitesseY += impulse * ny
	// P2  recoit l'impulsion dans le sens oppose de n
	p2.VitesseX -= impulse * nx
	p2.VitesseY -= impulse * ny

	overlap := collisionRadiusp1 + collisionRadiusp2 - distance

	if overlap > 0.0001 { // les particules sont imbriquées
		correction := (overlap / 2)

		p1.PositionX -= correction * nx
		p1.PositionY -= correction * ny

		p2.PositionX += correction * nx
		p2.PositionY += correction * ny
	}
}

// Interval represente un intervalle de collision pour une particule
type Interval struct {
	MinX      float64
	MaxX      float64
	rayon     float64
	Particule *particle.Particle
}

// buildIntervals construit les intervalles de collision pour chaque particule
func buildIntervals(s *System) []Interval {
	// liste des intervalles pour chaque particule
	intervals := make([]Interval, 0, s.Content.Len())

	config.General.Dx = float64(assets.ParticleImage.Bounds().Dx())
	rayon := config.General.Dx / 2

	// parcours de toutes les particules du systeme et creation des intervalles
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)

		radius := rayon * p.ScaleX
		// creation de l'intervalle pour la particule
		val := Interval{
			MinX:      p.PositionX - radius,
			MaxX:      p.PositionX + radius,
			rayon:     radius,
			Particule: p}
		// ajout de l'intervalle a la liste
		intervals = append(intervals, val)
	}
	return intervals
}

// Update_col parcourt toutes les particules du systeme et gere leurs collisions
func Update_col(s *System) {
	// construit les intervalles pour chaque particule
	intervals := buildIntervals(s)

	// trie les intervalles par leur borne minimale
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].MinX < intervals[j].MinX })

	var taille = len(intervals)
	// Parcours des intervalles pour detecter les collisions
	for i := 0; i < taille; i++ {
		var a = intervals[i]
		p1 := a.Particule
		rayon_p1 := a.rayon
		// Compare avec les intervalles suivants
		for j := i + 1; j < taille; j++ {
			var b = intervals[j]
			if b.MinX >= a.MaxX {
				break
			}
			p2 := b.Particule
			rayon_p2 := b.rayon

			// Verification de la collision
			if Collision(p1, p2, rayon_p1, rayon_p2) {
				Gerer_Collision(p1, p2, rayon_p1, rayon_p2)
			}
		}
	}
}
