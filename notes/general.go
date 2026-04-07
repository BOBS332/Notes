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
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	LastCall  time.Time
	Title     string
	Content   string
}
