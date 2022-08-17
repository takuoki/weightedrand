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
