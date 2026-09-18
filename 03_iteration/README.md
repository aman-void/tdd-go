# 03 — Iteration: Loops, Table Tests and Benchmarks

Now things get fun. We repeat work with `for` — the *only* loop in Go —
then level up our testing with table-driven tests and benchmarks.

If Chapter 01 was "how to test" and 02 was "how to document",
this one is "how to test smarter and measure faster".

> Based on [Learn Go with Tests — Iteration](https://github.com/quii/learn-go-with-tests/tree/main/for)

## What you'll learn here

- `for` is all you get (no `while`)
- `const`, `var` vs `:=`, why `+=` in loops can hurt
- `strings.Builder` and `strings.Repeat`
- Table-driven tests with `t.Run`
- Benchmarks with `b.Loop()` and `-benchmem`

## Project tour

```
03_iteration/
├── go.mod
├── iteration.go                # Repeat() + CompareStrings()
├── iteration_test.go           # TestRepeat + ExampleRepeat + TestCompareStrings
├── iteration_benchmark_test.go # BenchmarkRepeat + 4 intuition builders + BenchmarkCompareStrings
└── README.md
```

Core code:

```go
// iteration.go
package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
    return strings.Repeat(character, repeatCount)
}

func CompareStrings(str1, str2 string) int {
    return strings.Compare(str1, str2)
}
```

> Inside the real file I kept the journey in comments: `+=` loop ->
> `strings.Builder` -> `strings.Repeat`. Open `iteration.go` to see why each step mattered.

Tests — classic + table:

```go
func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"
    if expected != repeated {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func ExampleRepeat() {
    repeated := Repeat("*")
    fmt.Println(repeated)
    // Output: *****
}

// Table-driven: one logic, many rows
func TestCompareStrings(t *testing.T) {
    tests := []struct {
        name     string
        str1     string
        str2     string
        expected int
    }{
        {name: "equal strings", str1: "hi", str2: "hi", expected: 0},
        {name: "first string is smaller", str1: "apple", str2: "watermelon", expected: -1},
        {name: "first string is greater", str1: "california", str2: "boston", expected: 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := CompareStrings(tt.str1, tt.str2)
            if result != tt.expected {
                t.Errorf("CompareStrings(%q, %q) = %d and expected %d",
                    tt.str1, tt.str2, result, tt.expected)
            }
        })
    }
}
```

What you'll see:

```
=== RUN   TestRepeat
=== RUN   TestCompareStrings/equal_strings
=== RUN   TestCompareStrings/first_string_is_smaller
=== RUN   TestCompareStrings/first_string_is_greater
=== RUN   ExampleRepeat
--- PASS
```

## How we built it

`Repeat`: Red `undefined: Repeat` -> fake `return ""` -> real `for` + `+=` ->
refactor with `const`, then `Builder`, then `strings.Repeat` (the idiomatic finish).

`CompareStrings`: built to teach table tests. `strings.Compare` returns
`-1` if first < second, `0` if equal, `+1` if greater — perfect for a 3-row table.

## Go bits worth remembering

- No `while` in Go. Just `for`: `for i:=0;i<5;i++{}`, `for cond{}`, `for{}`,
  `for i,v := range s{}`. Braces always, parens never.
- `var repeated string` starts as `""`, then you fill it. `:=` is declare+init shorthand.
- `const repeatCount = 5` kills magic numbers. Next natural step:
  let the caller decide — `Repeat(s string, n int)`.
- Strings are immutable. `+=` in a big loop copies a lot. `strings.Builder`
  buffers instead, and `strings.Repeat` is best when you literally just repeat.

## Testing bits: table tests in plain English

Think of it like a spreadsheet:

1. `tests := []struct{...}` — your rows.
2. `for _, tt := range tests` — `tt` is convention for "this row".
3. `t.Run(tt.name, ...)` — each row becomes a named subtest, so failures tell you
   *which* row broke, not just "something broke".

Adding a new case? Just add one line to the table. No new logic needed.

## Benchmarks — let's measure, not guess

Benchmarks look like tests but take `*testing.B`:

```go
func BenchmarkRepeat(b *testing.B) {
    for b.Loop() {
        Repeat("*")
    }
}

func BenchmarkCompareStrings(b *testing.B) {
    for b.Loop() {
        CompareStrings("Hi", "Hi")
        CompareStrings("peach", "watermelon")
        CompareStrings("pomogranate", "grapes")
    }
}
```

Plus 4 extra I added to build intuition: `+=` vs `Builder` vs `json.Marshal` vs map lookup.

Run them:

```bash
go test -bench=. -benchmem
go test -bench=Repeat -benchmem -run=^$  # just one
```

How to read the output:

| Benchmark | What it tells you |
|---|---|
| `BenchmarkRepeat ~51 ns/op, 1 alloc` | `strings.Repeat` pre-sizes — 1 alloc total, fast |
| `StringConcatenation ~510 ns/op, 9 allocs` | `+=` x10 copies every time, ~10x slower |
| `StringBuilder ~114 ns/op, 2 allocs` | Buffering wins, ~4.5x faster than `+=` |
| `JSONEncoding ~471 ns/op` | Reflection + encoding is heavy, avoid in hot loops |
| `MapLookup ~12 ns/op, 0 allocs` | Map read is ~40x cheaper than JSON here |

> Tip: `b.Loop()` keeps the hot part timed and leaves setup outside.
> `-benchmem` adds `B/op` and `allocs/op` so you see memory, not just speed.

## Run it yourself

```bash
go test -v
go test -v -run TestCompareStrings
go test -bench=. -benchmem
gofmt -w .
```

## Takeaway in one line

> `for` does all looping. Tables make many cases cheap to test,
> and benchmarks turn "I think it's faster" into numbers you can trust.
>
> Reference: Iteration book page, `strings.Builder`, `strings.Repeat`, code in this folder.
