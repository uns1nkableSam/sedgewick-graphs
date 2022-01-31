package commons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIteratorFromPairs(t *testing.T) {
	p := Pairs[int, string]{
		{1, "1"},
		{2, "11"},
		{3, "111"},
		{4, "1111"},
		{5, "11111"},
	}

	ks := 0
	vs := ""
	p.Iterator().Each(func(k int, v string) bool {
		ks += k
		vs += v
		return true
	})

	assert.Equal(t, 15, ks)
	assert.Equal(t, "111111111111111", vs)
}
