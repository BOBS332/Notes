package notes

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNoteFromCache(t *testing.T) {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()
	testNote := Note{
		ID:      1,
		Title:   "Тестовая заметка",
		Content: "Содержание тестовой заметки",
	}
	AddNoteToCache(testNote)

	retrievedNote, exists := GetNoteFromCache(1)
	if !exists {
		t.Errorf("Ожидалось, что заметка с ID 1 будет найдена в кэше")
	}

	if retrievedNote.ID != testNote.ID || retrievedNote.Title != testNote.Title || retrievedNote.Content != testNote.Content {
		t.Errorf("Полученная заметка не совпадает с ожидаемой. Получено: %+v, Ожидалось: %+v", retrievedNote, testNote)
	}
}

func TestAddNoteToCache(t *testing.T) {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()

	testNote := Note{
		ID:      2,
		Title:   "Другая тестовая заметка",
		Content: "Содержание другой тестовой заметки",
	}
	AddNoteToCache(testNote)

	retrievedNote, exists := GetNoteFromCache(2)
	if !exists {
		t.Errorf("Ожидалось, что заметка с ID 2 будет найдена в кэше")
	}

	if retrievedNote.ID != testNote.ID || retrievedNote.Title != testNote.Title || retrievedNote.Content != testNote.Content {
		t.Errorf("Полученная заметка не совпадает с ожидаемой. Получено: %+v, Ожидалось: %+v", retrievedNote, testNote)
	}
}

func TestRemoveNoteFromCache(t *testing.T) {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()

	testNote := Note{
		ID:      3,
		Title:   "Тестовая заметка для удаления",
		Content: "Содержание тестовой заметки для удаления",
	}
	AddNoteToCache(testNote)

	RemoveNoteFromCache(3)

	_, exists := GetNoteFromCache(3)
	if exists {
		t.Errorf("Ожидалось, что заметка с ID 3 будет удалена из кэша")
	}
}

func TestIsNoteExpired(t *testing.T) {
	testNote := Note{
		ID:       4,
		Title:    "Тестовая заметка для проверки истечения срока",
		Content:  "Содержание тестовой заметки для проверки истечения срока",
		LastCall: time.Now().Add(-10 * time.Minute),
	}

	if !isNoteExpired(testNote) {
		t.Errorf("Ожидалось, что заметка с ID 4 будет признана истекшей")
	}
}

func TestDeleteExpiredNote(t *testing.T) {
	originalTTL := GetCacheTTL()
	defer SetCacheTTL(originalTTL)

	SetCacheTTL(1 * time.Millisecond)

	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()

	testNote := Note{
		ID:       5,
		Title:    "Тестовая заметка для удаления истекшей заметки",
		Content:  "Содержание тестовой заметки для удаления истекшей заметки",
		LastCall: time.Now().Add(-10 * time.Millisecond),
	}
	mu.Lock()
	cache[testNote.ID] = testNote
	mu.Unlock()

	time.Sleep(2 * time.Millisecond)

	deleted := deleteExpiredNote()
	if !deleted {
		t.Errorf("Ожидалось, что истекшая заметка с ID 5 будет удалена из кэша")
	}

	_, exists := GetNoteFromCache(5)
	if exists {
		t.Errorf("Ожидалось, что заметка с ID 5 будет удалена из кэша после удаления истекшей заметки")
	}
}

func TestClearNoteFromCache(t *testing.T) {
	originalTTL := GetCacheTTL()
	defer SetCacheTTL(originalTTL)

	SetCacheTTL(1 * time.Millisecond)

	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()

	testNote := Note{
		ID:       8,
		Title:    "Тестовая заметка для очистки кэша",
		Content:  "Содержание тестовой заметки для очистки кэша",
		LastCall: time.Now().Add(-10 * time.Millisecond),
	}
	mu.Lock()
	cache[testNote.ID] = testNote
	mu.Unlock()

	time.Sleep(2 * time.Millisecond)

	ClearNoteFromCache()

	_, exists := GetNoteFromCache(8)
	if exists {
		t.Errorf("Ожидалось, что заметка с ID 8 будет удалена из кэша после очистки")
	}

}

func TestRemoveAllNotesFromCache(t *testing.T) {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()

	for i := 1; i <= 3; i++ {
		testNote := Note{
			ID:      uint(i),
			Title:   fmt.Sprintf("Тестовая заметка %d", i),
			Content: fmt.Sprintf("Содержание тестовой заметки %d", i),
		}
		AddNoteToCache(testNote)
	}

	RemoveAllNotesFromCache()

	for i := 1; i <= 3; i++ {
		_, exists := GetNoteFromCache(uint(i))
		if exists {
			t.Errorf("Ожидалось, что заметка с ID %d будет удалена из кэша после удаления всех заметок", i)
		}
	}
}
