package texte

import (
	"fmt"

	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Texte struct {
	Screen *ebiten.Image
	Texte  string
	Y      int
}

func (t Texte) Draw() {
	var value string
	// affiche par rapport au Texte donne en parametre
	switch t.Texte {
	case "Gravity":
		value = fmt.Sprintf("%v", config.General.Gravity)
	case "Gravity_value":
		value = fmt.Sprintf("%.2f", config.General.Gravity_value)
	case "Collision":
		value = fmt.Sprintf("%v", config.General.Collison)
	case "Souris": // Spawn mouse
		value = fmt.Sprintf("%v", config.General.Spawn_mouse)
	case "Spawn Rate":
		value = fmt.Sprintf("%.2f", config.General.SpawnRate)
	case "Random Spawn":
		value = fmt.Sprintf("%v", config.General.RandomSpawn)
	case "Circle":
		value = fmt.Sprintf("%v", config.General.IsCircle)
	case "Circle_Radius":
		value = fmt.Sprintf("%.2f", config.General.Radius_circle)
	}

	// Affichage sur l'écran
	ebitenutil.DebugPrintAt(t.Screen, fmt.Sprint(t.Texte, " : ", value), 50, t.Y)
	ebitenutil.DebugPrintAt(t.Screen, fmt.Sprint(t.Texte, " : ", value), 51, t.Y)
}
