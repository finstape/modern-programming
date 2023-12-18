package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Text    string            `json:"text"`
	Options map[string]string `json:"options"`
}

type Novel map[string]Page

func main() {
	fmt.Println("Выберите действие:")
	fmt.Println("1. Играть новеллу")
	fmt.Println("2. Создать новеллу")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		playNovel()
	case 2:
		createNovel()
	default:
		fmt.Println("Неверный выбор.")
	}
}

func playNovel() {
	fmt.Print("Введите имя файла новеллы: ")
	var filename string
	fmt.Scan(&filename)

	novel, err := loadNovel(filename)
	if err != nil {
		fmt.Println("Ошибка при загрузке новеллы:", err)
		return
	}

	fmt.Print("Проверка наличия достижимого конца... ")
	if hasReachableEnd(novel) {
		fmt.Println("Граф связный, достижимый конец есть.")
	} else {
		fmt.Println("Ошибка: граф не связный или конец не достижим.")
	}

	currentPage := "1"
	for {
		page, ok := novel[currentPage]
		if !ok {
			fmt.Println("Ошибка: Страница не найдена.")
			break
		}

		fmt.Println(page.Text)

		if len(page.Options) == 0 {
			fmt.Println("Конец новеллы.")
			break
		}

		fmt.Println("Выберите действие:")
		printNumberedOptions(page.Options)

		var userChoice int
		fmt.Print("Введите номер вашего выбора (для завершения введите 0): ")
		fmt.Scan(&userChoice)

		if userChoice == 0 {
			break
		}

		foundOption := false
		i := 1
		for key := range page.Options {
			if i == userChoice {
				currentPage = page.Options[key]
				foundOption = true
				break
			}
			i++
		}

		if !foundOption {
			fmt.Println("Ошибка: Неверный выбор.")
			break
		}
	}
}

func printNumberedOptions(options map[string]string) {
	i := 1
	for key := range options {
		fmt.Printf("%d. %s\n", i, key)
		i++
	}
}

func hasReachableEnd(novel Novel) bool {
	visited := make(map[string]bool)
	startPage := "1"

	dfs(novel, visited, startPage)

	for page := range novel {
		if !visited[page] {
			return false
		}
	}

	return true
}

func dfs(novel Novel, visited map[string]bool, currentPage string) {
	if visited[currentPage] {
		return
	}

	visited[currentPage] = true
	for _, nextPage := range novel[currentPage].Options {
		dfs(novel, visited, nextPage)
	}
}

func createNovel() {
	novel := make(Novel)

	fmt.Println("Создание новой новеллы:")

	for {
		fmt.Print("Введите номер страницы (0 для завершения): ")
		var pageNumber string
		fmt.Scan(&pageNumber)

		if pageNumber == "0" {
			break
		}

		page := Page{}

		fmt.Print("Введите текст страницы (для завершения введите 'END'):\n")
		textScanner := bufio.NewScanner(os.Stdin)
		for textScanner.Scan() {
			line := textScanner.Text()
			if line == "END" {
				break
			}
			page.Text += line + "\n"
		}

		fmt.Print("Введите количество вариантов действий: ")
		var numOptions int
		fmt.Scan(&numOptions)

		page.Options = make(map[string]string)
		fmt.Println("Введите действия и номера страниц для каждого варианта (для завершения введите 'END'):")
		for i := 0; i < numOptions; i++ {
			fmt.Printf("Действие %d: ", i+1)

			actionScanner := bufio.NewScanner(os.Stdin)
			var action string
			for actionScanner.Scan() {
				line := actionScanner.Text()
				if line == "END" {
					break
				}
				action += line + "\n"
			}

			fmt.Printf("Номер страницы, на которую ведет действие %d: ", i+1)
			var nextPage string
			fmt.Scan(&nextPage)

			page.Options[action] = nextPage
		}

		novel[pageNumber] = page
	}

	saveNovel(novel, "novel.txt")
}

func saveNovel(novel Novel, filename string) {
	data, err := json.MarshalIndent(novel, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сохранении новеллы:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении новеллы:", err)
		return
	}

	fmt.Println("Новелла успешно сохранена в файле", filename)
}

func loadNovel(filename string) (Novel, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var novel Novel
	err = json.Unmarshal(data, &novel)
	if err != nil {
		return nil, err
	}

	return novel, nil
}
