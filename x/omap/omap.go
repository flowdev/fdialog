// Package omap implements a very simple ordered map with just the absolute
// minimum features for our purpose.
// It's not very performant of whatever, but it's nice to use and easy to maintain.
package omap

import "iter"

type OrderedMap[K comparable, V any] struct {
	m map[K]V
	s []K
}

// New returns a new ordered map with space for size elements.
func New[K comparable, V any](size int) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		m: make(map[K]V, size),
		s: make([]K, 0, size),
	}
}

func (om *OrderedMap[K, V]) Get(k K) (V, bool) {
	v, ok := om.m[k]
	return v, ok
}

// Add returns false if the key exists already,
// without adding the value.
func (om *OrderedMap[K, V]) Add(k K, v V) bool {
	if _, ok := om.m[k]; ok {
		return false
	}
	om.m[k] = v
	om.s = append(om.s, k)
	return true
}

func (om *OrderedMap[K, V]) Exists(k K) bool {
	_, ok := om.m[k]
	return ok
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.s)
}

// All returns an iterator that can be used to iterate
// like over normal maps (with a for range loop).
func (om *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(key K, value V) bool) {
		for _, k := range om.s {
			if !yield(k, om.m[k]) {
				return
			}
		}
	}
}

// Build adds key k and value v to the ordered map and returns the map itself.
// This allows elegant building of maps.
// But no feedback is given when the addition wasn't successful.
// So this shouldn't be used in production but rather for building test data.
func (om *OrderedMap[K, V]) Build(k K, v V) *OrderedMap[K, V] {
	om.Add(k, v)
	return om
}
