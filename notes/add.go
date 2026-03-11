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

	newNote := Note{
		Title:   title,
		Content: content,
	}

	err := DB.Create(&newNote).Error
	if err != nil {
		fmt.Println("Ошибка при добавлении заметки: ", err)
	} else {
		fmt.Println("Заметка успешно добавлена!")
	}
}
