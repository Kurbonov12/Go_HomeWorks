package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name   string
	Health int
}

func attack(h Hero) Hero {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	damage := r.Intn(30)

	fmt.Printf("Атака на %s! Нанесено %d урона.\n", h.Name, damage)

	h.Health -= damage

	if h.Health <= 0 {
		h.Health = 0
		fmt.Println("Критический удар! Урон превысил остаток здоровья.")
		fmt.Println("Герой погиб!")
	}

	return h
}

func heal(h Hero) Hero {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	plusHP := r.Intn(30)

	h.Health += plusHP
	fmt.Printf("%s лечится на %d HP.\n", h.Name, plusHP)

	if h.Health > 100 {
		fmt.Println("Здоровье больше 100!")
		h.Health = 100
	}

	return h
}

func main() {
	myHero := Hero{Name: "Warrior", Health: 100}

	myHero = attack(myHero)
	myHero = heal(myHero)

	fmt.Println("Итоговое здоровье:", myHero.Health)
}