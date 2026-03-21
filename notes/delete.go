package notes

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func DeleteNote(id uint) {
	result := DB.Delete(&Note{}, id)
	if result.Error != nil {
		fmt.Println("Ошибка при удалении заметки:", result.Error)
	} else if result.RowsAffected == 0 {
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
	DB.Model(&Note{}).Count(&count)
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
	result := DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&Note{})
	if result.Error != nil {
		fmt.Println("Ошибка при удалении заметок:", result.Error)
	} else {
		fmt.Printf("Удалено %d заметок!\n", result.RowsAffected)
		RemoveAllNotesFromCache()
	}
}
