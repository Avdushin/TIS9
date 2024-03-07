package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	name      string
	birthDate [3]uint32
	country   string
}

func main() {
	startTime := time.Now()

	var people []Person

	var n int
	fmt.Println("Введите количество людей (N):")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода числа. Пожалуйста, введите целое число.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Println("Введите имя:")
		scanner.Scan()
		name := scanner.Text()

		fmt.Println("Введите дату рождения (через пробел):")
		scanner.Scan()
		birthDateInput := scanner.Text()
		var birthDate [3]uint32
		birthDateInputSplit := strings.Fields(birthDateInput)
		for j := 0; j < 3; j++ {
			num, err := strconv.Atoi(birthDateInputSplit[j])
			if err != nil {
				fmt.Println("Ошибка ввода даты. Введите данные заново.")
				i--
				break
			}
			birthDate[j] = uint32(num)
		}

		fmt.Println("Введите страну проживания:")
		scanner.Scan()
		country := scanner.Text()

		person := Person{
			name:      name,
			birthDate: birthDate,
			country:   country,
		}

		people = append(people, person)
	}

	fmt.Println("Все люди:", people)

	fmt.Println("Введите значение для поиска (например, 'Россия'):")
	scanner.Scan()
	searchValue := scanner.Text()

	var filteredPeople []Person
	for _, person := range people {
		if person.country == searchValue {
			filteredPeople = append(filteredPeople, person)
		}
	}
	fmt.Println("Найденные люди:", filteredPeople)

	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println("Список людей по имени:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].birthDate[0] < people[j].birthDate[0] ||
			(people[i].birthDate[0] == people[j].birthDate[0] && people[i].birthDate[1] < people[j].birthDate[1]) ||
			(people[i].birthDate[0] == people[j].birthDate[0] && people[i].birthDate[1] == people[j].birthDate[1] && people[i].birthDate[2] < people[j].birthDate[2])
	})
	fmt.Println("Список людей по дате рождения:", people)

	//@ Измерение времени выполнения
	elapsed := time.Since(startTime)
	fmt.Printf("Время выполнения: %s\n", elapsed)

	//@ Измерение использования памяти
	f, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(f)
	f.Close()
}
