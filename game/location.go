package game

type location struct {
	name              string
	VillainMultiplier int
	HeroMultiplier    int
	egress            []location
}
