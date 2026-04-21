import { useState } from 'react';
import { NoteFormData } from '../types';

interface NoteFormProps {
  onSubmit: (formData: NoteFormData) => void;
  isLoading?: boolean;
}

export function NoteForm({ onSubmit, isLoading = false }: NoteFormProps) {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!title.trim() || !content.trim()) {
      alert('Пожалуйста, заполни оба поля!');
      return;
    }

    onSubmit({ title, content });
    setTitle('');
    setContent('');
  };

  return (
    <form onSubmit={handleSubmit} className="note-form">
      <div className="form-group">
        <label htmlFor="title">Название:</label>
        <input
          id="title"
          type="text"
          placeholder="Введи название заметки..."
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          disabled={isLoading}
        />
      </div>

      <div className="form-group">
        <label htmlFor="content">Содержание:</label>
        <textarea
          id="content"
          placeholder="Введи содержание заметки..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
          disabled={isLoading}
          rows={4}
        />
      </div>

      <button type="submit" disabled={isLoading} className="btn btn-primary">
        {isLoading ? 'Добавляю...' : 'Добавить заметку'}
      </button>
    </form>
  );
}
