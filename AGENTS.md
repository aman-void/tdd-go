# AGENTS.md — tdd-go

Learning repo for Go TDD (Learn Go with Tests). READMEs are curated mini-book notes for future reference — do not rewrite them unprompted; propose drafts first and wait for approval.

## Layout — independent modules, no root module

- No `go.work`, no root `go.mod`. Each chapter is its own module:
  - `01_hello_world` (`package main`, executable)
  - `02_integers` (packages `adder`, `subtractor`, `iseven` in subdirs — one package per directory is enforced by Go)
  - `03_iteration` (`package iteration`)
- Root `README.md` is a stub. Chapter `README.md` is the real doc.

## Commands — run per chapter dir, never from root

```bash
# 01 (only executable chapter)
go run .
go test -v              # from 01_hello_world/

# 02 (multi-package)
go test ./... -v        # from 02_integers/

# 03
go test -v                                   # tests + examples
go test -v -run TestCompareStrings           # single test
go test -bench=. -benchmem                   # all benchmarks
go test -bench=Repeat -benchmem -run=^$      # single benchmark
gofmt -w .
go vet ./...
```

`go test ./...` from repo root does nothing (no root module) — `cd` into the chapter first.

## Conventions to follow

- Stdlib `testing` only, no frameworks.
- Tests: `TestX` in `*_test.go`; table tests as `[]struct{name,...}` + `for _, tt := range tests` + `t.Run(tt.name, ...)`; errors use `%q` (strings), `%d` (ints), `%t` (bools).
- Examples: `func ExampleX()` + `// Output: ...` (without it, compile-only). Multiple examples need suffix: `ExampleIsEven_odd`.
- Benchmarks use the new `for b.Loop()` API (Go 1.24+, repo uses go 1.27.1), not manual `b.N` loops. Quote the pattern in Powershell: `go test -bench="."`.
- Helpers take `testing.TB` + call `t.Helper()` so failures point at the caller.
- READMEs must stay portable: never hardcode host-specific system info (CPU model, `goos`/`goarch`, `cpu:` lines). Present benchmark numbers as illustrative `~X ns/op` with a "your numbers will vary, focus on relative order and `allocs/op`" note; strip machine identifiers from pasted output.
- Final code favors idiomatic stdlib (`strings.Repeat`, `strings.Compare`, `strings.Builder`); `iteration.go` keeps the `+=` → `Builder` → `Repeat` journey in comments — preserve it.
