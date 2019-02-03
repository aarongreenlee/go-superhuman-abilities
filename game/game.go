// Package game structures the superhuman abilities game.
package game

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game
type Game struct {
	hero            *actor
	villain         []*actor
	currentLocation location
}

func New() *Game {
	return &Game{}
}

func (g *Game) Play() {
	fmt.Printf("Welcome to %s\n", gameName)
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("(press %q to quit)\n\n", "q")

	g.buildHero()
	g.buildVillains()
	g.fight()
}

func (g *Game) textInput(question string) string {
	fmt.Println(question)
	var text string
	if _, err := fmt.Scan(&text); err != nil {
		fatalError(err)
	}

	if strings.ToLower(text) == "q" {
		fmt.Printf("good bye")
		os.Exit(0)
	}

	return text
}

func (g *Game) multipleChoice(question string, choices []string) int {
	i := -1

	for i < 1 || i > len(choices) {
		fmt.Println(question)

		for i := 0; i < len(choices); i++ {
			fmt.Printf("  %d %s\n", i+1, choices[i])
		}

		var text string

		fmt.Printf("Enter your selection: ")

		if _, err := fmt.Scan(&text); err != nil {
			fatalError(err)
		}

		switch text {
		case "q", "Q":
			fmt.Printf("good bye")
			os.Exit(0)
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			i, _ = strconv.Atoi(text)
		}

		if i < 0 || i > len(choices) {
			fmt.Printf("\nTry again please...\n\n")
		}

	}

	return i
}

func (g *Game) fight() {

	herosDamage := characterCodes(g.hero.superpower) / 2
	hHealth := g.hero.health

	for _, villain := range g.villain {
		fmt.Printf("The hero has %d health and now faces %q who has %d health\n", g.hero.health, villain.name, villain.health)

		villainDamage := characterCodes(villain.superpower) / 4

		vHealth := villain.health

		for hHealth > 0 || vHealth > 0 {
			hHealth = hHealth - villainDamage
			fmt.Printf("%q used %q to deliver %d damage leaving the hero with %d health\n", villain.name, villain.superpower, villainDamage, hHealth)

			if hHealth <= 0 {
				fmt.Println("The hero lost!")
				os.Exit(0)
			}

			vHealth = vHealth - herosDamage
			if vHealth <= 0 {
				fmt.Printf("%q took the l\n", villain.name)
				break
			}
		}

		if g.hero.superpower == rapidRegeneration {
			hHealth = g.hero.health + 300
		} else {
			hHealth += 50
		}
	}

	fmt.Println("The hero won!")
}

func (g *Game) buildHero() {
	h := &actor{}
	h.name = g.textInput("What is your hero's name?")

	choice := g.multipleChoice("What is your hero's superhuman ability?", heroSuperpowers)

	// Convert the choice (1-3) to the index (0-2)
	h.superpower = heroSuperpowers[choice-1]

	h.health = characterCodes(h.name + h.superpower)

	fmt.Printf("Your hero %q will use %q to defeat the villians!\n", h.name, h.superpower)
	fmt.Printf("Your hero has a health of %d\n", h.health)

	g.hero = h
}

func (g *Game) buildVillains() {
	g.villain = make([]*actor, 0, 3)

	for i := 1; i <= 3; i++ {
		v := actor{}
		v.name = g.textInput(fmt.Sprintf("What is the name of villain number %d?", i))

		choice := g.multipleChoice("What villain #%d's superhuman ability?", villainSuperpowers)
		v.superpower = villainSuperpowers[choice-1]
		v.health = characterCodes(v.name + string(v.superpower))
		g.villain = append(g.villain, &v)
	}

	for i := 0; i < len(g.villain); i++ {
		v := g.villain[i]
		fmt.Printf("%d: %q with the superhuman ability %q and %d health\n", i+1, v.name, v.superpower, v.health)
	}
}

type actor struct {
	name       string
	superpower string
	health     int
}

func fatalError(err error) {
	fmt.Printf("ABORT: A fatal error occured: %s", err)
	os.Exit(1)
}

func characterCodes(s string) int {
	i := 0

	for _, c := range s {
		code, _ := strconv.Atoi(fmt.Sprintf("%d", c))
		i = i + code
	}

	return i
}
