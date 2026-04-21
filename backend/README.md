# Backend - Go REST API

Go приложение для управления заметками с REST API + интерактивным CLI.

## 📦 Зависимости

- Go 1.25.7+
- PostgreSQL 16+
- Gin Web Framework
- lib/pq (PostgreSQL driver)

## 🏃 Запуск

### API сервер (порт 8080)
```bash
./start-api.sh
# или
go run main.go
```

### CLI режим
```bash
go run main.go
# Выбрать вариант 1 в меню
```

### Оба режима
```bash
go run main.go
# Выбрать вариант 3 в меню
```

## 📚 REST API Endpoints

| Метод | Путь | Описание |
|-------|------|---------|
| GET | `/api/notes` | Получить все заметки |
| POST | `/api/notes` | Создать заметку |
| GET | `/api/notes/:id` | Получить заметку по ID |
| PUT | `/api/notes/:id` | Обновить заметку |
| DELETE | `/api/notes/:id` | Удалить заметку |
| DELETE | `/api/notes` | Удалить все |
| GET | `/api/stats` | Статистика |

Подробнее: [API_GUIDE.md](API_GUIDE.md)

## 🏗️ Docker

```bash
docker build -t notes-backend .
docker run -p 8080:8080 notes-backend
```

## 📂 Структура

- `main.go` - точка входа
- `notes/` - пакет с бизнес-логикой (DB, handlers, команды)
- `Dockerfile` - контейнер для продакшена
