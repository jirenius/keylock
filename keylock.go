package keylock

import "sync"

// KeyLock provides read/write mutex locking based on key name.
//
// A KeyLock must not be copied after first use.
type KeyLock struct {
	mu    sync.Mutex
	locks map[string]*lock
}

type lock struct {
	sync.RWMutex
	count uint32
}

var lockPool = sync.Pool{
	New: func() interface{} {
		return new(lock)
	},
}

// New creates a new KeyLock.
func New() *KeyLock {
	return &KeyLock{}
}

// Lock locks the mutex for the given key.
func (kl *KeyLock) Lock(key string) {
	kl.getLock(key).Lock()
}

// Unlock unlocks the mutex for the given key.
func (kl *KeyLock) Unlock(key string) {
	kl.releaseLock(key).Unlock()
}

// RLock read locks the mutex for the given key.
func (kl *KeyLock) RLock(key string) {
	kl.getLock(key).RLock()
}

// RUnlock read unlocks the mutex for the given key.
func (kl *KeyLock) RUnlock(key string) {
	kl.releaseLock(key).RUnlock()
}

func (kl *KeyLock) getLock(key string) *lock {
	kl.mu.Lock()
	if kl.locks == nil {
		kl.locks = make(map[string]*lock)
	}
	l, ok := kl.locks[key]
	if !ok {
		l = lockPool.Get().(*lock)
		kl.locks[key] = l
	}
	l.count++
	kl.mu.Unlock()
	return l
}

func (kl *KeyLock) releaseLock(key string) *lock {
	kl.mu.Lock()
	defer kl.mu.Unlock()
	l, ok := kl.locks[key]
	if !ok {
		panic("no lock for " + key + " found")
	}
	l.count--
	if l.count == 0 {
		delete(kl.locks, key)
		lockPool.Put(l)
	}
	return l
}
