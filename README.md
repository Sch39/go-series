# Go Engineering Notes — Runnable Examples

Code examples for **[Go Engineering Notes](https://sch39.dev)** — a practical guide to Go semantics, concurrency, HTTP APIs, databases, and production patterns.

**Target audience:** Developers who already understand programming and want to use Go to build real applications and solve technical problems.

---

## How to Use This Repository

### 1. Reading an Article

Read the article on [sch39.dev](https://sch39.dev) by series name.

### 2. Running Code Examples

Each topic has its own Go module. Navigate to the topic folder and run:

```bash
cd G01-02-value-pointer
go run ./...
go test ./...
```

For concurrent code, always verify thread-safety:

```bash
go test -race ./...
```

## Publication Status

### Batch 1 — Go Fundamentals for Developers (Target: 16 articles)
- [x] G01-01 - Go Project Structure: Packages, Modules, and Imports
- [x] G01-02 - Value vs Pointer Semantics in Go
- [ ] G01-03 - Slices Explained: Length, Capacity, and Backing Arrays
- [ ] G01-04 - Maps in Go: Zero Values, Mutability, and Concurrency
- [ ] G01-05 - Interfaces in Go: Implicit Implementation and Practical Design
- [ ] G01-06 - Error Handling in Go: Wrapping, Checking, and Sentinel Errors
- [ ] G01-07 - `defer`, `panic`, and `recover`: When to Use Them
- [ ] G01-08 - Generics in Go: Practical Use Cases Without Overengineering
- [ ] G01-09 - Functional Options Pattern: When It Helps and When It Doesn't
- [ ] G01-10 - Go Context: Cancellation, Deadlines, and Request Scope
- [ ] G01-11 - `io.Reader` and `io.Writer`: The Interfaces Everything Is Built On
- [ ] G01-12 - Struct Embedding and Composition (Not Inheritance)
- [ ] G01-13 - Loop Variable Semantics and Closures
- [ ] G01-14 - Iterators with `range` over Functions
- [ ] G01-15 - Time in Go: `time.Time`, Monotonic Clock, Timezones, Timers
- [ ] G01-16 - JSON Pitfalls: `omitempty`, Numbers, Unknown Fields, Custom Marshal

### Batch 2 — Application Design & CRUD (Target: 15 articles)
- [ ] G02-01 — Building a REST API with `net/http`
- [ ] G02-02 — Request Validation and JSON Handling in Go
- [ ] G02-03 — Implementing CRUD with PostgreSQL and `pgx`
- [ ] G02-04 — Database Transactions in Go: Commit, Rollback, and Failure Handling
- [ ] G02-05 — Structuring a Go Application Without Overengineering
- [ ] G02-06 - Repository Pattern in Go: When to Use It
- [ ] G02-07 - Pagination, Filtering, and Sorting for REST APIs
- [ ] G02-08 - Middleware Design in Go HTTP Servers
- [ ] G02-09 - Graceful Shutdown for Go HTTP Servers
- [ ] G02-10 - Idempotent API Design and Safe Retry Handling
- [ ] G02-11 - Database Migrations in Go (goose/atlas)
- [ ] G02-12 - Type-Safe SQL with `sqlc` vs Hand-Written Queries
- [ ] G02-13 - Authentication: Sessions vs JWT in Go
- [ ] G02-14 - Authorization Middleware and Per-Resource Access Checks
- [ ] G02-15 - Consistent API Error Responses (RFC 9457 Problem Details)
- [ ] G02-16 - CORS and CSRF for Go APIs

---

## Standards & Quality Gates

All code passes these checks:

```bash
go vet ./...
go test ./...
go test -race ./...  # for concurrent code
```

**Code quality expectations:**
- ✓ Runnable with provided instructions
- ✓ Error handling (no silent failures)
- ✓ Resource cleanup (files, connections, goroutines)
- ✓ Context cancellation propagation where applicable
- ✓ Clear test coverage of happy path and failure cases
- ✓ No credentials or sensitive data
- ✓ Performance claims backed by benchmarks or profiles

---

## Learning Path

**Recommended order for beginners:**
1. G01-02 — Value vs Pointer Semantics (foundation)
2. G01-03 — Slices (memory model)
3. G01-05 — Interfaces (design)
4. G01-06 — Error Handling (robustness)
5. G02-01 — REST API (practical use case)

---

## References

- [Go Official Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Go Specification](https://go.dev/ref/spec)
- [Go Blog](https://go.dev/blog/) — release notes, best practices

---

## License & Attribution

Code and articles are published on **[sch39.dev](https://sch39.dev)**.

This repository is the source for examples articles. All content and code are available for learning and reference.