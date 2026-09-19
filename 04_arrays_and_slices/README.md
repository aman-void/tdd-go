# 04 — Arrays and Slices: Sums, Tails and Real Coverage

Hey, welcome back. Chapters 01-03 were warm-up: strings, ints, loops.

This chapter is where Go gets opinionated. Arrays have fixed size,
slices have flexible size — and almost always you want slices.
We prove it by building `Sum`, `SumAll`, and `SumAllTails` test-first.

> Based on [Learn Go with Tests — Arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
> plus Rob Pike's [The cover story](https://go.dev/blog/cover) for the coverage half.

## What you'll learn here

- Arrays `[5]int` vs slices `[]int` — why `Sum(numbers []int)` wins
- `for _, num := range numbers` for safe iteration
- Variadics `...[]int` for `SumAll` / `SumAllTails`
- `append` vs `make`, slicing `numbers[1:]`, guarding `len == 0`
- `slices.Equal` because slices can't use `!=` (only vs `nil`)
- Table tests, `testing.TB` helpers, more `Example_suffix` docs
- Benchmarks with `b.Loop()` + `-benchmem`
- Coverage the Go way: `go test -cover -coverprofile + go tool cover`

## Project tour

```
04_arrays_and_slices/
├── go.mod                  # module github.com/aman-void/04_arrays_and_slices
├── sum.go                  # Sum(), SumAll(), SumAllTails()
├── sum_test.go             # TestSum + TestSumAll + TestSumAllTails + 4 Examples
├── sum_benchmark_test.go   # BenchmarkSum + Large + SumAll + SumAllTails
└── README.md
```

`sum.go` — the whole library:

```go
package arraysandslices

// If we pass `numbers [5]int` to the Sum func parameter
// it works for array but not for slice. So we are using slice here.

// Sum returns the total of a slice of ints.
// An empty or nil slice sums to 0.
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// SumAll returns the sum of each slice passed as a variadic argument.
// For example, SumAll([]int{1, 2}, []int{0, 9}) returns []int{3, 9}.
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	var sums []int
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		// more efficient using append

		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns the sum of each slice's tail (all elements except the first).
// Empty slices and single-element slices contribute 0, so no panic on numbers[1:].
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
```

> The commented-out `make` lines are kept on purpose — they show the journey
> from pre-sized `make([]int, len)` + index assignment to `var sums []int` + `append`,
> same spirit as the `+=` -> `Builder` -> `Repeat` comments in `03_iteration`.

Tests — all table-driven now:

```go
func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{name: "collection of 5 numbers", numbers: []int{1, 2, 3, 4, 5}, want: 15},
		{name: "collection of any size", numbers: []int{1, 2, 3}, want: 6},
		{name: "empty slice", numbers: []int{}, want: 0},
		{name: "nil slice", numbers: nil, want: 0},
		{name: "single element", numbers: []int{7}, want: 7},
		{name: "negatives cancel out", numbers: []int{-1, -2, 3}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.numbers)
			if got != tt.want {
				t.Errorf("got %d want %d given %v", got, tt.want, tt.numbers)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{name: "two slices", input: [][]int{{1, 2}, {0, 9}}, want: []int{3, 9}},
		{name: "three slices of different sizes", input: [][]int{{1, 2, 3}, {4}, {}}, want: []int{6, 4, 0}},
		{name: "no slices", input: nil, want: nil},
		{name: "nil and empty slices", input: [][]int{nil, {}}, want: []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumAll(tt.input...)
			checkSums(t, got, tt.want)
		})
	}
}

func TestSumAllTails(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{name: "make the sums of tail of", input: [][]int{{1, 2}, {3, 4}}, want: []int{2, 4}},
		{name: "safely sum empty slices", input: [][]int{{}, {3, 4, 5}}, want: []int{0, 9}},
		{name: "single element tails are empty", input: [][]int{{1}, {3, 4, 5}}, want: []int{0, 9}},
		{name: "nil slice is safe", input: [][]int{nil, {1, 2, 3}}, want: []int{0, 5}},
		{name: "no slices", input: nil, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumAllTails(tt.input...)
			checkSums(t, got, tt.want)
		})
	}
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

Living docs — 4 examples:

```go
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Println(result)
	// Output: 15
}

