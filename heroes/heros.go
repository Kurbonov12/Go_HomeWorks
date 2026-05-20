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

// Теперь функция принимает *Hero (указатель). 
// Нам больше не нужно возвращать Hero, так как изменения происходят в оригинале.
func attack(h *Hero) {
	damage := rand.Intn(30)
	fmt.Printf("Атака на %s! Нанесено %d урона.\n", h.Name, damage)

	if damage >= h.Health {
		fmt.Println("Критический удар! Урон превысил остаток здоровья.")
		h.Health = 0
	} else {
		h.Health -= damage
	}

	if h.Health <= 0 {
		fmt.Println("Герой погиб!")
	}
}

func heal(h *Hero) {
	plusHP := rand.Intn(30)
	h.Health += plusHP
	fmt.Printf("%s лечится на %d HP.\n", h.Name, plusHP)

	if h.Health > 100 {
		fmt.Println("Здоровье превысило лимит (100)! Установлен максимум.")
		h.Health = 100
	}
}

func main() {
	// Исправлено: rand.Seed нужен для того, чтобы числа были всегда разными
	rand.Seed(time.Now().UnixNano())

	myHero := Hero{Name: "Warrior", Health: 100}

	// Передаем адрес переменной с помощью символа &
	attack(&myHero)
	heal(&myHero)

	fmt.Printf("Итоговое здоровье %s: %d\n", myHero.Name, myHero.Health)
}