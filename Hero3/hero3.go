package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name   string
	Health int
}

func (h *Hero) Attack() {
	damage := rand.Intn(50) + 1
	fmt.Println(h.Name, "Получает удар!", damage)

	if damage > h.Health {
		fmt.Println(" Урон превысил остаток здоровья!")
		h.Health = 0
	} else {
		h.Health = h.Health - damage
	}
}

func (h *Hero) Heal() {
	healAmount := rand.Intn(50) + 1
	fmt.Println(h.Name, "Плюс к здоровью:", healAmount)

	if h.Health+healAmount > 100 {
		fmt.Println(" Здоровье не может быть больше 100.")
		h.Health = 100
	} else {
		h.Health = h.Health + healAmount
	}
}

func main() {
	myHero := Hero{
		Name:   "Марио",
		Health: 100,
	}

	fmt.Println(" Наш герой:", myHero.Name, "| Здоровье:", myHero.Health)

	myHero.Attack()
	fmt.Println(" Текущее здоровье:", myHero.Health)

	myHero.Heal()
	fmt.Println(" Текущее здоровье:", myHero.Health)
}