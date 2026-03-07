package notes

import (
	"bufio"
	"encoding/json"
	"os"
)

var Reader *bufio.Reader

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

const notesFile = "notes.json"

func InitializeNotesFile() {
	if _, err := os.Stat(notesFile); os.IsNotExist(err) {
		notes := make([]Note, 0)
		data, _ := json.MarshalIndent(notes, "", "  ")
		os.WriteFile(notesFile, data, 0644)
	}
}

func LoadNotes() []Note {
	data, err := os.ReadFile(notesFile)
	if err != nil {
		return []Note{}
	}

	notes := make([]Note, 0)
	json.Unmarshal(data, &notes)
	return notes
}

func SaveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(notesFile, data, 0644)
}

func getMaxID(notes []Note) int {
	maxID := 0
	for _, note := range notes {
		if note.ID > maxID {
			maxID = note.ID
		}
	}
	return maxID
}
