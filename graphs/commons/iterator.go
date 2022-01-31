package commons

type IIterator[K comparable, V any] interface {
	Next() bool
	EOF() bool
	Key() K
	Value() V
	Data() (K,V)
    Each(f func(K,V) bool)
}

var _ IIterator[int, string] = &Iterator[int,string]{}

type Iterator[K comparable, V any] struct {
	currentElem K
	hasStarted  bool
}

func NewSliceIterator[V any](s *[]V) IIterator[int, V] {
	result := &SliceIterator[int, V]{
		currentElem: 0,
		hasStarted:  false,
	}
	return result
}

type SliceIterator [K comparable, V any] struct {
	values *[]V
}

/*
func NewIteratorFromPairs[K comparable, V any](m Pairs[K, V]) IIterator[K, V] {
	result := &Iterator[K, V]{
		keys:   make([]K, len(m)),
		values: make([]V, len(m)),

		currentElem: 0,
		hasStarted:  false,
	}

	for n, e := range m {
		result.keys[n] = e.First
		result.values[n] = e.Second
	}

	return result
}

func NewIteratorFromMap[K comparable, V any](m map[K]V) IIterator[K, V] {
	result := &Iterator[K, V]{
		keys:   make([]K, len(m)),
		values: make([]V, len(m)),

		currentElem: 0,
		hasStarted:  false,
	}

	n := 0
	for k, v := range m {
		result.keys[n] = k
		result.values[n] = v
		n++
	}

	return result
}

func (i *Iterator[K, V]) Next() bool {
	if i.hasStarted {
		i.currentElem++
	} else {
		i.hasStarted = true
	}
	return !i.EOF()
}

func (i Iterator[K, V]) Key() K {
	if i.keys == nil {
		var key K
		return key
	}
	return i.keys[i.currentElem]

}

func (i Iterator[K, V]) Value() V {
	return i.values[i.currentElem]
}

func (i Iterator[K, V]) Data() (K, V) {
	return i.Key(), i.values[i.currentElem]
}

func (i Iterator[K, V]) EOF() bool {
	return i.currentElem >= len(i.values)
}

func (i *Iterator[K, V]) Each(f func(k K, v V) bool) {
	for i.Next() {
		if !f(i.Data()) {
			return
		}
	}
}
*/
