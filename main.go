package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := class()

	fmt.Println(runTraining(name, class))
}

// обратите внимание на имя функции и имена переменных
func class() string {
	var confirmation string
	var class string

	for confirmation != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &class)

		if class == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if class == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if class == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &confirmation)
		confirmation = strings.ToLower(confirmation)
	}
	return class
}

// здесь обратите внимание на имена параметров
func runTraining(name, class string) string {
	if class == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", name)
	}

	if class == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", name)
	}

	if class == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", name)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defense — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(name, class))
		}

		if cmd == "defense" {
			fmt.Println(defense(name, class))
		}

		if cmd == "special" {
			fmt.Println(special(name, class))
		}
	}

	return "тренировка окончена"
}

func attack(name, class string) string {
	switch class {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+combatModifier(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+combatModifier(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+combatModifier(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}

}

// обратите внимание на "if else" и на "else"
func defense(name, class string) string {
	switch class {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+combatModifier(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+combatModifier(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+combatModifier(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

func combatModifier(min, max int) int {
	return rand.Intn(max-min) + min
}

// обратите внимание на "if else" и на "else"
func special(name, class string) string {

	switch class {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)
	default:
		return "неизвестный класс персонажа"
	}
}
