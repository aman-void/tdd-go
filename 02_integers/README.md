# 02 — Integers: Docs, Verbs and Testable Examples

Okay, strings are done. Now let's play with numbers.

Same TDD rhythm as Chapter 01, but now functions return `int` and `bool`.
You'll also meet two Go favorites: doc comments and `Example` tests —
documentation that actually runs, so it can't go stale.

## What you'll learn here

- TDD with `int` / `bool` (verbs `%d`, `%t`)
- One package per folder rule in Go
- Doc comments that show up in `go doc`
- `Example` tests with `// Output:`
- Edge cases: zero, negatives

## Project tour

Go doesn't allow two packages in one folder, so each tiny function gets its own home:

```
02_integers/
├── go.mod
├── adder/adder.go + adder_test.go           # Add(x,y int) int
├── subtractor/subtractor.go + _test.go      # Subtract(a,b int) int
└── is_even/is_even.go + _test.go            # IsEven(num int) bool, my extension
```

The code — small on purpose:

```go
// adder/adder.go
package adder

// Add takes two parameters and return their sum
func Add(x, y int) int {
    return x + y
}
```

```go
// subtractor/subtractor.go
package subtractor

// Subtract function takes two parameters and return the subtract
func Subtract(a, b int) int {
    return a - b
}
```

```go
// is_even/is_even.go
package iseven

func IsEven(num int) bool {
    if num%2 == 0 {
        return true
    } else {
        return false
    }
}
```

Tests + living docs:

```go
// adder/adder_test.go
func TestAdder(t *testing.T) {
    sum := Add(3, 4)
    expected := 7
    if expected != sum {
        t.Errorf("expected %d but got %d", expected, sum)
    }
}

func ExampleAdd() {
    sum := Add(3, 4)
    fmt.Println(sum)
    // Output: 7
}
```

`is_even` goes further — 4 examples for one function:

```go
func ExampleIsEven_odd() {
    fmt.Println(IsEven(7))
    // Output: false
}

func ExampleIsEven_even() {
    fmt.Println(IsEven(8))
    // Output: true
}

func ExampleIsEven_zero() {
    fmt.Println(IsEven(0))
    // Output: true
}

func ExampleIsEven_negative() {
    fmt.Println(IsEven(-2))
    // Output: true
}
```

## How we built it

Same story: `undefined: Add` (Red) -> `return 0` (fake Green, fails with `expected 7 got 0`) ->
`return x + y` (real Green). Repeated for `Subtract` and `IsEven`.

Friendly warning from the book: if you return a hard-coded `7`, one test passes.
That's why you need more examples. TDD only forces general code if your tests demand it.

## Go bits worth remembering

- `func Add(x, y int)` is shorthand for `(x int, y int)`. Use it when types match.
- Format verbs cheat: `%d` ints, `%t` bools, `%q` strings, `%v` anything.
- Doc comment must start with the name: `// Add takes...`. It powers `go doc` and pkg.go.dev.
- `IsEven` can slim down to `return num%2 == 0` — good first refactor to try.
- Fun facts: `0` is even, and Go's `%` keeps the dividend sign so `-2 % 2 == 0`.

## Testing bits worth remembering

- `Example` must be in a `_test.go` file, start with `Example`, and `import "fmt"`.
- The magic line is `// Output: 7`. With it, `go test` runs and checks it.
  Without it, it only checks that it compiles.
- Suffix pattern `ExampleIsEven_odd` lets you document happy path + edges in one place.
  Future-you will thank you for that zero/negative coverage.

## Run it yourself

```bash
go test ./... -v
# you'll see TestAdder, TestSubtract, TestIsEven + all Example* passing
```

No benchmarks here yet — Chapter 03 is where we measure speed.

## Takeaway in one line

> Same TDD loop, new types. Write doc comments, cover edges with `Example_suffix`,
> and let `// Output:` keep your docs honest.
