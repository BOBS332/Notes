package notes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllNotes(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

	notes, err := GetAllNotesByUserID(userID.(uint))
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

func CreateNote(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

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

	note.UserID = userID.(uint)
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

func GetNoteByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

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

	if note.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Доступ запрещен",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": note,
	})
}

func UpdateNoteAPI(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный ID заметки",
		})
		return
	}

	var existingNote Note
	err = GetNoteFromDB(uint(id), &existingNote)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Заметка не найдена",
		})
		return
	}

	if existingNote.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Доступ запрещен",
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

func DeleteNoteByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

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

	if note.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Доступ запрещен",
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

func DeleteAllNotesHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

	err, deletedCount := DeleteAllNotesByUserID(userID.(uint))
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

func GetStats(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неавторизованный доступ",
		})
		return
	}

	count, err := GetNotesCountByUserID(userID.(uint))
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
