package game

type gameEvent string

const (
	villainAttack gameEvent = "Villain Attack"
	heroDefend    gameEvent = "Hero Defend"
	sleep         gameEvent = "Sleep"
	runaway       gameEvent = "Run Away"
	recharge      gameEvent = "Recharge Health"
)
