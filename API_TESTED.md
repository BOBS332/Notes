# 🎉 REST API готов к использованию!

## ✅ Проверено и работает

Ваш REST API полностью функционален! Все endpoints тестированы и работают корректно.

## 📡 Примеры работающих запросов

### 1️⃣ Получить все заметки
```bash
curl http://localhost:8080/api/notes
```
**Ответ (200 OK):**
```json
{
  "count": 6,
  "data": [
    {
      "ID": 83,
      "CreatedAt": "2026-03-21T23:01:21.39994+03:00",
      "UpdatedAt": "2026-03-21T23:01:21.39994+03:00",
      "DeletedAt": null,
      "Title": "Автоматический заголовок 1",
      "Content": "Автоматическое содержание 1"
    },
    ...
  ]
}
```

### 2️⃣ Создать новую заметку
```bash
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Тестовая заметка",
    "content": "Это тест API"
  }'
```
**Ответ (201 Created):**
```json
{
  "message": "Заметка успешно создана",
  "note": {
    "Title": "Тестовая заметка",
    "Content": "Это тест API"
  }
}
```

### 3️⃣ Получить заметку по ID
```bash
curl http://localhost:8080/api/notes/83
```
**Ответ (200 OK):**
```json
{
  "data": {
    "ID": 83,
    "Title": "Автоматический заголовок 1",
    "Content": "Автоматическое содержание 1",
    "CreatedAt": "2026-03-21T23:01:21.39994+03:00",
    "UpdatedAt": "2026-03-21T23:01:21.39994+03:00"
  }
}
```

### 4️⃣ Обновить заметку
```bash
curl -X PUT http://localhost:8080/api/notes/83 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Обновлённый заголовок",
    "content": "Новое содержание"
  }'
```
**Ответ (200 OK):**
```json
{
  "message": "Заметка успешно обновлена",
  "note": {
    "ID": 83,
    "Title": "Обновлённый заголовок",
    "Content": "Новое содержание"
  }
}
```

### 5️⃣ Удалить заметку
```bash
curl -X DELETE http://localhost:8080/api/notes/83
```
**Ответ (200 OK):**
```json
{
  "message": "Заметка успешно удалена"
}
```

### 6️⃣ Получить статистику
```bash
curl http://localhost:8080/api/stats
```
**Ответ (200 OK):**
```json
{
  "total_notes": 6
}
```

## 🌐 Использование в фронтенде

### JavaScript/Fetch API
```javascript
// Получить все заметки
const response = await fetch('http://localhost:8080/api/notes');
const data = await response.json();
console.log(data);

// Создать новую заметку
const newNote = await fetch('http://localhost:8080/api/notes', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    title: 'Новая заметка',
    content: 'Содержание'
  })
});
const result = await newNote.json();
console.log(result);
```

### Axios
```javascript
// Получить все заметки
const notes = await axios.get('http://localhost:8080/api/notes');
console.log(notes.data);

// Создать заметку
const created = await axios.post('http://localhost:8080/api/notes', {
  title: 'Новая заметка',
  content: 'Содержание'
});
console.log(created.data);
```

## 🔧 Запуск сервера

### Режим 1 - Только CLI (интерактивный)
```bash
./learninggo
# Выберите 1
```

### Режим 2 - Только API сервер
```bash
./learninggo
# Выберите 2
```

### Режим 3 - CLI + API одновременно
```bash
./learninggo
# Выберите 3
```

## 📊 HTTP Статус коды

| Код | Значение | Когда возвращается |
|-----|----------|-------------------|
| **200** | OK | Успешная операция (GET, PUT) |
| **201** | Created | Заметка успешно создана (POST) |
| **400** | Bad Request | Неверные данные (пустой title/content) |
| **404** | Not Found | Заметка не найдена |
| **500** | Server Error | Ошибка базы данных |

## 🎯 Готово к использованию в фронтенде!

Теперь вы можете создавать фронтенд приложение на любом фреймворке:
- **React** - используйте hooks (useState, useEffect) с fetch или axios
- **Vue** - используйте composition API с fetch или axios
- **Angular** - используйте HttpClient service
- **HTML + JS** - простые fetch запросы
- **jQuery** - $.ajax() запросы

Все endpoints работают и готовы к интеграции! 🚀
