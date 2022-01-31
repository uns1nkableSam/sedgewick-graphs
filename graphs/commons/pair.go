package commons

type Pair[F comparable, S any] struct {
	First F
	Second S
}

type Pairs[F comparable, S any] []Pair[F,S]

func (p Pairs[F,S]) Iterator() IIterator[F,S] {
	return NewIteratorFromPairs[F,S](p)
}