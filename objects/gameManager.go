package objects

type Manager interface {
	Update(layer Layer) error
}

// Implement Object
type GameManager struct{}

func (gm *GameManager) Update(layer Layer) error {

	return nil
}

func NewGameManager() *GameManager {
	return &GameManager{}
}
