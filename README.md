# weightedrand

A golang package for weighted random.

## Usage

```go
	items := []weightedrand.WeightedItem[string]{
		{Weight: 5, Item: "Hi!"},
		{Weight: 3, Item: "Hello!"},
		{Weight: 2, Item: "What's up?"},
	}

	wr := weightedrand.New(time.Now().UnixNano(), items)

	fmt.Println(wr.GetItem()) // "Hi!", "Hello!" or "What's up?"
```
