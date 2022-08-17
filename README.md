# weightedrand

[![GoDoc](https://godoc.org/github.com/takuoki/weightedrand?status.svg)](https://godoc.org/github.com/takuoki/weightedrand)
![CI](https://github.com/takuoki/weightedrand/actions/workflows/auto-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/takuoki/weightedrand/branch/main/graph/badge.svg?token=YJGUO2OZDC)](https://codecov.io/gh/takuoki/weightedrand)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

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

| Weight | Item         |            Expected Rate |
| -----: | :----------- | -----------------------: |
|      5 | "Hi!"        | 50 % (= 5 / (5 + 3 + 2)) |
|      3 | "Hello!"     | 30 % (= 3 / (5 + 3 + 2)) |
|      2 | "What's up?" | 20 % (= 2 / (5 + 3 + 2)) |
