package objects

type Manager interface {
	Object
}

// Implement Object
type GameManager struct {
	ID uint32
}

func (gm *GameManager) SetID(ID uint32) {
	gm.ID = ID
}

func (gm *GameManager) GetID() uint32 {
	return gm.ID
}

func NewGameManager() *GameManager {
	return &GameManager{}
}
