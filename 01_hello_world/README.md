# 01 — Hello, World: Learn TDD the Friendly Way

Hey, welcome to the start of the journey.

This chapter isn't really about printing "Hello, World".
It's about learning the rhythm you'll use for the whole repo:

**Red -> Green -> Refactor.**

We build a tiny `Hello()` function that greets people in 3 languages.
Small enough to finish in one sitting, big enough to meet packages,
`if`, `switch`, constants, and real tests.

> Based on [Learn Go with Tests — Hello World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

## What you'll learn here

- How TDD actually feels in real life
- Go modules, packages, functions
- `if`, `switch`, constants, named returns
- Exported vs unexported (capital vs small letter)
- Subtests with `t.Run`, helpers with `t.Helper`

## Project tour

```
01_hello_world/
├── go.mod        # module github.com/aman-void/tdd-go/01_hello_world
├── hello.go      # Hello(), greetingPrefix(), main()
└── hello_test.go # TestHello + helper
```

`hello.go` — the whole app:

```go
package main

import "fmt"

const (
    spanish = "Spanish"
    french  = "French"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }
    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("", ""))
    fmt.Println(Hello("Rob Pike", spanish))
    fmt.Println(Hello("Robert Griesemer", french))
    fmt.Println(Hello("Ken Thompson", ""))
}
```

`hello_test.go` — 4 behaviors, one test function:

```go
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Ken Thompson!", "")
        want := "Hello, Ken Thompson!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying 'Hello, World' when empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Rob Pike", "Spanish")
        want := "Hola, Rob Pike"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Robert Griesemer", "French")
        want := "Bonjour, Robert Griesemer"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## How we built it (the TDD story)

Think of TDD like this: test first, then just enough code, then tidy up.

1. **Red:** Wrote `TestHello` before `Hello` existed. Compiler shouted `undefined: Hello`.
   That's good — it means your test is actually connected.
2. **Green:** Returned `""` just to compile, then added the real `if` + `switch` logic.
3. **Refactor:** Pulled out `greetingPrefix()`, added constants, added the helper.
   Tests stayed green the whole time.

That loop is everything: write test, see it fail nicely, make it pass, clean up.

## Go bits worth remembering

- `Hello` returns a string, it doesn't print. `main()` does the printing.
  This separation is why testing is easy — no need to capture terminal output.
- `if name == "" { name = "World" }` — friendly default lives inside the function,
  callers don't have to remember it.
- `const` avoids magic strings. `case french:` is clearer than `case "French":`.
- `switch` beats long `if-else` chains here. `default:` falls back to English.
- `(prefix string)` + bare `return` is a named return. Cute for short funcs,
  but plain `return prefix` is often easier to read. Don't overuse it.
- Capital `Hello` = public to other packages. Small `greetingPrefix` = private detail.

## Testing bits worth remembering

- No framework needed, just `import "testing"`. File must end in `_test.go`.
- `t.Run("in Spanish", ...)` gives you named subtests. When one fails,
  you know exactly which language broke.
- `assertCorrectMessage` + `t.Helper()` keeps failure messages pointing
  at the right line, not inside the helper.
- `%q` shows strings with quotes: `got "Hello, Bob " want "Hello, Bob"` —
  that trailing space is suddenly obvious.

## Run it yourself

```bash
go run .
go test -v
```

No benchmarks in this chapter — we meet those properly in `03_iteration`.

## Takeaway in one line

> Tests drive the design. `Hello()` makes a value, `main()` prints it,
> and `t.Run` + helper keep the tests readable as behaviors grow.
