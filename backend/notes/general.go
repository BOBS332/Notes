package notes

import (
	"bufio"
	"database/sql"
	"sync"
	"time"
)

var DB *sql.DB
var Reader *bufio.Reader
var mu sync.RWMutex

type Note struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	LastCall  time.Time    `json:"last_call"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
}
