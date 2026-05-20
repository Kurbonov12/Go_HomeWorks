package main

import (
	"fmt"
)

func AddGrade(grades *[5]int, grade int, index int) {
	if index < 0 || index >= 5 {
		fmt.Println(" Ошибка: Допустимо от 0 до 4.")
		return
	}

	if grade > 5 {
		fmt.Println(" Ошибка: Оценка выше 5 баллов!")
		return
	}

	grades[index] = grade
	fmt.Println(" Оценка успешно добавлена:", grade)
}

func PrintStats(grades [5]int) {
	sum := 0

	for i := 0; i < len(grades); i++ {
		sum = sum + grades[i]
	}

	average := float64(sum) / float64(len(grades))

	fmt.Println("\n Оценки студента:", grades)
	fmt.Println(" Сумма оценок:", sum)
	fmt.Println(" Средняя оценка:", average)
}

func main() {
	studentGrades := [5]int{}

	AddGrade(&studentGrades, 5, 0)
	AddGrade(&studentGrades, 4, 1)
	AddGrade(&studentGrades, 10, 2)
	AddGrade(&studentGrades, 3, 2)
	AddGrade(&studentGrades, 5, 3)
	AddGrade(&studentGrades, 2, 4)

	PrintStats(studentGrades)
}
