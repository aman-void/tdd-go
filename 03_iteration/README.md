# 03 — Iteration

Core idea from this chapter: same Red → Green → Refactor loop as `01_hello_world` and `02_integers`,
but applied to repeated work with `for` — the *only* looping construct in Go — plus **Benchmarks** and **Testable Examples**.

Based on [Learn Go with Tests — Iteration](https://github.com/quii/learn-go-with-tests/tree/main/for).

## What I built

Single `iteration` package (one package per directory rule, same as before):

```
03_iteration/
├── go.mod                         # module github.com/aman-void/tdd-go/03_iteration
├── iteration.go                   # package iteration, func Repeat(character string) string
├── iteration_test.go              # TestRepeat + ExampleRepeat (my example)
├── iteration_benchmark_test.go    # BenchmarkRepeat + 4 extra benchmarks (my examples)
└── README.md                      # this file
```

Run everything:

```bash
go test ./... -v
go test -bench=. -benchmem
```

All pass on my machine:

```
=== RUN   TestRepeat
--- PASS: TestRepeat
=== RUN   ExampleRepeat
--- PASS: ExampleRepeat
```

```text
BenchmarkRepeat-4                 25870501    51.24 ns/op    8 B/op    1 allocs/op
BenchmarkStringConcatenation-4     2374438   509.8  ns/op  128 B/op    9 allocs/op
BenchmarkStringBuilder-4          10853826   114.0  ns/op   24 B/op    2 allocs/op
BenchmarkJSONEncoding-4            2506603   471.4  ns/op   48 B/op    2 allocs/op
BenchmarkMapLookup-4              98561480    12.42 ns/op    0 B/op    0 allocs/op
```

## TDD cycle, concretely

1. **Red — write failing test first** (`iteration_test.go` — my example):

   ```go
   func TestRepeat(t *testing.T) {
       repeated := Repeat("a")
       expected := "aaaaa"

       if expected != repeated {
           t.Errorf("expected %q but got %q", expected, repeated)
       }
   }
   ```

   First failure is a *compile* failure: `undefined: Repeat`. That's expected —
   it proves the test is actually wired up.

2. **Green (fake it) — minimal code to compile** (`iteration.go`):

   ```go
   package iteration

   func Repeat(character string) string {
       return ""
   }
   ```

   Now the test fails for the *right* reason: `expected 'aaaaa' but got ''`.

   Then the real fix — `for` follows most C-like languages, except no
   parentheses and braces `{ }` are always required:

   ```go
   func Repeat(character string) string {
       var repeated string
       for i := 0; i < 5; i++ {
           repeated = repeated + character
       }
       return repeated
   }
   ```

3. **Refactor — introduce `const` + `+=`**:

   ```go
   const repeatCount = 5

   func Repeat(character string) string {
       var repeated string
       for i := 0; i < repeatCount; i++ {
           repeated += character
       }
       return repeated
   }
   ```

   `+=` is the *Add AND assignment operator*: adds the right operand to the
   left and assigns the result back. Works for strings and integers.

## Go notes worth remembering

### 1. `for` is the only loop

No `while`, `do`, `until` keywords in Go — only `for`:

```go
for i := 0; i < repeatCount; i++ {
    repeated += character
}
```

Other variants from [Go by Example](https://gobyexample.com/for):

```go
for i < 10 { }        // while-style
for { }               // infinite
for i, v := range s { } // over slice/map/string
```

### 2. `var x string` vs `x := ...`

So far we used `:=` (declare + initialize shorthand). Here we declare only:

```go
var repeated string // zero value "" — then fill it in the loop
```

`:=` is shorthand for both steps. Explicit `var` makes the
"start empty, accumulate" intent clear.

### 3. `const repeatCount = 5`

Magic numbers scattered in loops rot. Naming it:

- documents intent,
- gives one place to change repetition count,
- matches the practice exercise: let the caller pass the count in.

### 4. Strings are immutable — why `+=` in a loop can hurt

Every `repeated += character` can allocate a new backing array and copy.
For 5 iterations it's insignificant; for large loops it gets expensive.

Standard library fix: `strings.Builder` — keeps an internal buffer,
`WriteString` appends without repeated copying, `String()` returns the result:

```go
import "strings"

const repeatCount = 5

func Repeat(character string) string {
    var repeated strings.Builder
    for i := 0; i < repeatCount; i++ {
        repeated.WriteString(character)
    }
    return repeated.String()
}
```

### 5. My final version — `strings.Repeat` (my example)

Since the operation is literally "repeat this string N times", the stdlib
already has the idiomatic solution. No manual loop / Builder needed
(`iteration.go` in this repo):

```go
package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
    // A simple approach would be to concatenate strings using +=:
    //
    // var repeated string
    // for i := 0; i < repeatCount; i++ {
    //     repeated += character
    // }
    //
    // However, strings in Go are immutable. Repeatedly using += in a loop
    // can cause multiple allocations and copies as the string grows.
    //
    // For a small number of iterations this is usually insignificant,
    // but for larger loops it can become unnecessarily expensive.

    // strings.Builder is designed for efficiently constructing strings:
    //
    // var repeated strings.Builder
    // for i := 0; i < repeatCount; i++ {
    //     repeated.WriteString(character)
    // }
    //
    // return repeated.String()
    //
    // The Builder maintains an internal buffer, avoiding the repeated
    // allocation and copying that can happen with string concatenation.

    // Since our actual operation is simply "repeat this string N times",
    // the standard library already provides the most idiomatic solution.
    //
    // strings.Repeat handles the allocation and construction internally,
    // so there is no need to manually manage a Builder or a loop here.
    return strings.Repeat(character, repeatCount)
}
```

The comments preserve the whole journey: `+=` → `Builder` → `strings.Repeat`.
That history is the note itself.

## Testable Examples — my example

`iteration_test.go` in this repo (beyond the book's test):

```go
package iteration

import (
    "fmt"
    "testing"
)

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
```

Key rules (same as `02_integers`):

- File ends in `_test.go`, function starts with `Example`.
- Must `import "fmt"`.
- Trailing `// Output: *****` turns it from "compile-only" into
  "compile + run + assert". Remove it → `go test -v` won't execute it.
- `go test -v` shows it: `=== RUN ExampleRepeat --- PASS: ExampleRepeat`.
- Appears in `go doc` / `pkgsite` / `pkg.go.dev` as tested documentation.

Practice extension from the book I should still do: `Write ExampleRepeat`
is done; remaining are *caller-specified count* (`Repeat(s string, n int)`)
and exploring `strings` package with tests.

## Benchmarking — my examples

Benchmarks are first-class in Go, structured like tests but with `*testing.B`
(`iteration_benchmark_test.go` in this repo):

```go
func BenchmarkRepeat(b *testing.B) {
    for b.Loop() {
        Repeat("*")
    }
}
```

- `b.Loop()` returns true while the benchmark should keep running.
  Only the loop body is timed; setup/cleanup outside is excluded.
- Framework picks iteration count (`b.N`) to get stable numbers.
- Run with `go test -bench=.` (Powershell: `go test -bench="."`).
- Add `-benchmem` for allocation stats: `B/op` (bytes per op),
  `allocs/op` (allocations per op).

### My extra benchmarks beyond the book

I added 4 more to build intuition for what is cheap vs expensive:

```go
func BenchmarkStringConcatenation(b *testing.B) {
    for b.Loop() {
        var result string

        for i := 0; i < 10; i++ {
            result += "[]"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for b.Loop() {
        var builder strings.Builder

        for i := 0; i < 10; i++ {
            builder.WriteString("&")
        }
        // builder.String()
    }
}

type MyStruct struct {
    Name string
    Age  int
}

func BenchmarkJSONEncoding(b *testing.B) {
    data := MyStruct{
        Name: "Void",
        Age:  20,
    }

    for b.Loop() {
        json.Marshal(data)
    }
}

func BenchmarkMapLookup(b *testing.B) {
    data := map[string]int{
        "foo": 42,
    }

    for b.Loop() {
        _ = data["foo"]
    }
}
```

What they teach (`go test -bench=. -benchmem`):

| Benchmark | ns/op | B/op | allocs/op | Lesson |
| --------- | ----- | ---- | --------- | ------ |
| `BenchmarkRepeat` (`strings.Repeat`) | ~51 | 8 | 1 | 1 alloc total — stdlib pre-sizes correctly |
| `BenchmarkStringConcatenation` (`+=` ×10) | ~510 | 128 | 9 | ~10× slower, 9 allocs — immutable-string copying |
| `BenchmarkStringBuilder` (`WriteString` ×10) | ~114 | 24 | 2 | ~4.5× faster than `+=`, far fewer allocs |
| `BenchmarkJSONEncoding` | ~471 | 48 | 2 | reflection + encoding dominates; avoid in hot loops |
| `BenchmarkMapLookup` | ~12 | 0 | 0 | map read is ~40× cheaper than JSON encode here |

Note: `BenchmarkStringBuilder` above comments out `builder.String()` —
calling it adds one alloc + copy. Uncomment to measure the full cost.

## Takeaways

- TDD with loops is identical — Red (`undefined: Repeat`), fake Green
  (`return ""`), real Green (`for` + `+=`), Refactor (`const`, `Builder`, `strings.Repeat`).
- `for` is Go's only loop; no parens, braces always required.
- `var s string` declares the zero value `""`; `:=` is declare + init shorthand.
- `const repeatCount` removes the magic number; next step is making it a parameter.
- Strings are immutable: `+=` in a loop copies; `strings.Builder` buffers;
  `strings.Repeat` is idiomatic when you just need repetition.
- `ExampleRepeat` + `// Output: *****` gives tested docs for free.
- `BenchmarkX(b *testing.B)` + `for b.Loop()` + `go test -bench=. -benchmem`
  quantifies `ns/op`, `B/op`, `allocs/op` — my 4 extra benchmarks make the
  cost of `+=` vs `Builder` vs `encoding/json` vs map lookup concrete.

## Commands Cheat Sheet

```bash
# Run tests in this package
go test -v

# Run all tests from repo root
go test ./... -v

# Run benchmarks (current package)
go test -bench=.
go test -bench=. -benchmem

# Run one benchmark
go test -bench=Repeat -benchmem -run=^$

# Format
gofmt -w .
```

## Reference

- [Learn Go with Tests — Iteration](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration)
- Code in this repo: `iteration.go`, `iteration_test.go`, `iteration_benchmark_test.go`
- Stdlib: [`strings.Builder`](https://pkg.go.dev/strings#Builder), [`strings.Repeat`](https://pkg.go.dev/strings#Repeat), [`strings` package](https://golang.org/pkg/strings)
