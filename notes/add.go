package notes

import (
	"fmt"
	"strings"
)

func AddNote() {
	fmt.Printf("Введите название заметки: ")
	title, _ := Reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Введите содержание заметки: ")
	content, _ := Reader.ReadString('\n')
	content = strings.TrimSpace(content)

	notes := LoadNotes()
	newID := getMaxID(notes) + 1

	newNote := Note{
		ID:      newID,
		Title:   title,
		Content: content,
	}

	notes = append(notes, newNote)
	err := SaveNotes(notes)

	if err != nil {
		fmt.Println("Ошибка при добавлении заметки: ", err)
	} else {
		fmt.Println("Заметка успешно добавлена!")
	}
}
