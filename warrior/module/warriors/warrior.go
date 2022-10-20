package warriors

type LvL uint8
type Health uint8

type Warrior struct {
	Name    string
	Level   LvL
	Health  Health
	Attack  map[string]int
	Defence map[string]int
	// Inventory []string
}

func New(name string, bonus Health) *Warrior {
	return &Warrior{
		Name:    name,
		Level:   0,
		Health:  100 + bonus,
		Attack:  map[string]int{"heavy attack": 25, "lite attack": 12, "bow": 5, "melee attack": 10},
		Defence: map[string]int{"foot": 2, "hand": 1, "shield": 20},
	}
}
