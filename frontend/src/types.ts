// Структура одной заметки (совпадает с бэкендом)
export interface Note {
  id: number;
  title: string;
  content: string;
  created_at: string;
}

// Формат ответа от API
export interface ApiResponse<T> {
  data: T;
  count?: number;
  error?: string;
}

// Форма для добавления/редактирования заметки
export interface NoteFormData {
  title: string;
  content: string;
}
