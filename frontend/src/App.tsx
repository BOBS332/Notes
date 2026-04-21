import { useState, useEffect } from 'react';
import { Note, NoteFormData } from './types';
import { NoteForm } from './components/NoteForm';
import { NotesList } from './components/NotesList';
import * as api from './api';
import './App.css';

export function App() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [editingNote, setEditingNote] = useState<Note | null>(null);
  const [editTitle, setEditTitle] = useState('');
  const [editContent, setEditContent] = useState('');

  // Загружаем заметки при загрузке компонента
  useEffect(() => {
    loadNotes();
  }, []);

  const loadNotes = async () => {
    setIsLoading(true);
    setError(null);
    const loadedNotes = await api.getAllNotes();
    setNotes(loadedNotes);
    setIsLoading(false);
  };

  const handleAddNote = async (formData: NoteFormData) => {
    setIsLoading(true);
    setError(null);
    
    const newNote = await api.createNote(formData);
    if (newNote) {
      setNotes([...notes, newNote]);
    } else {
      setError('Не удалось добавить заметку');
    }
    
    setIsLoading(false);
  };

  const handleDeleteNote = async (id: number) => {
    setIsLoading(true);
    setError(null);
    
    const success = await api.deleteNote(id);
    if (success) {
      setNotes(notes.filter((note) => note.id !== id));
    } else {
      setError('Не удалось удалить заметку');
    }
    
    setIsLoading(false);
  };

  const handleEditNote = (note: Note) => {
    setEditingNote(note);
    setEditTitle(note.title);
    setEditContent(note.content);
  };

  const handleSaveEdit = async () => {
    if (!editingNote) return;
    
    if (!editTitle.trim() || !editContent.trim()) {
      setError('Пожалуйста, заполни оба поля!');
      return;
    }

    setIsLoading(true);
    setError(null);

    const updated = await api.updateNote(editingNote.id, {
      title: editTitle,
      content: editContent,
    });

    if (updated) {
      setNotes(
        notes.map((note) =>
          note.id === editingNote.id
            ? { ...note, title: editTitle, content: editContent }
            : note
        )
      );
      setEditingNote(null);
    } else {
      setError('Не удалось обновить заметку');
    }

    setIsLoading(false);
  };

  const handleCancelEdit = () => {
    setEditingNote(null);
    setEditTitle('');
    setEditContent('');
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>📝 Notes App</h1>
        <p className="subtitle">Управление заметками</p>
      </header>

      <main className="app-main">
        <div className="container">
          {/* Форма добавления */}
          <section className="section">
            <h2>Добавить новую заметку</h2>
            <NoteForm onSubmit={handleAddNote} isLoading={isLoading} />
          </section>

          {/* Ошибка */}
          {error && (
            <div className="error-message">
              ❌ {error}
            </div>
          )}

          {/* Редактирование заметки */}
          {editingNote && (
            <section className="section edit-section">
              <h2>✏️ Редактировать заметку</h2>
              <div className="form-group">
                <label htmlFor="edit-title">Название:</label>
                <input
                  id="edit-title"
                  type="text"
                  value={editTitle}
                  onChange={(e) => setEditTitle(e.target.value)}
                  disabled={isLoading}
                />
              </div>
              <div className="form-group">
                <label htmlFor="edit-content">Содержание:</label>
                <textarea
                  id="edit-content"
                  value={editContent}
                  onChange={(e) => setEditContent(e.target.value)}
                  disabled={isLoading}
                  rows={4}
                />
              </div>
              <div className="button-group">
                <button
                  onClick={handleSaveEdit}
                  disabled={isLoading}
                  className="btn btn-primary"
                >
                  💾 Сохранить
                </button>
                <button
                  onClick={handleCancelEdit}
                  disabled={isLoading}
                  className="btn btn-secondary"
                >
                  ✕ Отмена
                </button>
              </div>
            </section>
          )}

          {/* Список заметок */}
          <section className="section">
            <div className="section-header">
              <h2>Ваши заметки ({notes.length})</h2>
              {notes.length > 0 && (
                <button
                  onClick={() => {
                    if (window.confirm('Удалить ВСЕ заметки? Это нельзя отменить!')) {
                      api.deleteAllNotes().then(() => loadNotes());
                    }
                  }}
                  disabled={isLoading}
                  className="btn btn-danger"
                >
                  Удалить все
                </button>
              )}
            </div>
            <NotesList
              notes={notes}
              onDelete={handleDeleteNote}
              onEdit={handleEditNote}
              isLoading={isLoading}
            />
          </section>
        </div>
      </main>

      <footer className="app-footer">
        <p>API: http://localhost:8080</p>
      </footer>
    </div>
  );
}
