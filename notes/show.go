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
	result := DB.First(&note, id)
	if result.Error != nil {
		fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
		return
	}
	fmt.Printf("\n========== ЗАМЕТКА #%d ==========\n", note.ID)
	fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
	fmt.Printf("\t%s\n", note.Content)
	fmt.Print("================================\n\n")
}

func ShowNoteByChoice() {
	var notes []Note
	result := DB.Find(&notes)
	if result.Error != nil {
		fmt.Println("Ошибка загрузки заметок:", result.Error)
		return
	}

	if len(notes) == 0 {
		fmt.Println("❌ Доступный заметок нет")
		return
	}

	fmt.Println("\n========== ДОСТУПНЫЕ ЗАМЕТКИ ==========")
	for _, note := range notes {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
	}
	fmt.Println("=======================================")

	var inp uint
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&inp)

	ShowNote(inp)
}
