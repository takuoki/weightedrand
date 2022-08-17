package weightedrand_test

import (
	"fmt"

	"github.com/takuoki/weightedrand"
)

func Example() {

	items := []weightedrand.WeightedItem[string]{
		{Weight: 5, Item: "Hi!"},
		{Weight: 3, Item: "Hello!"},
		{Weight: 2, Item: "What's up?"},
	}

	// please set the appropriate seed like `time.Now().UnixNano()`.
	wr := weightedrand.New(1, items)

	fmt.Println(wr.GetItem())
	fmt.Println(wr.GetItem())
	fmt.Println(wr.GetItem())
	fmt.Println(wr.GetItem())
	fmt.Println(wr.GetItem())

	// Output:
	// Hi!
	// Hello!
	// Hello!
	// What's up?
	// Hi!
}

func Example_checkRate() {

	const repeat = 100000
	rate := func(a, b int) float32 {
		return float32(a) / float32(b) * 100
	}

	items := []weightedrand.WeightedItem[int]{
		{Weight: 5, Item: 0},
		{Weight: 3, Item: 1},
		{Weight: 2, Item: 2},
	}

	wr := weightedrand.New(1, items)

	counts := make([]int, len(items))
	for i := 0; i < repeat; i++ {
		counts[wr.GetItem()]++
	}

	for i := 0; i < len(items); i++ {
		fmt.Printf("- %d: weight=%d, count=%d, rate=%.2f%%\n", i, items[i].Weight, counts[i], rate(counts[i], repeat))
	}

	// Output:
	// - 0: weight=5, count=49965, rate=49.97%
	// - 1: weight=3, count=30295, rate=30.30%
	// - 2: weight=2, count=19740, rate=19.74%
}
