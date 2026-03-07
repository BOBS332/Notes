package notes

import (
	"fmt"
	"strconv"
	"strings"
)

func DeleteNote(i int) {
	notes := LoadNotes()
	found := false

	for idx, note := range notes {
		if note.ID == i {
			notes = append(notes[:idx], notes[idx+1:]...)
			found = true
			break
		}
	}

	if !found {
		if !found {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", i)
			return
		}
	}

	err := SaveNotes(notes)

	if err != nil {
		fmt.Println("Ошибка при удалении заметки: ", err)
	} else {
		fmt.Println("Заметка успешно удалена!")
	}
}

func DeleteNoteByChoice() {
	notes := LoadNotes()

	if len(notes) == 0 {
		fmt.Println("❌ Доступный заметок нет")
		return
	}

	fmt.Println("\n========== ДОСТУПНЫЕ ЗАМЕТКИ ==========")
	for _, note := range notes {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
	}
	fmt.Println("=======================================")

	fmt.Print("Ваш выбор: ")
	inp, _ := Reader.ReadString('\n')
	inp = strings.TrimSpace(inp)

	id, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("❌ Пожалуйста, введите число!")
		return
	}

	fmt.Print("Вы уверены? (y/n): ")
	approve, _ := Reader.ReadString('\n')
	approve = strings.TrimSpace(approve)

	if approve == "y" || approve == "Y" {
		DeleteNote(id)
	} else {
		fmt.Println("Удаление отменено")
	}

}

func DeleteAllNotes() {
	notes := LoadNotes()
	if len(notes) == 0 {
		fmt.Println("Заметок нет. Удалять нечего. ¯\\_(ツ)_/¯")
		return
	}

	fmt.Print("Вы уверены? Данные будут удалены безвозвратно. (y/n): ")
	approve, _ := Reader.ReadString('\n')
	approve = strings.TrimSpace(approve)

	if approve == "y" || approve == "Y" {
		DeleteAll()
	} else {
		fmt.Println("Удаление отменено")
	}
}

func DeleteAll() {
	notes := make([]Note, 0)
	err := SaveNotes(notes)

	if err != nil {
		fmt.Println("Ошибка при удалении заметок: ", err)
	} else {
		fmt.Println("Заметки успешно удалены!")
	}
}
