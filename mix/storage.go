// Package source models a single audio source
package mix

import (
	"sync"
)

// Prepare a source by ensuring it is stored in memory.
func SourcePrepare(src string) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	if _, exists := storage[src]; !exists {
		storage[src] = SourceNew(src)
	}
}

// Get a source from storage
func SourceGet(src string) *Source {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	if _, ok := storage[src]; ok {
		return storage[src]
	} else {
		return nil
	}
}

func GetLength(src string) Tz {
	source := SourceGet(src)
	if source != nil {
		return source.Length()
	} else {
		return Tz(0)
	}
}

// Prune to keep only the sources in this list
func SourcePrune(keep map[string]bool) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	for key, _ := range storage {
		if _, exists := keep[key]; !exists {
			delete(storage, key)
		}
	}
}

// Count the number of sources in memory
func SourceCount() int {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	return len(storage)
}

/*
 *
 private */

var (
	storage      map[string]*Source
	storageMutex = &sync.Mutex{}
)

func init() {
	storage = make(map[string]*Source, 0)
}
