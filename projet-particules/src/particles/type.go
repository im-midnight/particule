package particles

import "container/list"

// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content *list.List
}
