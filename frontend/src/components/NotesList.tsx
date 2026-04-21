import { Note } from '../types';

interface NotesListProps {
  notes: Note[];
  onDelete: (id: number) => void;
  onEdit: (note: Note) => void;
  isLoading?: boolean;
}

export function NotesList({ notes, onDelete, onEdit, isLoading = false }: NotesListProps) {
  if (notes.length === 0) {
    return (
      <div className="empty-state">
        <p>📝 Нет заметок</p>
        <p className="text-secondary">Добавь первую заметку выше!</p>
      </div>
    );
  }

  return (
    <div className="notes-list">
      {notes.map((note) => (
        <div key={note.id} className="note-card">
          <div className="note-header">
            <h3 className="note-title">{note.title}</h3>
            <div className="button-group">
              <button
                onClick={() => onEdit(note)}
                disabled={isLoading}
                className="btn btn-primary btn-small"
                title="Редактировать"
              >
                ✎
              </button>
              <button
                onClick={() => {
                  if (window.confirm(`Удалить заметку "${note.title}"?`)) {
                    onDelete(note.id);
                  }
                }}
                disabled={isLoading}
                className="btn btn-danger btn-small"
                title="Удалить"
              >
                ✕
              </button>
            </div>
          </div>
          <p className="note-content">{note.content}</p>
          <p className="note-date">
            {new Date(note.created_at).toLocaleString('ru-RU')}
          </p>
        </div>
      ))}
    </div>
  );
}
