import { Note, ApiResponse, NoteFormData } from './types';

const API_BASE = 'http://localhost:8080/api';

// Получить все заметки
export async function getAllNotes(): Promise<Note[]> {
  try {
    const response = await fetch(`${API_BASE}/notes`);
    if (!response.ok) throw new Error('Failed to fetch notes');
    const data: ApiResponse<Note[]> = await response.json();
    return data.data || [];
  } catch (error) {
    console.error('Error fetching notes:', error);
    return [];
  }
}

// Получить одну заметку по ID
export async function getNote(id: number): Promise<Note | null> {
  try {
    const response = await fetch(`${API_BASE}/notes/${id}`);
    if (!response.ok) throw new Error('Failed to fetch note');
    const data: ApiResponse<Note> = await response.json();
    return data.data || null;
  } catch (error) {
    console.error('Error fetching note:', error);
    return null;
  }
}

// Создать новую заметку
export async function createNote(formData: NoteFormData): Promise<Note | null> {
  try {
    const response = await fetch(`${API_BASE}/notes`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });
    if (!response.ok) throw new Error('Failed to create note');
    const data: ApiResponse<Note> = await response.json();
    return data.data || null;
  } catch (error) {
    console.error('Error creating note:', error);
    return null;
  }
}

// Обновить заметку
export async function updateNote(id: number, formData: NoteFormData): Promise<Note | null> {
  try {
    const response = await fetch(`${API_BASE}/notes/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });
    if (!response.ok) throw new Error('Failed to update note');
    const data: ApiResponse<Note> = await response.json();
    return data.data || null;
  } catch (error) {
    console.error('Error updating note:', error);
    return null;
  }
}

// Удалить заметку по ID
export async function deleteNote(id: number): Promise<boolean> {
  try {
    const response = await fetch(`${API_BASE}/notes/${id}`, {
      method: 'DELETE',
    });
    return response.ok;
  } catch (error) {
    console.error('Error deleting note:', error);
    return false;
  }
}

// Удалить все заметки
export async function deleteAllNotes(): Promise<boolean> {
  try {
    const response = await fetch(`${API_BASE}/notes`, {
      method: 'DELETE',
    });
    return response.ok;
  } catch (error) {
    console.error('Error deleting all notes:', error);
    return false;
  }
}
