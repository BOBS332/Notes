# 📝 Notes Application - REST API Ready

Приложение для управления заметками с REST API и CLI интерфейсом.

## 🎯 Быстрый старт

### 1. Запустить REST API сервер
```bash
./start-api.sh
```
Сервер запустится на `http://localhost:8080`

### 2. Запустить интерактивный CLI режим
```bash
./start-cli.sh
```

### 3. Запустить оба режима одновременно
```bash
./start-all.sh
```
API будет на порту 8080, CLI будет ждать команд

## 📚 REST API Endpoints

| Метод | Путь | Описание |
|-------|------|---------|
| **GET** | `/api/notes` | Получить все заметки |
| **POST** | `/api/notes` | Создать новую заметку |
| **GET** | `/api/notes/:id` | Получить заметку по ID |
| **PUT** | `/api/notes/:id` | Обновить заметку |
| **DELETE** | `/api/notes/:id` | Удалить заметку |
| **DELETE** | `/api/notes` | Удалить все заметки |
| **GET** | `/api/stats` | Получить статистику |

## 🧪 Тестирование API

### Получить все заметки
```bash
curl http://localhost:8080/api/notes | jq
```

### Создать новую заметку
```bash
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"Новая заметка","content":"Содержание"}' \
  | jq
```

### Получить заметку по ID
```bash
curl http://localhost:8080/api/notes/1 | jq
```

### Обновить заметку
```bash
curl -X PUT http://localhost:8080/api/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Обновлённо","content":"Новое"}' \
  | jq
```

### Удалить заметку
```bash
curl -X DELETE http://localhost:8080/api/notes/1
```

### Статистика
```bash
curl http://localhost:8080/api/stats | jq
```

## 📋 Структура проекта

```
LearningGo/
├── main.go                 # Главный файл с логикой запуска и маршрутизацией
├── go.mod                  # Зависимости (gin, pq)
├── go.sum                  # Хеши зависимостей
├── learninggo              # Скомпилированный исполняемый файл
│
├── notes/                  # Пакет с бизнес-логикой
│   ├── handlers.go         # REST обработчики (ручки для API)
│   ├── general.go          # Структуры и переменные
│   ├── db.go               # Работа с базой данных
│   ├── add.go              # Добавление заметок (CLI)
│   ├── show.go             # Отображение заметок (CLI)
│   ├── update.go           # Обновление заметок (CLI)
│   ├── delete.go           # Удаление заметок (CLI)
│   ├── cache.go            # Кэширование
│   ├── cat.go              # ASCII Art котика
│   └── notes_test.go       # Тесты
│
├── start-api.sh            # Скрипт для запуска API сервера
├── start-cli.sh            # Скрипт для запуска CLI режима
├── start-all.sh            # Скрипт для запуска API + CLI
├── requests.http           # Примеры запросов для REST Client
├── API_GUIDE.md            # Полная документация API
└── API_TESTED.md           # Результаты тестирования
```

## 🛠️ Что использовалось

- **язык**: Go 1.25.7
- **фреймворк**: Gin (для REST API)
- **БД**: PostgreSQL
- **кэширование**: In-memory с sync.RWMutex

## 📦 Структура заметки

```json
{
  "ID": 1,
  "Title": "Заголовок",
  "Content": "Содержание заметки",
  "CreatedAt": "2026-04-08T10:00:00Z",
  "UpdatedAt": "2026-04-08T12:00:00Z",
  "DeletedAt": null,
  "LastCall": "2026-04-08T12:00:00Z"
}
```

## 🔌 Подключение базы данных

Переменные окружения (по умолчанию):
- `DB_HOST=localhost`
- `DB_PORT=5432`
- `DB_USER=postgres`
- `DB_PASSWORD=newpassword`
- `DB_NAME=test`

Для переопределения:
```bash
DB_HOST=your_host DB_PORT=5432 DB_USER=user DB_PASSWORD=pass DB_NAME=db ./learninggo
```

## 💡 Использование в фронтенде

### React пример
```jsx
import { useState, useEffect } from 'react';

function Notes() {
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/notes')
      .then(r => r.json())
      .then(data => setNotes(data.data));
  }, []);

  const addNote = async (title, content) => {
    const res = await fetch('http://localhost:8080/api/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, content })
    });
    const newNote = await res.json();
    setNotes([...notes, newNote.note]);
  };

  return (
    <div>
      <h1>Мои заметки ({notes.length})</h1>
      {notes.map(note => (
        <div key={note.ID}>
          <h3>{note.Title}</h3>
          <p>{note.Content}</p>
        </div>
      ))}
    </div>
  );
}
```

## ✅ Проверка работы

1. Откройте новый терминал
2. Запустите сервер: `./start-api.sh`
3. В другом терминале тестируйте API запросами выше

## 📖 Дополнительная документация

- [API_GUIDE.md](API_GUIDE.md) — подробное описание всех endpoints
- [API_TESTED.md](API_TESTED.md) — результаты тестирования

## 🚀 Готово к использованию!

Ваши "ручки" (REST endpoints) полностью готовы для интеграции с фронтенд приложением!
