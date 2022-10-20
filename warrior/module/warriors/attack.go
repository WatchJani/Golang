package warriors

func (w *Warrior) GetDamage(attackName string) int {
	return w.Attack[attackName]
}

//Health je tip podatka i to je sve, malo sam ga izkomplikova :D
func Damage(attacker, victim *Warrior, attackName string) {
	victim.Health -= Health(attacker.GetDamage(attackName))
}
