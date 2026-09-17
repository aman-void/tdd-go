# 01 — Hello World

This exercise is based on the **Hello World** chapter from _Learn Go with Tests_.

The goal is not just to print `"Hello, World"`, but to learn the fundamentals of Go through a small example and introduce **Test-Driven Development (TDD)**.

---

## What This Exercise Covers

- Go modules
- Packages
- Functions
- Parameters and return values
- Constants
- `if` statements
- `switch` statements
- Named return values
- Exported and unexported functions
- Table-like test organization with subtests
- Test helpers
- `testing.T`
- `t.Run()`
- `t.Helper()`
- Basic TDD workflow
- Separating logic from side effects

---

## Project Structure

```text
01_hello_world/
├── go.mod
├── hello.go
└── hello_test.go
```

### `hello.go`

Contains the application code:

- `Hello()` is the main function being tested.
- `greetingPrefix()` determines which greeting prefix to use.
- `main()` demonstrates the function.

### `hello_test.go`

Contains tests for the behavior of `Hello()`.

### `go.mod`

Defines the Go module:

```go
module github.com/aman-void/01_hello_world

go 1.27.1
```

The `go` directive specifies the Go language/toolchain version the module is intended to use.

---

# `Hello()` Function

```go
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
```

The function accepts two strings:

```text
name
language
```

and returns a string.

Examples:

```go
Hello("", "")
// Hello, World

Hello("Rob Pike", "Spanish")
// Hola, Rob Pike

Hello("Robert Griesemer", "French")
// Bonjour, Robert Griesemer
```

### Default name

If the caller supplies an empty name:

```go
if name == "" {
	name = "World"
}
```

the function uses `"World"` as the default.

This keeps the behavior inside `Hello()` instead of forcing every caller to handle the default themselves.

---

# Constants

The greetings and language names are defined as constants:

```go
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)
```

Constants are useful here because these values do not change during program execution.

They also avoid scattering **magic strings** throughout the code.

For example:

```go
case french:
```

is clearer than:

```go
case "French":
```

when the same value has semantic meaning throughout the program.

---

# `greetingPrefix()`

```go
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
```

This function is responsible only for determining the greeting prefix.

The logic is:

```text
French  → Bonjour,
Spanish → Hola,
Other   → Hello,
```

This is a useful example of **separation of responsibilities**.

`Hello()` handles the overall greeting while `greetingPrefix()` handles language selection.

---

## Named Return Values

This function uses a named return value:

```go
func greetingPrefix(language string) (prefix string)
```

The return variable is named `prefix`.

Therefore, this:

```go
return
```

returns the current value of `prefix`.

For example:

```go
prefix = frenchHelloPrefix

return
```

is equivalent in effect to:

```go
return frenchHelloPrefix
```

### Important

Named returns are a Go feature, but they should not automatically be used everywhere.

For short functions they can sometimes improve readability, but unnecessary named returns can make code harder to follow because the returned value is no longer visible at the `return` statement.

---

# `switch`

The language selection uses a `switch`:

```go
switch language {
case french:
	prefix = frenchHelloPrefix
case spanish:
	prefix = spanishHelloPrefix
default:
	prefix = englishHelloPrefix
}
```

This is clearer than repeatedly checking:

```go
if language == french {
	...
} else if language == spanish {
	...
}
```

The `default` case provides the English greeting for any unsupported or empty language.

---

# Exported vs Unexported Functions

The function:

```go
func Hello(...)
```

starts with an uppercase letter.

Therefore, it is **exported**.

The function:

```go
func greetingPrefix(...)
```

starts with a lowercase letter.

Therefore, it is **unexported**.

Go uses capitalization to control visibility between packages:

```text
Hello         → exported
greetingPrefix → unexported
```

`Hello()` is the public behavior of this package, while `greetingPrefix()` is an implementation detail.

---

# Testing

The project uses Go's standard `testing` package:

```go
import "testing"
```

No third-party testing framework is required.

Run the tests with:

```bash
go test
```

For more detailed output:

```bash
go test -v
```

---

# Subtests

The test uses `t.Run()`:

```go
t.Run("saying hello to people", func(t *testing.T) {
	// test
})
```

This allows several related scenarios to exist under the same `TestHello` function.

The current test cases cover:

```text
Normal greeting
Empty name
Spanish greeting
French greeting
```

This makes the test output descriptive and makes it easier to identify which behavior failed.

