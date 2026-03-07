package notes

import "fmt"

func ShowAll() {
	notes := LoadNotes()

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

func ShowNote(i int) {
	notes := LoadNotes()
	for _, note := range notes {
		if note.ID == i {
			fmt.Printf("\n========== ЗАМЕТКА #%d ==========\n", note.ID)
			fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
			fmt.Printf("\t%s\n", note.Content)
			fmt.Print("================================\n\n")
			return
		}
	}
	fmt.Printf("❌ Заметка с номером %d не найдена.\n", i)
}

func ShowNoteByChoice() {
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

	var inp int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&inp)
	if inp < 1 || inp > getMaxID(notes) {
		fmt.Println("❌ Заметка не найдена")
		return
	}

	ShowNote(inp)
}
