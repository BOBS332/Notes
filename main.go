package main

import (
	"bufio"
	"fmt"
	"learninggo/notes"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

const maxactions = 9

func main() {
	notes.InitDB()
	notes.Reader = reader
	notes.InitializeCacheFile()
	actions()
	go func() {
		time.Sleep(notes.GetCacheTTL())
		for {
			notes.ClearNoteFromCache()
			time.Sleep(10 * time.Second)
		}
	}()
	for true {
		num := chooseAction()

		switch num {
		case 0:
			actions()
		case 1:
			notes.AddNote()
		case 2:
			notes.Update()
		case 3:
			notes.ShowAll()
		case 4:
			notes.ShowNoteByChoice()
		case 5:
			notes.DeleteNoteByChoice()
		case 6:
			notes.DeleteAllNotes()
		case 7:
			notes.PrintCat()
		case 8:
			notes.AutoAdd()
		case 9:
			fmt.Println("👋 До свидания!")
			os.Exit(0)
		}
	}

}

func actions() {
	fmt.Println(`
	0 - Показать действия
	1 - Добавить заметку
	2 - Изменить заметку
	3 - Показать все заметки
	4 - Показать заметку по номеру
	5 - Удалить заметку по номеру
	6 - Удалить ВСЕ
	7 - Нарисовать котика
	8 - Автоматически добавить 3 заметки
	9 - Выход
	`)
}

func chooseAction() int {
	for {
		fmt.Print("Выбрать действие🟩: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 || num > maxactions {
			fmt.Printf("❌ Пожалуйста, введите число от 1 до %d!\n", maxactions)
			continue
		}
		return num
	}
}
