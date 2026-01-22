# Projet Particules

SAE particules en Go - PHOMASONE KEVIN et  KERLOCH NOE

## Structure Principale du Projet
```
src/
├── assets/           # Gestion des ressources (images)
├── circle/          # Formation circulaire des particules
├── config/          # Configuration du projet
├── couleur/         # Gestion des couleurs
├── gravity/         # Système de gravité
├── particle/        # Création des particules
├── particles/       # Système de particules et collisions
├── texte/           # Affichage de texte sur l'ecran
├── draw.go          # Rendu graphique
├── game.go          # Logique du jeu
├── main.go          # Point d'entrée
└── update.go        # Mise à jour de l'état
```
### EXTENSIONS:
- `Gravity` - 5.1
- `Exterieur de l'ecran` - 5.2
- `Variation de Couleur,..` - 5.4
- `Forme du Generateur` - 5.5
- `Générateur mobile` - 5.6
- `Modification dynamique du système de particules` - 5.7
- `Collisions` - 5.10

### Particles (Collisions)
Gère la détection et la résolution des collisions entre particules.

**Fonctions principales:**
- `Collision(p1, p2, radius1, radius2)` - Détecte si deux particules entrent en collision
- `Gerer_Collision(p1, p2, radius1, radius2)` - Gère la collision des particules
- `Update_col(system)` - Met à jour toutes les collisions du système

### Affichage sur l'ecran
- `'W' ,'S' pour bouger la fleche, 'Enter' pour modifer la valeur ( ou augmenter la valeur), 'Shift' pour diminuer la valeur`

**Nombres de particules :**
- `ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombres particules :", g.system.Content.Len()), config.General.WindowSizeX-160, 20)` - Affiche le nombres de particules

**Affichage Modifiable:**
- `Gravity`: Enter pour appliquer la gravity et inverse
- `Gravity_value`: Enter pour augmenter la gravity et shift pour diminuer
- `Collision`: Enter pour appliquer les collisions et inverse
- `Souris`: Enter pour que les particules spawn a la positon de la souris et inverse
- `Spawn Rate`: Enter pour augmenter le spawn rate et Shift pour diminuner
- `Random Spawn`: Enter pour appliquer Random Spawn et inverse
- `Circle`: Enter pour appliquer la forme du cercle et inverse
- `Circle_Radius`: Enter pour augmenter le rayon du cercle et shift pour diminuer

### Tout les Tests disponible

- `collision_test.go` - Test de collision
- `new_test.go` - 2 fichiers Test pour les Particules, 1 avec System et l'autre sans
- `update_test.go` - Test mise à jour de l'état du système de particule
- `gravity_test.go` - Test de gravity
- `color_test.go` - Test pour les couleurs
- `circle_test.go` - Test pour l'affichage du cercle

