# 02 — Integers

Core idea from this chapter: same Red → Green → Refactor loop as `01_hello_world`,
but applied to `int` return values, doc comments, and **Testable Examples**.

## What I built

Deviated slightly from the book: instead of one `integers` package,
each function lives in its own package (one package per directory rule).

```
02_integers/
├── go.mod                    # module github.com/aman-void/tdd-go/02_integers
├── adder/
│   ├── adder.go              # package adder, func Add(x, y int) int
│   └── adder_test.go         # TestAdder + ExampleAdd
├── subtractor/
│   ├── subtractor.go         # package subtractor, func Subtract(a, b int) int
│   └── subtractor_test.go    # TestSubtract + ExampleSubtract
└── is_even/                  # my own extension beyond the book
    ├── is_even.go            # package iseven, func IsEven(num int) bool
    └── is_even_test.go       # TestIsEven + 4 Examples
```

Run everything:

```bash
go test ./... -v
```

All 3 packages pass: `TestAdder`, `TestSubtract`, `TestIsEven`, plus all `Example*`.

## TDD cycle, concretely

1. **Red — write failing test first:**

   ```go
   func TestAdder(t *testing.T) {
       sum := Add(3, 4)
       expected := 7
       if expected != sum {
           t.Errorf("expected %d but got %d", expected, sum)
       }
   }
   ```

   First failure is a _compile_ failure: `undefined: Add`. That's expected —
   it proves the test is actually wired up.

2. **Green (fake it) — minimal code to compile:**

   ```go
   func Add(x, y int) int {
       return 0
   }
   ```

   Now the test fails for the _right_ reason: `expected 7 but got 0`.

   Pedantic TDD step the book calls out: `return 4` / `return 7` would
   pass one test. That's why you don't stop at one example — the real fix:

   ```go
   // Add takes two parameters and return their sum
   func Add(x, y int) int {
       return x + y
   }
   ```

3. **Refactor — here there was almost nothing to refactor.**
   The improvement was documentation (see below), not logic.

I repeated the same loop for `Subtract(a, b int) int { return a - b }`
and `IsEven(num int) bool { return num%2 == 0 }`.

## Go notes worth remembering

### 1. One package per directory

Go enforces this. That's why `adder`, `subtractor`, `iseven` are separate
folders — you can't put `package adder` and `package subtractor` in the same dir.

### 2. Shortened parameter types

When consecutive params share a type, write it once:

```go
func Add(x, y int) int        // same as (x int, y int)
func Subtract(a, b int) int
```

### 3. Format verbs

Strings chapter used `%q` (quoted string). Integers need different verbs:

| Verb | Use                | Example in this repo                                   |
| ---- | ------------------ | ------------------------------------------------------ |
| `%d` | integers           | `t.Errorf("expected %d but got %d", expected, sum)`    |
| `%t` | booleans           | `t.Errorf("expected %t but got %t", expected, isEven)` |
| `%v` | default / anything | useful fallback                                        |

### 4. Named return values — skipped here on purpose

Chapter 01 used `func Hello(...) (greeting string)`. Here it's just `int`:

```go
func Add(x, y int) int
```

Rule of thumb from [CodeReviewComments](https://go.dev/wiki/CodeReviewComments#named-result-parameters):
use named returns only when the meaning isn't obvious from context.
`Add` returning `x + y` is obvious, so plain `int` is clearer.

### 5. Doc comments become documentation

```go
// Add takes two parameters and return their sum
func Add(x, y int) int {
```

`go doc`, editors, `pkgsite`, and `pkg.go.dev` all surface this.
Write it as a full sentence starting with the function name.

Missing here (improvement for later): `Subtract` and `IsEven` deserve the
same treatment — e.g. `// IsEven reports whether num is even.`

## Testable Examples — the big new idea

Examples live in `_test.go` files, start with `Example`, compile on every
`go test` run, and show up in docs. They prevent README-style docs from rotting.

Basic form (`adder/adder_test.go`):

```go
func ExampleAdd() {
    sum := Add(3, 4)
    fmt.Println(sum)
    // Output: 7
}
```

Key rules:

- Must `import "fmt"` (rely on `goimports` / editor auto-import).
- The trailing `// Output: 7` comment is what turns it from
  "compile-only" into "compile + run + assert".
  Remove it → `go test -v` won't execute the Example anymore.
- `go test -v` shows them explicitly:
  `=== RUN ExampleAdd --- PASS: ExampleAdd`.

Multiple examples for one function need a suffix (`is_even/is_even_test.go`):

```go
func ExampleIsEven_odd()     // IsEven(7)  → false
func ExampleIsEven_even()    // IsEven(8)  → true
func ExampleIsEven_zero()    // IsEven(0)  → true (0 % 2 == 0)
func ExampleIsEven_negative() // IsEven(-2) → true (Go % keeps sign of dividend, -2%2==0)
```

This is the idiomatic way to document edge cases: zero and negatives,
not just the happy path. Note the `if/else` in `IsEven` can simplify to
`return num%2 == 0` — good refactor candidate.

View them rendered with:

```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .
# → your package → func Add → Example
```

Or publish and read on `pkg.go.dev`.

## Takeaways

- TDD with ints is identical to TDD with strings — only the verbs (`%d`, `%t`) change.
- A compile error (`undefined: Add`) counts as Red. Then fake Green (`return 0`),
  then real Green (`return x + y`).
- One hard-coded return value passing is a TDD smell; more examples (or
  property-based testing, previewed in the book) force the general solution.
- `ExampleX` + `// Output:` gives you tested documentation for free.
  Use `_suffix` for multiple examples per function.
- Keep doc comments on every exported function; keep packages tiny and focused.
