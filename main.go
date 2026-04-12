package main

import (
	"bufio"
	"fmt"
	"learninggo/notes"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

const maxactions = 9

func main() {
	notes.InitDB()
	notes.Reader = reader
	notes.InitializeCacheFile()

	// Запускаем очистку кэша в отдельной goroutine
	go startCacheCleaner()

	// Спрашиваем пользователя - запустить CLI или API сервер
	fmt.Println(`
╔════════════════════════════════════╗
║     Выберите режим запуска:        ║
║  1 - CLI (интерактивный режим)     ║
║  2 - REST API (веб-сервер на :8080)║
║  3 - Оба режима (CLI + API)        ║
╚════════════════════════════════════╝
	`)

	fmt.Print("Ваш выбор: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		startCLI()
	case "2":
		startAPIServer()
	case "3":
		go startAPIServer()
		startCLI()
	default:
		fmt.Println("❌ Неверный выбор! Запускаю CLI...")
		startCLI()
	}
}

func startCacheCleaner() {
	time.Sleep(notes.GetCacheTTL())
	for {
		notes.ClearNoteFromCache()
		time.Sleep(10 * time.Second)
	}
}

func startAPIServer() {
	// Используем ReleaseMode для минимального вывода
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Группа маршрутов для API
	api := router.Group("/api")
	{
		// GET все заметки
		api.GET("/notes", notes.GetAllNotes)
		// POST новую заметку
		api.POST("/notes", notes.CreateNote)
		// GET заметку по ID
		api.GET("/notes/:id", notes.GetNoteByID)
		// PUT обновить заметку
		api.PUT("/notes/:id", notes.UpdateNoteAPI)
		// DELETE удалить заметку
		api.DELETE("/notes/:id", notes.DeleteNoteByID)
		// DELETE удалить все заметки
		api.DELETE("/notes", notes.DeleteAllNotesHandler)
		// GET статистика
		api.GET("/stats", notes.GetStats)
	}

	fmt.Println("🚀 REST API сервер запущен на http://localhost:8080")
	fmt.Println("📚 Доступные endpoints:")
	fmt.Println("  GET    /api/notes          - получить все заметки")
	fmt.Println("  POST   /api/notes          - создать новую заметку")
	fmt.Println("  GET    /api/notes/:id      - получить заметку по ID")
	fmt.Println("  PUT    /api/notes/:id      - обновить заметку")
	fmt.Println("  DELETE /api/notes/:id      - удалить заметку")
	fmt.Println("  DELETE /api/notes          - удалить все заметки")
	fmt.Println("  GET    /api/stats          - получить статистику")
	fmt.Println()

	router.Run(":8080")
}

func startCLI() {
	actions()
	for true {
		num := chooseAction()

		switch num {
		case 0:
			actions()
		case 1:
			notes.AddNote()
		case 2:
			notes.Update()
		case 3:
			notes.ShowAll()
		case 4:
			notes.ShowNoteByChoice()
		case 5:
			notes.DeleteNoteByChoice()
		case 6:
			notes.DeleteAllNotes()
		case 7:
			notes.PrintCat()
		case 8:
			notes.AutoAdd()
		case 9:
			fmt.Println("👋 До свидания!")
			os.Exit(0)
		}
	}
}

func actions() {
	fmt.Println(`
	0 - Показать действия
	1 - Добавить заметку
	2 - Изменить заметку
	3 - Показать все заметки
	4 - Показать заметку по номеру
	5 - Удалить заметку по номеру
	6 - Удалить ВСЕ
	7 - Нарисовать котика
	8 - Автоматически добавить 3 заметки
	9 - Выход
	`)
}

func chooseAction() int {
	for {
		fmt.Print("Выбрать действие🟩: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 || num > maxactions {
			fmt.Printf("❌ Пожалуйста, введите число от 1 до %d!\n", maxactions)
			continue
		}
		return num
	}
}