---

# Test Helper

Repeated assertion logic is extracted into:

```go
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Instead of repeating:

```go
if got != want {
	t.Errorf("got %q want %q", got, want)
}
```

in every subtest, the tests call:

```go
assertCorrectMessage(t, got, want)
```

---

## `testing.TB`

The helper accepts:

```go
testing.TB
```

rather than:

```go
*testing.T
```

`testing.TB` is an interface implemented by Go's testing and benchmarking types.

This makes the helper usable in both tests and benchmarks.

For this exercise, the important idea is simply:

> `testing.TB` allows a test helper to work with multiple testing types.

---

## `t.Helper()`

Inside the helper:

```go
t.Helper()
```

tells Go that this function is a test helper.

When the assertion fails, Go can report the failure at the caller's location rather than making the helper implementation look like the source of the failure.

This becomes particularly useful when you have many reusable test helpers.

---

# `%q` in Test Errors

The assertion uses:

```go
t.Errorf("got %q want %q", got, want)
```

`%q` formats strings with quotation marks.

For example:

```text
got "Hello, Bob" want "Hello, Alice"
```

This makes invisible or confusing characters easier to notice.

For example, a whitespace difference becomes much easier to spot:

```text
got "Hello, Bob " want "Hello, Bob"
```

---

# Why `Hello()` Returns a String

The `Hello()` function does not print anything.

It returns the greeting:

```go
return greetingPrefix(language) + name
```

The actual printing happens in `main()`:

```go
fmt.Println(Hello("", ""))
```

This separation is important.

### Logic

```go
Hello("Rob Pike", "Spanish")
```

produces:

```text
Hola, Rob Pike
```

### Side effect

```go
fmt.Println(...)
```

prints that value to the terminal.

Keeping these separate makes the logic easy to test without having to capture terminal output.

---

# TDD Discipline

The development process used in this exercise follows:

```text
1. Write a test
2. Make the compiler pass
3. Run the test
4. Confirm that it fails
5. Check that the failure message is meaningful
6. Write enough code to make the test pass
7. Refactor
```

This is commonly summarized as:

```text
Red → Green → Refactor
```

### Red

Write a test for behavior that does not exist yet.

The test should fail.

### Green

Implement the smallest amount of code necessary to make the test pass.

### Refactor

Improve the implementation while keeping the tests passing.

The important idea is that **the tests drive the design**, rather than being something added after the implementation is finished.

---

# Current Behavior

| Input                            | Output                      |
| -------------------------------- | --------------------------- |
| `("", "")`                       | `Hello, World`              |
| `("Ken Thompson!", "")`          | `Hello, Ken Thompson!`      |
| `("Rob Pike", "Spanish")`        | `Hola, Rob Pike`            |
| `("Robert Griesemer", "French")` | `Bonjour, Robert Griesemer` |
| `("Any Name", "Unknown")`        | `Hello, Any Name`           |

Unsupported languages fall back to English.

---

# Running the Exercise

From this directory:

```bash
go run .
```

Run tests:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

Run all tests recursively from the repository root:

```bash
go test ./...
```

---

# Key Takeaways

### Go fundamentals

- `package main` creates an executable package.
- `main()` is the executable entry point.
- Functions can accept multiple parameters.
- Functions can return values.
- Constants are declared with `const`.
- `switch` is useful for selecting between multiple cases.
- Uppercase identifiers are exported.
- Lowercase identifiers are unexported.

### Testing

- Go has a built-in testing framework.
- Test files use the `_test.go` suffix.
- Test functions start with `Test`.
- `t.Run()` creates subtests.
- `t.Helper()` marks reusable test helpers.
- `testing.TB` can make helpers reusable across tests and benchmarks.
- `go test` runs the tests.

### Design

The exercise demonstrates an important design principle:

```text
Pure logic
    ↓
Hello()
    ↓
String result
    ↓
Side effect
    ↓
fmt.Println()
```

The function produces a value, while `main()` decides what to do with that value.

That makes the core behavior easier to test, reuse, and change.

---

# Commands Cheat Sheet

```bash
# Run the program
go run .

# Run tests
go test

# Run tests with verbose output
go test -v

# Run all tests in the module
go test ./...

# Format Go files
gofmt -w .

# Inspect module dependencies
go list -m all

# Verify module dependencies
go mod verify
```

---

## Reference

- [Learn Go with Tests — Hello World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
