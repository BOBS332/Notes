package notes

import "fmt"

func ShowAll() {
	var notes []Note

	result := DB.Find(&notes)
	if result.Error != nil {
		fmt.Println("Ошибка загрузки заметок:", result.Error)
		return
	}

	if len(notes) == 0 {
		fmt.Println("❌ Заметок нет")
		return
	}

	fmt.Println("\n===Все заметки===")
	for _, note := range notes {
		fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
		fmt.Printf("\t%s\n", note.Content)
	}
	fmt.Println("=================")
}

func ShowNote(id uint) {
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
		result := DB.First(&note, id)
		if result.Error != nil {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
			return
		}
		fmt.Printf("Заметка #%d загружена из базы данных и добавлена в кэш.\n", note.ID)
	}

	fmt.Printf("\n========== ЗАМЕТКА #%d ==========\n", note.ID)
	fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
	fmt.Printf("\t%s\n", note.Content)
	fmt.Print("================================\n\n")

	RemoveNoteFromCache(id)
	AddNoteToCache(note)
}

func ShowNoteByChoice() {
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

	ShowNote(inp)
}

func ShowAvailableNotes() (bool, error) {
	var notes []Note

	result := DB.Find(&notes)
	if result.Error != nil {
		fmt.Println("Ошибка загрузки заметок:", result.Error)
		return false, result.Error
	}

	if len(notes) == 0 {
		fmt.Println("❌ Доступных заметок нет")
		return false, nil
	}

	fmt.Println("\n========== ДОСТУПНЫЕ ЗАМЕТКИ ==========")
	for _, note := range notes {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
	}
	fmt.Println("=======================================")

	return true, nil
}
