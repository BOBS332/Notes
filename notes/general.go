package notes

import (
	"bufio"
	"log"
)

var Reader *bufio.Reader

func LoadNotes() []Note {
	var notes []Note
	result := DB.Find(&notes)
	if result.Error != nil {
		log.Println("Error loading notes:", result.Error)
		return []Note{}
	}
	return notes
}

func SaveNotes(notes []Note) error {
	// В GORM сохранение происходит через Create/Update, но для простоты перезапишем
	// Сначала удалим все, потом добавим новые
	DB.Unscoped().Delete(&Note{}) // Удаляем все (включая soft delete)
	for _, note := range notes {
		if err := DB.Create(&note).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetNoteByID(id int) (Note, bool) {
	var note Note
	result := DB.First(&note, id)
	if result.Error != nil {
		return Note{}, false
	}
	return note, true
}
