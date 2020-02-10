/*
Package keylock provides mutex locking based on a key string.

Usage
	// Zero value of KeyLock is ready to use
	kl := &KeyLock{}

	// Lock key foo
	kl.Lock("foo")
	defer kl.Unlock("foo")

	// Read lock key bar
	kl.RLock("bar")
	defer kl.RUnlock("bar")
*/
package keylock
