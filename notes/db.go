package notes

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=db user=notes_user password=notes_password dbname=notes_db port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	err = DB.AutoMigrate(&Note{})
	if err != nil {
		log.Fatal("Ошибка при миграции базы данных:", err)
	}

	fmt.Println("✅ Подключение к базе данных успешно!")
}

func AddNoteToDB(note Note) error {
	return DB.Create(&note).Error
}

func GetNoteFromDB(id uint, note *Note) error {
	result := DB.First(note, id)
	return result.Error
}

func GetAllNotesFromDB() ([]Note, error) {
	var notes []Note
	result := DB.Find(&notes)
	return notes, result.Error
}

func UpdateNoteInDB(note Note) error {
	return DB.Save(&note).Error
}

func DeleteNoteFromDB(id uint) (error, int) {
	result := DB.Delete(&Note{}, id)
	return result.Error, int(result.RowsAffected)
}

func DeleteAllNotesFromDB() (error, int) {
	result := DB.Where("1 = 1").Delete(&Note{})
	return result.Error, int(result.RowsAffected)
}

func SaveNoteToDB(note Note) error {
	return DB.Save(&note).Error
}

func GetNotesCountFromDB() (int64, error) {
	var count int64
	result := DB.Model(&Note{}).Count(&count)
	return count, result.Error
}
