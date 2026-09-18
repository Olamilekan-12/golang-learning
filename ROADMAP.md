# Go Learning Roadmap

Zero to professional, project-based, following the official Go documentation
(A Tour of Go → Effective Go → Go Modules Reference → the standard library docs)
as curriculum. Each numbered folder is a small, runnable project for one topic.

Status legend: ✅ done · 🚧 current · ⬜ upcoming

## Phase 0 — Toolchain
- ✅ Install Go (Homebrew), verify `go version`
- ✅ Init git repo, `.gitignore`, GitHub remote (`origin` → github.com/Olamilekan-12/golang-learning)

## Phase 1 — Language fundamentals (Tour of Go: Basics)
- ✅ `01-hello-world` — packages, `go run` vs `go build`, modules (`go mod init`)
- ✅ `02-variable-types` — var/const, zero values, type inference, `:=`
- ✅ `03-control-flow` — if/for/switch (no while/ternary in Go)
- ✅ `04-functions` — multiple returns, named returns, variadics, closures, defer

## Phase 2 — Composite types (Tour of Go: More types)
- ✅ `05-arrays-slices` — arrays vs slices, append, aliasing, copy
- ✅ `06-maps` — CRUD, comma-ok idiom, randomized iteration order
- ✅ `07-structs-methods` — value vs pointer receivers, pointers (&, *)
- ✅ `08-interfaces` — implicit satisfaction, any, type switches

## Phase 3 — Errors & panics
- ✅ `09-error-handling` — error interface, wrapping (%w), sentinel errors, errors.Is
- ✅ `10-panic-recover` — when (not) to use them

## Phase 4 — Concurrency (Tour of Go: Concurrency)
- ✅ `11-goroutines-channels` — goroutines, sync.WaitGroup, buffered/unbuffered channels
- ✅ `12-select-context` — select, time.After timeouts, context cancellation
- ✅ `13-sync-primitives` — race conditions, `-race`, Mutex, atomic

## Phase 5 — Testing & tooling
- ✅ `14-testing` — go test, table-driven tests, subtests, coverage
- ✅ `15-benchmarks` — go test -bench, -benchmem, pprof
- ✅ `16-tooling` — go doc, golangci-lint (plus gofmt, go vet)

## Phase 6 — Modules & dependencies
- ✅ `17-modules` — go get, go.mod/go.sum, semver, go mod tidy, module cache

## Phase 7 — Standard library deep dive
- ✅ `18-net-http` — handlers, routing with methods, query params, status codes
- 🚧 `19-encoding-json`
- ⬜ `20-io-and-files`
- ⬜ `21-database-sql` — sqlite/postgres driver

## Phase 8 — Generics
- ⬜ `22-generics` — type parameters, constraints

## Phase 9 — Idiomatic Go & project layout
- ⬜ `23-project-structure` — standard Go project layout, Effective Go conventions

## Phase 10 — Capstone project 🏆
- ⬜ `capstone/` — full service (proposal: concurrent URL-shortener + analytics API —
  REST API, Postgres, background workers, structured logging, tests, Docker, CI).
  Final design confirmed once we get there.

---

## Git practices track (introduced progressively)
- ✅ `git init`, `.gitignore`, meaningful commits, `--amend` on unpushed commits
- ✅ Commit hygiene: small, atomic, one concern per commit
- ✅ GitHub remote via `gh repo create`, `git push -u origin main`
- ⬜ Branching: feature branches, `main` always green
- ⬜ Merging vs rebasing, resolving conflicts
- ⬜ Tags & releases (semver)
- ⬜ Pull requests, code review workflow, CI (GitHub Actions)
- ⬜ Undo tools: `git restore`, `git revert` vs `git reset`, reflog
