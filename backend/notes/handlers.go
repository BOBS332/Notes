package notes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllNotes - получить все заметки
// GET /api/notes
func GetAllNotes(c *gin.Context) {
	notes, err := GetAllNotesFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при загрузке заметок",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  notes,
		"count": len(notes),
	})
}

// CreateNote - создать новую заметку
// POST /api/notes
func CreateNote(c *gin.Context) {
	var note Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	if note.Title == "" || note.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Название и содержание не могут быть пустыми",
		})
		return
	}

	createdNote, err := AddNoteToDBAndReturn(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при добавлении заметки",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdNote,
	})
}

// GetNoteByID - получить заметку по ID
// GET /api/notes/:id
func GetNoteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный ID заметки",
		})
		return
	}

	var note Note
	err = GetNoteFromDB(uint(id), &note)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Заметка не найдена",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": note,
	})
}

// UpdateNoteAPI - обновить заметку
// PUT /api/notes/:id
func UpdateNoteAPI(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный ID заметки",
		})
		return
	}

	var updateData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	if updateData.Title == "" || updateData.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Название и содержание не могут быть пустыми",
		})
		return
	}

	note := Note{
		ID:      uint(id),
		Title:   updateData.Title,
		Content: updateData.Content,
	}

	updatedNote, err := UpdateNoteInDBAndReturn(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при обновлении заметки",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedNote,
	})
}

// DeleteNoteByID - удалить заметку
// DELETE /api/notes/:id
func DeleteNoteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный ID заметки",
		})
		return
	}

	err, deletedCount := DeleteNoteFromDB(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при удалении заметки",
		})
		return
	}

	if deletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Заметка не найдена",
		})
		return
	}

	RemoveNoteFromCache(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"message": "Заметка успешно удалена",
	})
}

// DeleteAllNotesHandler - удалить все заметки
// DELETE /api/notes
func DeleteAllNotesHandler(c *gin.Context) {
	err, deletedCount := DeleteAllNotesFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при удалении заметок",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Все заметки успешно удалены",
		"deleted_count": deletedCount,
	})
}

// GetStats - получить статистику
// GET /api/stats
func GetStats(c *gin.Context) {
	count, err := GetNotesCountFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении статистики",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_notes": count,
	})
}
