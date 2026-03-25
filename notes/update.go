package notes

import (
	"fmt"
	"strings"
)

func Update() {
	exists, err := ShowAvailableNotes()
	if err != nil {
		fmt.Println("Ошибка при загрузке заметок:", err)
		return
	}
	if !exists {
		return
	}

	var inp uint
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&inp)

	UpdateNote(inp)
}

func UpdateNote(id uint) {
	var note Note
	isFromCache := false

	if !ShouldBypassCache() {
		var exists bool
		note, exists = GetNoteFromCache(id)
		if exists {
			isFromCache = true
		}
	}

	if !isFromCache {
		err := GetNoteFromDB(id, &note)
		if err != nil {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
			return
		}
	}

	fmt.Printf("Введите название заметки: ")
	note.Title, _ = Reader.ReadString('\n')
	note.Title = strings.TrimSpace(note.Title)

	fmt.Printf("Введите содержание заметки: ")
	note.Content, _ = Reader.ReadString('\n')
	note.Content = strings.TrimSpace(note.Content)

	fmt.Print("Вы уверены? Данные будут обновлены. (y/n): ")
	approve, _ := Reader.ReadString('\n')
	approve = strings.TrimSpace(approve)

	if approve == "y" || approve == "Y" {
		err := SaveNoteToDB(note)
		if err != nil {
			fmt.Println("Ошибка при обновлении заметки:", err)
		} else {
			fmt.Println("Заметка успешно обновлена!")
			RemoveNoteFromCache(id)
			AddNoteToCache(note)
		}
	} else {
		fmt.Println("Обновление отменено")
	}
}
