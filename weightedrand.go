// Package weightedrand implements weighted random selection.
// This package is intended for situations where you want to
// randomly generate bot comments, for example.
//
// Note that at the moment, the algorithm is based on a simple
// cumulative sum, so the order of the selection is `O(N)`.
// The number of choice items should not be large,
// as it is not performance optimized.
package weightedrand

import (
	"math/rand"
)

// WeightedItem is a weighted selection item. Any type can be used.
type WeightedItem[T any] struct {
	Weight int
	Item   T
}

// WeightedRand is an object that performs weighted random selection.
// It should be created using the New function.
type WeightedRand[T any] struct {
	rand    *rand.Rand
	items   []WeightedItem[T]
	sum     int
	borders []int
}

// New is a function that generates WeightedRand seeded with the given value.
func New[T any](seed int64, items []WeightedItem[T]) *WeightedRand[T] {
	sum := 0
	borders := make([]int, 0, len(items))
	for _, item := range items {
		sum += item.Weight
		borders = append(borders, sum)
	}
	return &WeightedRand[T]{
		rand:    rand.New(rand.NewSource(seed)),
		items:   items,
		sum:     sum,
		borders: borders,
	}
}

// GetItem returns one of the selection items at random.
func (r *WeightedRand[T]) GetItem() T {
	num := r.rand.Intn(r.sum)
	for i, b := range r.borders {
		if num < b {
			return r.items[i].Item
		}
	}
	panic("sorry, something wrong...")
}
