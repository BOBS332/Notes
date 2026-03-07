package main

import (
	"bufio"
	"fmt"
	"notes/notes"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const maxactions = 8

func main() {
	notes.Reader = reader
	notes.InitializeNotesFile()
	actions()
	for true {
		num := chooseAction()

		switch num {
		case 1:
			notes.AddNote()
		case 2:
			notes.ShowAll()
		case 3:
			notes.ShowNoteByChoice()
		case 4:
			notes.DeleteNoteByChoice()
		case 5:
			notes.DeleteAllNotes()
		case 6:
			actions()
		case 7:
			notes.PrintCat()
		case 8:
			fmt.Println("👋 До свидания!")
			os.Exit(0)
		}
	}

}

func actions() {
	fmt.Println(`
	1 - Добавить заметку
	2 - Показать все заметки
	3 - Показать заметку по номеру
	4 - Удалить заметку по номеру
	5 - Удалить ВСЕ
	6 - Показать действия
	7 - Нарисовать котика
	8 - Выход
	`)
}

func chooseAction() int {
	for {
		fmt.Print("Выбрать действие🟩: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > maxactions {
			fmt.Printf("❌ Пожалуйста, введите число от 1 до %d!\n", maxactions)
			continue
		}
		return num
	}
}
