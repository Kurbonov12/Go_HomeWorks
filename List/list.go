package main

import (
	"fmt"
	"os"
)

// Предварительно создаем слайсы под категории продуктов
var (
	fruits     = []string{"Яблоки", "Бананы", "Апельсины"}
	vegetables = []string{"Картофель", "Огурцы", "Помидоры"}
	dairy      = []string{"Молоко", "Творог", "Сыр"}
)

// Функция создания списка (добавления предмета в нужный слайс)
// Принимает номер группы и название покупки
func addItem(categoryIdx int, itemName string) {
	switch categoryIdx {
	case 1:
		fruits = append(fruits, itemName)
		fmt.Printf("✅ %s успешно добавлен в категорию 'Фрукты'\n", itemName)
	case 2:
		vegetables = append(vegetables, itemName)
		fmt.Printf("✅ %s успешно добавлен в категорию 'Овощи'\n", itemName)
	case 3:
		dairy = append(dairy, itemName)
		fmt.Printf("✅ %s успешно добавлен в категорию 'Молочные продукты'\n", itemName)
	default:
		fmt.Println("❌ Ошибка: неверный номер категории!")
	}
}

// Функция чтения списка (вывод покупок из конкретной категории)
func printCategoryList(categoryIdx int) {
	fmt.Println("\n--- Список покупок ---")
	switch categoryIdx {
	case 1:
		fmt.Println("🍎 Категория 'Фрукты':")
		printSlice(fruits)
	case 2:
		fmt.Println("🥦 Категория 'Овощи':")
		printSlice(vegetables)
	case 3:
		fmt.Println("🥛 Категория 'Молочные продукты':")
		printSlice(dairy)
	default:
		fmt.Println("❌ Ошибка: категория не найдена!")
	}
	fmt.Println("----------------------")
}

// Вспомогательная функция для красивого вывода элементов слайса
func printSlice(slice []string) {
	if len(slice) == 0 {
		fmt.Println("  (список пуст)")
		return
	}
	for i, item := range slice {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
}

func main() {
	for {
		// Главное меню в терминале
		fmt.Println("\n--- ГЛАВНОЕ МЕНЮ ---")
		fmt.Println("1. Посмотреть список покупок")
		fmt.Println("2. Добавить покупку в список")
		fmt.Println("3. Выйти из программы")
		fmt.Print("Выберите действие (1-3): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Меню чтения списка
			fmt.Println("\nДоступные категории:")
			fmt.Println("1. Фрукты")
			fmt.Println("2. Овощи")
			fmt.Println("3. Молочные продукты")
			fmt.Print("Выберите номер категории для просмотра: ")

			var catChoice int
			fmt.Scan(&catChoice)
			printCategoryList(catChoice)

		case 2:
			// Меню добавления в список
			fmt.Println("\nДоступные категории для добавления:")
			fmt.Println("1. Фрукты")
			fmt.Println("2. Овощи")
			fmt.Println("3. Молочные продукты")
			fmt.Print("Введите номер категории: ")

			var catChoice int
			fmt.Scan(&catChoice)

			if catChoice < 1 || catChoice > 3 {
				fmt.Println("❌ Неверный номер категории.")
				continue
			}

			fmt.Print("Введите название покупки: ")
			var itemName string
			fmt.Scan(&itemName) // Считает одно слово. Если нужны пробелы, лучше использовать bufio.NewScanner

			// Вызов функции добавления
			addItem(catChoice, itemName)

		case 3:
			fmt.Println("Выход из программы. Хорошего дня!")
			os.Exit(0)

		default:
			fmt.Println("❌ Неверный пункт меню. Попробуйте еще раз.")
		}
	}
}