func ExampleSum_empty() {
	fmt.Println(Sum([]int{}))
	fmt.Println(Sum(nil))
	// Output:
	// 0
	// 0
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2, 3, 4, 5}, []int{3, 4})
	fmt.Println(got)
	// Output: [15 7]
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
	fmt.Println(got)
	// Output: [5 9]
}
```

What you'll see:

```
=== RUN   TestSum/collection_of_5_numbers
=== RUN   TestSum/collection_of_any_size
=== RUN   TestSum/empty_slice
=== RUN   TestSum/nil_slice
=== RUN   TestSum/single_element
=== RUN   TestSum/negatives_cancel_out
=== RUN   TestSumAll/two_slices
=== RUN   TestSumAll/three_slices_of_different_sizes
=== RUN   TestSumAll/no_slices
=== RUN   TestSumAll/nil_and_empty_slices
=== RUN   TestSumAllTails/make_the_sums_of_tail_of
=== RUN   TestSumAllTails/safely_sum_empty_slices
=== RUN   TestSumAllTails/single_element_tails_are_empty
=== RUN   TestSumAllTails/nil_slice_is_safe
=== RUN   TestSumAllTails/no_slices
=== RUN   ExampleSum / ExampleSum_empty / ExampleSumAll / ExampleSumAllTails
--- PASS
```

## How we built it (the TDD story)

1. **Sum:** Red `undefined: Sum` -> fake `return 0` -> real `for _, num := range`.
   First version took `[5]int` (array). That only works for size 5.
   Switching the signature to `[]int` (slice) made any size pass — restored
   the old `collection of 5 numbers` case as a table row so both stay green.
2. **SumAll:** Red with variadic `...[]int`. Green via `make([]int, len)` + index.
   Refactor to `var sums []int` + `append(Sum(...))` — simpler, no index bookkeeping.
   Old comments kept in file to show the trade.
3. **SumAllTails:** Red asked for `numbers[1:]`. First green panicked on `[]int{}`:
   `slice bounds out of range`. Fix was the `if len(numbers) == 0 { append 0 }` guard.
   That same guard also covers `nil` and makes single-element tails (`[1] -> [] -> 0`) safe.

## Bugs fixed in this pass

- **Swapped `t.Errorf` args in `TestSum`:** was `t.Errorf("got %d want %d given %v", want, got, numbers)` —
  message printed backwards. Fixed to `(got, want, numbers)`. The commented-out
  block had the same swap.
- **Helper took `*testing.T`, now `testing.TB`:** `func checkSums(t testing.TB, ...)` +
  `t.Helper()` so it works in tests, examples, and benchmarks, and failures point
  at the caller. This matches Chapters 01 and 03.
- **Thin coverage dressed as 100%:** baseline was already `100.0% of statements`
  but with only 1-2 cases per function. Added nil/empty/single/negative/no-arg rows
  so the 100% actually means something. See coverage section below.

## Go bits worth remembering

- `[5]int` is an array (value, fixed size, part of the type). `[]int` is a slice
  (header: pointer + len + cap over a backing array). `Sum([]int)` accepts any length;
  `Sum([5]int)` would not accept `[3]int`.
- `range` gives index + copy: `for _, num := range numbers`. Blank `_` when you don't need the index.
- Variadic `...[]int` collects args into `[][]int` inside. Call with `SumAll(a, b)` or spread with `SumAll(slices...)` — tests use `tt.input...`.
- `tail := numbers[1:]` shares the backing array, no copy. `numbers[1:]` on empty panics,
  hence the `len == 0` guard. `Sum(nil)` is safe: ranging over `nil` runs zero times.
- `var sums []int` starts as `nil`. `append(nil, x)` allocates as needed.
  `slices.Equal(nil, nil)` is `true`, so `no slices -> nil` passes cleanly.
- Format verbs: `%d` ints, `%v` slices, `%q` strings. Keep `got` first, `want` second.

## Testing bits worth remembering

- Slices can't use `!=` (compile error except vs `nil`). Use `slices.Equal(got, want)` from stdlib.
- Table pattern again: `tests := []struct{name, input, want}` + `for _, tt := range tests` + `t.Run(tt.name, ...)`.
  New edge = one new row, no new logic.
- Shared helper `checkSums(t testing.TB, ...)` removes duplication between `TestSumAll` and `TestSumAllTails`.
- `ExampleSum_empty` / `ExampleSumAllTails` use the suffix pattern from Chapter 02
  (`ExampleIsEven_odd`). With `// Output:` they are checked tests; without it compile-only.

## Benchmarks — let's measure, not guess

