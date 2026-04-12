# 🚀 REST API для Notes - Инструкция по использованию

Теперь у вас есть полнофункциональный REST API для управления заметками!

## 📝 Что было добавлено

### 1. **Файл handlers.go** (`notes/handlers.go`)
   - REST обработчики для всех операций с заметками
   - Автоматическая валидация JSON данных
   - Обработка ошибок с правильными HTTP статусами

### 2. **Обновлённый main.go**
   - Выбор режима запуска (CLI, API или оба)
   - Запуск Gin веб-сервера на порту 8080

## 📚 REST API Endpoints

| Метод | Endpoint | Описание |
|-------|----------|---------|
| **GET** | `/api/notes` | Получить все заметки |
| **POST** | `/api/notes` | Создать новую заметку |
| **GET** | `/api/notes/:id` | Получить заметку по ID |
| **PUT** | `/api/notes/:id` | Обновить заметку |
| **DELETE** | `/api/notes/:id` | Удалить заметку по ID |
| **DELETE** | `/api/notes` | Удалить все заметки |
| **GET** | `/api/stats` | Получить статистику |

## 🧪 Примеры запросов (curl)

### 1. Получить все заметки
```bash
curl -X GET http://localhost:8080/api/notes
```

### 2. Создать новую заметку
```bash
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Моя первая заметка",
    "content": "Это содержание заметки"
  }'
```

### 3. Получить заметку по ID
```bash
curl -X GET http://localhost:8080/api/notes/1
```

### 4. Обновить заметку
```bash
curl -X PUT http://localhost:8080/api/notes/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Обновлённый заголовок",
    "content": "Новое содержание"
  }'
```

### 5. Удалить заметку
```bash
curl -X DELETE http://localhost:8080/api/notes/1
```

### 6. Удалить все заметки
```bash
curl -X DELETE http://localhost:8080/api/notes
```

### 7. Получить статистику
```bash
curl -X GET http://localhost:8080/api/stats
```

## 🚀 Запуск приложения

```bash
cd /home/kilobytik/LearningGo
./learninggo
```

Выберите режим:
- **1** - только CLI (интерактивный режим)
- **2** - только REST API (веб-сервер)
- **3** - оба режима (API + CLI в фоне)

## 💡 Примеры с PostMan или VS Code REST Client

Для VS Code можно установить расширение **REST Client** и создать файл `requests.http`:

```http
### Получить все заметки
GET http://localhost:8080/api/notes

### Создать заметку
POST http://localhost:8080/api/notes
Content-Type: application/json

{
  "title": "Новая заметка",
  "content": "Содержание заметки"
}

### Получить заметку по ID
GET http://localhost:8080/api/notes/1

### Обновить заметку
PUT http://localhost:8080/api/notes/1
Content-Type: application/json

{
  "title": "Обновлённая заметка",
  "content": "Новое содержание"
}

### Удалить заметку
DELETE http://localhost:8080/api/notes/1

### Получить статистику
GET http://localhost:8080/api/stats
```

## 🎯 Следующие шаги для фронтенда

Теперь ваш фронтенд может использовать эти endpoints для:

1. **Получения данных** с помощью `fetch()` или `axios`
2. **Отправки новых заметок** через POST запросы
3. **Обновления заметок** через PUT запросы
4. **Удаления заметок** через DELETE запросы
5. **Отображения данных** в таблице, карточках или других элементах UI

## 📝 Ответы API в формате JSON

### Успешный ответ (GET /api/notes):
```json
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-04-08T12:00:00Z",
      "UpdatedAt": "2026-04-08T12:00:00Z",
      "DeletedAt": null,
      "LastCall": "2026-04-08T12:00:00Z",
      "Title": "Моя заметка",
      "Content": "Содержание"
    }
  ],
  "count": 1
}
```

## ✅ Что дальше?

1. **Добавить CORS** для работы с фронтенд приложением
2. **Добавить аутентификацию** (JWT токены)
3. **Добавить фильтрацию и поиск** по заметкам
4. **Добавить пагинацию** для больших списков
5. **Создать фронтенд** (React, Vue, Angular и т.д.)

Ваши "ручки" готовы! 🎉
