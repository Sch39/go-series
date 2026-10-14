# Go Engineering Notes — Runnable Examples

Code examples for **[Go Engineering Notes](https://sch39.dev)** — a practical guide to Go semantics, concurrency, HTTP APIs, databases, and production patterns.

**Target audience:** Developers who already understand programming and want to use Go to build real applications and solve technical problems.

---

## Series Overview

Go Engineering Notes is organized into **10 major sections**, each focusing on a specific aspect of Go development:

| ID | Series | Focus |
|----|--------|-------|
| **G01** | **Go Fundamentals for Developers** | Semantics (value vs pointer), packages, slices, maps, interfaces, errors, context |
| **G02** | **Application Design & CRUD** | REST APIs with `net/http`, PostgreSQL with `pgx`, validation, transactions, auth |
| **G03** | **Concurrency & Goroutines** | Goroutines, channels, mutexes, race conditions, deadlocks, goroutine leaks, worker pools |
| **G04** | **Memory, Performance & Leaks** | Memory allocation, leak detection, profiling with `pprof`, benchmarking, optimization |
| **G05** | **HTTP & Network Protocols** | HTTP lifecycle, streaming, file uploads (TUS), SSE, WebSocket, gRPC, reverse proxy |
| **G06** | **Database & Data Processing** | Connection pooling, N+1 queries, batch inserts, pagination, transactions, optimistic locking |
| **G07** | **Background Jobs & Distributed Systems** | Worker lifecycle, retries, idempotency, job queues, graceful shutdown, outbox pattern |
| **G08** | **Testing & Debugging** | Table-driven tests, HTTP testing, integration tests, race detection, fuzzing, profiling |
| **G09** | **Production & Deployment** | Configuration, structured logging, health checks, observability, security, containers, CI |
| **G10** | **Go Implementation Recipes** | Practical patterns: ZIP streaming, CSV parsing, idempotency keys, email delivery, CLI tools |

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

### Batch 1 — Go Essentials (Target: 6 articles)
- [ ] G01-03 — Slices Explained: Length, Capacity, and Backing Arrays
- [ ] G01-02 — Value vs Pointer Semantics in Go
- [ ] G01-05 — Interfaces in Go: Implicit Implementation and Practical Design
- [ ] G01-06 — Error Handling in Go: Wrapping, Checking, and Sentinel Errors
- [ ] G01-11 — `io.Reader` and `io.Writer`: Foundation for Upload & Streaming
- [ ] G01-10 — Go Context: Cancellation, Deadlines, and Request Scope

### Batch 2 — Practical REST API (Target: 5 articles)
- [ ] G02-01 — Building a REST API with `net/http`
- [ ] G02-02 — Request Validation and JSON Handling in Go
- [ ] G02-03 — Implementing CRUD with PostgreSQL and `pgx`
- [ ] G02-04 — Database Transactions: Commit, Rollback, and Failure Handling
- [ ] G08-02 — Testing HTTP Handlers with `httptest`

### Batch 3 — Concurrency & Leaks (Target: 6 articles)
- [ ] G03-01 — Goroutines Explained: Execution, Scheduling, and Lifecycle
- [ ] G03-02 — Channels vs Mutexes: Choosing the Right Synchronization
- [ ] G03-03 — Race Conditions in Go and How to Detect Them
- [ ] G03-04 — Common Goroutine Leaks and How to Prevent Them
- [ ] G03-05 — Worker Pool Pattern with Bounded Concurrency
- [ ] G04-06 — Profiling Go Applications with `pprof`

### Batch 4 — HTTP Streaming & Uploads (Target: 6 articles)
- [ ] G05-06 — Large File Uploads Without Loading Everything into Memory
- [ ] G05-07 — Resumable Uploads: Why Multipart Upload Is Not Enough
- [ ] G05-08 — Implementing the TUS Protocol in Go
- [ ] G05-12 — Implementing Server-Sent Events (SSE) in Go
- [ ] G05-17 — Implementing WebSocket in Go
- [ ] G05-18 — WebSocket Read/Write Pumps and Connection Lifecycle

### Batch 5 — Production Patterns (Target: 6 articles)
- [ ] G07-01 — Building Background Jobs in Go
- [ ] G07-02 — Retry Strategies, Exponential Backoff, and Jitter
- [ ] G07-03 — Idempotency in Background Jobs
- [ ] G02-09 — Graceful Shutdown for Go HTTP Servers
- [ ] G09-02 — Structured Logging with `log/slog`
- [ ] G09-03 — Health Checks: Readiness and Liveness

**Current Progress:** Foundational modules and tests under development.

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