```go
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumLarge(b *testing.B) {
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i + 1
	}
	b.ResetTimer()
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{3, 4}
	for b.Loop() {
		SumAll(a, c)
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{3, 4}
	for b.Loop() {
		SumAllTails(a, c)
	}
}
```

Run them:

```bash
go test -bench=. -benchmem
go test -bench=SumLarge -benchmem -run=^$  # just one
```

How to read the output (example output — your numbers will vary by machine,
focus on relative order and `allocs/op` rather than exact `ns/op`):

| Benchmark                                              | What it tells you                                                           |
| ------------------------------------------------------ | --------------------------------------------------------------------------- |
| `BenchmarkSum ~2.4 ns/op, 0 B/op, 0 allocs/op`         | Tiny 5-elem loop, no heap — pure stack/range                                |
| `BenchmarkSumLarge ~4984 ns/op, 0 B/op, 0 allocs/op`   | 10k elems ~2000x work of 5 elems, still 0 allocs — `range` doesn't allocate |
| `BenchmarkSumAll ~85 ns/op, 24 B/op, 2 allocs/op`      | `append` from `nil` grows the result slice — 2 allocs for 2 sums            |
| `BenchmarkSumAllTails ~83 ns/op, 24 B/op, 2 allocs/op` | Same shape as `SumAll`; `numbers[1:]` is free (no copy)                     |

> Tip from Chapter 03 still holds: `b.Loop()` keeps setup outside the timer
> (plus explicit `b.ResetTimer()` for the 10k fill), `-benchmem` adds `B/op` and `allocs/op`.

## Coverage notes — Rob Pike's "The cover story"

Go 1.2 added coverage with an unusual trick. Instead of `gcov`-style binary
breakpoints (hard to port, arch/OS-specific, debug-info fragile), Go **rewrites
source before compile**: each basic block gets `GoCover.Count[n] = 1`, then it
compiles, runs tests, and counts which counters fired. One `MOV`, ~3% overhead —
cheap enough to run daily.

Commands used in this chapter:

```bash
go test -cover
# PASS coverage: 100.0% of statements

go test -coverprofile=coverage.out  # implies -cover, saves profile
go tool cover -func=coverage.out
# sum.go: Sum           100.0%
# sum.go: SumAll        100.0%
# sum.go: SumAllTails   100.0%
# total: (statements)   100.0%

go tool cover -html=coverage.out    # opens browser to visually watch coverage:
                                    # green=covered, red=uncovered, grey=uninstrumented
```

This last command writes a temporary HTML page from `coverage.out` and opens it
in your browser, so you can _watch_ each line of `sum.go` light up.
Click through `Sum` / `SumAll` / `SumAllTails` — anything red is a branch your
tables never exercised. `coverage.out` itself is a generated artifact, don't
commit it; delete it after with `rm coverage.out`.

Modes via `-covermode`: `set` (did it run? default), `count` (how many times? heat map),
`atomic` (precise counts for parallel code, uses `sync/atomic`, expensive).

Why this matters here: baseline was already 100% with 4 cases, which proves
Pike's warning — coverage is statement-based, per basic block (brace-bounded),
not per value. `f() && g()` counts together; a single `{1,2,3} -> 6` test lights
every block in `Sum` green without testing `nil`, `empty`, or negatives.
The extra table rows in this pass don't raise the number (still 100%) — they raise
confidence. If someone removes the `len == 0` guard, `single element` / `nil` rows
go red in `-html` immediately, which is exactly what the tool is for.

> Reference: Pike's `Size(a int)` example scored 42.9% because `zero/big/huge/enormous`
> branches were never called — HTML showed the red branches at a glance.
> Our `SumAllTails` guard is the same lesson in miniature.

## Run it yourself

```bash
go test -v                                   # tests + examples
go test -v -run TestSumAllTails              # single test
go test -cover && go test -coverprofile=coverage.out && go tool cover -func=coverage.out
go tool cover -html=coverage.out    # watch it in the browser, then rm coverage.out
go test -bench=. -benchmem
go test -bench=SumLarge -benchmem -run=^$    # just one
gofmt -w .
go vet ./...
```

## Takeaway in one line

> Slices beat arrays for flexible code, `append` + `slices.Equal` + `TB` helpers keep
> variadic sums clean to test, and 100% coverage only counts when your tables
> force every branch — including empty, nil, and single-element tails — to run.
>
> Reference: Arrays/slices book page, Rob Pike cover blog, `slices` package, code in this folder.
