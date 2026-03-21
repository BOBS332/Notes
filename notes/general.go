package notes

import (
	"bufio"
	"sync"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
var Reader *bufio.Reader
var mu sync.RWMutex

type Note struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	LastCall time.Time
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
}
