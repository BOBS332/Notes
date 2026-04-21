package notes

import (
	"sync"
	"time"
)

var (
	cache           = make(map[uint]Note)
	cacheTTL        = 10 * time.Minute
	isCacheClearing bool
	cacheStateMu    sync.RWMutex
)

func GetNoteFromCache(id uint) (Note, bool) {
	mu.RLock()
	defer mu.RUnlock()
	note, exists := cache[id]
	return note, exists
}

func AddNoteToCache(note Note) {
	note.LastCall = time.Now()
	mu.Lock()
	cache[note.ID] = note
	mu.Unlock()
	SaveCacheToFile()
}

func RemoveNoteFromCache(id uint) {
	mu.Lock()
	delete(cache, id)
	mu.Unlock()
	SaveCacheToFile()
}

func RemoveAllNotesFromCache() {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()
	SaveCacheToFile()
}

func isNoteExpired(note Note) bool {
	duration := time.Since(note.LastCall)
	return duration > cacheTTL
}

func deleteExpiredNote() bool {
	mu.Lock()
	var expiredID uint
	var found bool
	for id, note := range cache {
		if isNoteExpired(note) {
			expiredID = id
			found = true
			break
		}
	}

	if !found {
		mu.Unlock()
		return false
	}

	delete(cache, expiredID)
	mu.Unlock()
	SaveCacheToFile()
	return true
}

func ClearNoteFromCache() {
	setCacheClearingState(true)

	for deleteExpiredNote() {
		time.Sleep(3 * time.Second)
	}

	time.Sleep(100 * time.Millisecond)

	setCacheClearingState(false)
}

func ShouldBypassCache() bool {
	cacheStateMu.RLock()
	defer cacheStateMu.RUnlock()
	return isCacheClearing
}

func setCacheClearingState(clearing bool) {
	cacheStateMu.Lock()
	defer cacheStateMu.Unlock()
	isCacheClearing = clearing
}

func GetCacheTTL() time.Duration {
	return cacheTTL
}
func SetCacheTTL(ttl time.Duration) {
	cacheTTL = ttl
}
