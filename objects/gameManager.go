package objects

type Manager interface {
	Update() error
}

// Implement Object
type GameManager struct {
	Layers Layers
}

func (gm *GameManager) Update() error {
	return nil
}

func NewGameManager(layers *Layers) *GameManager {
	return &GameManager{
		Layers: *layers,
	}
}
