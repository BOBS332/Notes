package notes

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "newpassword"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "test"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Ошибка при проверке подключения:", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP,
		last_call TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL DEFAULT 1 REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}

	var userCount int
	err = DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = 1").Scan(&userCount)
	if err == nil && userCount == 0 {
		_, err := DB.Exec(`
			INSERT INTO users (id, username, email, password_hash)
			VALUES (1, 'default', 'default@example.com', '$2a$10$dummy')
			ON CONFLICT (id) DO NOTHING
		`)
		if err != nil {
			fmt.Println("⚠️ Не удалось создать default пользователя:", err)
		}
	}

	fmt.Println("✅ Подключение к базе данных успешно!")
}

func AddNoteToDB(note Note) error {
	insertSQL := `
	INSERT INTO notes (title, content, last_call, created_at, updated_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := DB.Exec(insertSQL, note.Title, note.Content, note.LastCall)
	return err
}

func AddNoteToDBAndReturn(note Note) (Note, error) {
	insertSQL := `
	INSERT INTO notes (title, content, user_id, last_call, created_at, updated_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING id, created_at, updated_at, user_id
	`
	var createdNote Note
	createdNote.Title = note.Title
	createdNote.Content = note.Content
	createdNote.LastCall = note.LastCall
	createdNote.UserID = note.UserID

	err := DB.QueryRow(insertSQL, note.Title, note.Content, note.UserID).Scan(
		&createdNote.ID,
		&createdNote.CreatedAt,
		&createdNote.UpdatedAt,
		&createdNote.UserID,
	)

	return createdNote, err
}

func GetNoteFromDB(id uint, note *Note) error {
	query := `
	SELECT id, created_at, updated_at, deleted_at, last_call, title, content, user_id
	FROM notes
	WHERE id = $1 AND deleted_at IS NULL
	`
	err := DB.QueryRow(query, id).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.UpdatedAt,
		&note.DeletedAt,
		&note.LastCall,
		&note.Title,
		&note.Content,
		&note.UserID,
	)
	return err
}

func GetAllNotesFromDB() ([]Note, error) {
	query := `
	SELECT id, created_at, updated_at, deleted_at, last_call, title, content, user_id
	FROM notes
	WHERE deleted_at IS NULL
	ORDER BY id
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.ID,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.DeletedAt,
			&note.LastCall,
			&note.Title,
			&note.Content,
			&note.UserID,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func UpdateNoteInDB(note Note) error {
	updateSQL := `
	UPDATE notes
	SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP
	WHERE id = $3 AND deleted_at IS NULL
	`
	_, err := DB.Exec(updateSQL, note.Title, note.Content, note.ID)
	return err
}

func UpdateNoteInDBAndReturn(note Note) (Note, error) {
	updateSQL := `
	UPDATE notes
	SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP
	WHERE id = $3 AND deleted_at IS NULL
	`
	_, err := DB.Exec(updateSQL, note.Title, note.Content, note.ID)
	if err != nil {
		return Note{}, err
	}

	var updatedNote Note
	err = GetNoteFromDB(note.ID, &updatedNote)
	return updatedNote, err
}

func DeleteNoteFromDB(id uint) (error, int) {
	deleteSQL := `
	UPDATE notes
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = $1 AND deleted_at IS NULL
	`
	result, err := DB.Exec(deleteSQL, id)
	if err != nil {
		return err, 0
	}

	rowsAffected, err := result.RowsAffected()
	return err, int(rowsAffected)
}

func DeleteAllNotesFromDB() (error, int) {
	deleteSQL := `
	UPDATE notes
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE deleted_at IS NULL
	`
	result, err := DB.Exec(deleteSQL)
	if err != nil {
		return err, 0
	}

	rowsAffected, err := result.RowsAffected()
	return err, int(rowsAffected)
}

func SaveNoteToDB(note Note) error {
	return UpdateNoteInDB(note)
}

func GetNotesCountFromDB() (int64, error) {
	query := `
	SELECT COUNT(*)
	FROM notes
	WHERE deleted_at IS NULL
	`
	var count int64
	err := DB.QueryRow(query).Scan(&count)
	return count, err
}

func GetAllNotesByUserID(userID uint) ([]Note, error) {
	query := `
	SELECT id, created_at, updated_at, deleted_at, last_call, title, content, user_id
	FROM notes
	WHERE user_id = $1 AND deleted_at IS NULL
	ORDER BY id DESC
	`
	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.ID,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.DeletedAt,
			&note.LastCall,
			&note.Title,
			&note.Content,
			&note.UserID,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func DeleteAllNotesByUserID(userID uint) (error, int) {
	deleteSQL := `
	UPDATE notes
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE user_id = $1 AND deleted_at IS NULL
	`
	result, err := DB.Exec(deleteSQL, userID)
	if err != nil {
		return err, 0
	}

	rowsAffected, err := result.RowsAffected()
	return err, int(rowsAffected)
}

func GetNotesCountByUserID(userID uint) (int64, error) {
	query := `
	SELECT COUNT(*)
	FROM notes
	WHERE user_id = $1 AND deleted_at IS NULL
	`
	var count int64
	err := DB.QueryRow(query, userID).Scan(&count)
	return count, err
}
