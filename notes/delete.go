package notes

import (
	"fmt"
	"strconv"
	"strings"
)

func DeleteNote(id uint) {
	err, deletedCount := DeleteNoteFromDB(id)
	if err != nil {
		fmt.Println("Ошибка при удалении заметки:", err)
	} else if deletedCount == 0 {
		fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
	} else {
		fmt.Println("Заметка успешно удалена!")
		RemoveNoteFromCache(id)
	}
}

func DeleteNoteByChoice() {
	exists, err := ShowAvailableNotes()
	if err != nil {
		fmt.Println("Ошибка при загрузке заметок:", err)
		return
	}
	if !exists {
		return
	}

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
		DeleteNote(uint(id))
	} else {
		fmt.Println("Удаление отменено")
	}
}

func DeleteAllNotes() {
	var count int64
	count, err := GetNotesCountFromDB()
	if err != nil {
		fmt.Println("Ошибка при подсчете заметок:", err)
		return
	}
	if count == 0 {
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
	err, deletedCount := DeleteAllNotesFromDB()
	if err != nil {
		fmt.Println("Ошибка при удалении заметок:", err)
	} else {
		fmt.Printf("Успешно удалено %d заметок!\n", deletedCount)
		RemoveAllNotesFromCache()
	}
}
