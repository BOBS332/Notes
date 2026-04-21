# Frontend - React TypeScript

React приложение для управления заметками с красивым UI.

## 📦 Технологии

- React 18+
- TypeScript 5+
- Vite (быстрый bundler)
- CSS3
- Fetch API

## 🏃 Быстрый старт

### 1. Установить зависимости
```bash
npm install
```

### 2. Запустить dev сервер (порт 3000)
```bash
npm run dev
```
Автоматически откроется в браузере на `http://localhost:3000`

### 3. Production build
```bash
npm run build
npm run preview
```

### Docker
```bash
docker build -t notes-frontend .
docker run -p 3000:3000 notes-frontend
```

## 📂 Структура проекта

```
src/
├── main.tsx          ← Точка входа React приложения
├── index.tsx         ← Альтернативный entry point
├── App.tsx           ← Главный компонент
├── App.css           ← Стили приложения
├── types.ts          ← TypeScript типы (совпадают с бэком)
├── api.ts            ← Функции для общения с API
└── components/
    ├── NoteForm.tsx  ← Форма добавления заметки
    └── NotesList.tsx ← Список всех заметок
```

## 🔌 API интеграция

Frontend подключается к Backend API на `http://localhost:8080/api`

**Endpoints:**
- `GET /api/notes` - получить все заметки
- `POST /api/notes` - создать заметку
- `GET /api/notes/:id` - получить одну заметку
- `PUT /api/notes/:id` - обновить заметку
- `DELETE /api/notes/:id` - удалить заметку
- `DELETE /api/notes` - удалить все заметки

## 🎨 Дизайн

Минималистичный, интуитивный интерфейс:
- ✨ Modern gradient header
- 📋 Карточки заметок (grid layout)
- 🎯 Простая форма добавления
- 📱 Полностью responsive (мобильный + desktop)
- ⚡ Smooth animations и transitions

## 🧪 Тестирование локально

```bash
# Terminal 1: Backend (Go API)
cd backend
./start-api.sh

# Terminal 2: Frontend (React dev server)
cd frontend
npm run dev
```

Отворить браузер на `http://localhost:3000`
