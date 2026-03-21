package notes

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=localhost user=postgres password=newpassword dbname=test port=5432 sslmode=disable"

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
