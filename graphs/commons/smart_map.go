package commons

import "sync"

type ISmartMap[K comparable, V any] interface {
	Set(key K, value V) bool
	Remove(key K)
	RemoveBetween(first, last K)
	Get(key K) (K, V)
	GetBetween(first, last K) Pairs[K, V]

	Len() int
}

type basicSmartMap[K comparable, V any] struct {}

func (s *basicSmartMap[K, V]) Set(key K, value V) bool {
	panic("not implemented")
}

func (s basicSmartMap[K, V]) Remove(key K) {
	panic("not implemented")
}

func (s basicSmartMap[K, V]) RemoveBetween(first, last K) {
	panic("not implemented")
}

func (s basicSmartMap[K, V]) Get(key K) (K, V) {
	panic("not implemented")
}

func (s basicSmartMap[K, V]) GetBetween(first, last K) Pairs[K, V] {
	panic("not implemented")
}

func (s basicSmartMap[K, V]) Len() int {
	panic("not implemented")
}

type SmartMap[K comparable, V any] struct {
	basicSmartMap[K,V]

	m     ISmartMap[K, V]
	mutex sync.RWMutex

	mapAmount   int
	sliceAmount int
	mapMode     bool
}

func (s SmartMap[K, V]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.len()
}

func (s SmartMap[K, V]) len() int {
	return s.Len()
}

func (s SmartMap[K, V]) isConversionNeeded() bool {
	return s.isConversionNeededAt(s.len())
}

func (s SmartMap[K, V]) isConversionNeededAt(cnt int) bool {
	return (cnt > s.mapAmount && !s.mapMode) || (cnt < s.sliceAmount && s.mapMode)
}

func (s SmartMap[K, V]) convertToMap() bool {
	if !s.mapMode {
		return true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	return true
}

func (s *SmartMap[K, V]) Set(key K, value V) bool {
	return s.m.Set(key, value)
}

func (s SmartMap[K, V]) Remove(key K) {
	s.m.Remove(key)
}

func (s SmartMap[K, V]) RemoveBetween(first, last K) {
	s.m.RemoveBetween(first, last)
}

func (s SmartMap[K, V]) Get(key K) (K, V) {
	return s.m.Get(key)
}

func (s SmartMap[K, V]) GetBetween(first, last K) Pairs[K, V] {
	return s.m.GetBetween(first, last)
}

type smartList[K comparable, V any] struct {
	basicSmartMap[K,V]

	keys   map[K]int
	values []V
}

type smartMap[K comparable, V any] struct {
	basicSmartMap[K,V]

	m map[K]V
}
