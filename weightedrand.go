package weightedrand

import (
	"math/rand"
)

type WeightedItem[T any] struct {
	Weight int
	Item   T
}

type WeightedRand[T any] struct {
	rand    *rand.Rand
	items   []WeightedItem[T]
	sum     int
	borders []int
}

func New[T any](seed int64, items []WeightedItem[T]) *WeightedRand[T] {
	sum := 0
	borders := []int{}
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

func (r *WeightedRand[T]) GetItem() T {
	num := r.rand.Intn(r.sum)
	for i, b := range r.borders {
		if num < b {
			return r.items[i].Item
		}
	}
	panic("sorry, something wrong...")
}
