import { Note, ApiResponse, NoteFormData } from './types';

const API_BASE = 'http://localhost:8080/api';

// Получить токен из localStorage
function getAuthToken(): string | null {
  return localStorage.getItem('token');
}

// Получить стандартные заголовки с авторизацией
function getHeaders(): Record<string, string> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };
  
  const token = getAuthToken();
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  return headers;
}

// === AUTH ENDPOINTS ===

export interface AuthResponse {
  token: string;
  user: {
    id: number;
    username: string;
    email: string;
  };
}

export async function register(
  username: string,
  email: string,
  password: string
): Promise<AuthResponse | null> {
  try {
    const response = await fetch(`${API_BASE}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, email, password }),
    });
    if (!response.ok) throw new Error('Registration failed');
    return await response.json();
  } catch (error) {
    console.error('Error registering:', error);
    return null;
  }
}

export async function login(
  email: string,
  password: string
): Promise<AuthResponse | null> {
  try {
    const response = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    if (!response.ok) throw new Error('Login failed');
    return await response.json();
  } catch (error) {
    console.error('Error logging in:', error);
    return null;
  }
}

export function logout(): void {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
}

export function isAuthenticated(): boolean {
  return !!getAuthToken();
}

// === NOTES ENDPOINTS ===

// Получить все заметки
export async function getAllNotes(): Promise<Note[]> {
  try {
    const response = await fetch(`${API_BASE}/notes`, {
      headers: getHeaders(),
    });
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
    const response = await fetch(`${API_BASE}/notes/${id}`, {
      headers: getHeaders(),
    });
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
      headers: getHeaders(),
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
      headers: getHeaders(),
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
      headers: getHeaders(),
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
      headers: getHeaders(),
    });
    return response.ok;
  } catch (error) {
    console.error('Error deleting all notes:', error);
    return false;
  }
}

