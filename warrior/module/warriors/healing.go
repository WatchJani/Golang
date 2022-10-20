package warriors

import "fmt"

// Med Kit.
// First Aid Kit.
// Bandage.
// Energy Drink.
// Mega Energy Drink.
// Painkillers.
// Adrenaline Syringe.
// Hermostatic Device.

func (w *Warrior) GetHelth() Health {
	return w.Health
}

func CanUse(Health, MaxHealing Health) bool {
	if Health < 100 && Health < 100-MaxHealing {
		return true
	} else {
		return false
	}
}

func ErrorMsg(medical string) {
	fmt.Printf("Imaš previse zdravlja da bi iskoristio %s!\n", medical)
}

//PITANJE ZA OVE DVIJE FUNKCIJE

func (w *Warrior) Painkillers() {
	const MaxHealing Health = 3

	if CanUse(w.Health, MaxHealing) {
		w.Health += MaxHealing
	} else {
		ErrorMsg("Paint killer")
	}
}

func (w *Warrior) MadKit() {
	const MaxHealing Health = 80

	if CanUse(w.Health, MaxHealing) {
		w.Health += MaxHealing
	} else {
		ErrorMsg("Mad Kit")
	}
}
