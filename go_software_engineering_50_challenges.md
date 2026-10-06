# Go Software Engineering Challenge Pack
## 50 Problems — Moderate → High

**Goal:** Move beyond syntax exercises. These challenges are designed to make you practice Go while gradually learning software-engineering concepts such as APIs, concurrency, persistence, testing, error handling, interfaces, caching, observability, rate limiting, background workers, and clean architecture.

## Rules of engagement

1. Use the standard library first. Avoid frameworks unless a challenge explicitly allows them.
2. Write tests for non-trivial logic.
3. Keep packages small and responsibilities clear.
4. Handle errors deliberately; don't ignore returned errors.
5. Add a README to every substantial project.
6. Use `go fmt`, `go vet`, and `go test ./...`.
7. For concurrency challenges, use the race detector: `go test -race ./...`.
8. Prefer context-aware functions for operations that can block.
9. Don't copy solutions from tutorials. Search documentation when you are stuck, but implement the idea yourself.
10. For every project, ask: **What happens when input is invalid, data is missing, the program is interrupted, or two things happen at the same time?**

---

# LEVEL 1 — Strong Go Foundations

## 1. Log File Analyzer
**Difficulty:** Moderate

Build a CLI that reads a large web-server log file and reports:
- total requests
- requests by HTTP method
- requests by status code
- top 10 requested paths
- top 10 IP addresses
- average response size

### Engineering requirements
- Process the file as a stream; don't load the whole file into memory.
- Separate parsing, aggregation, and presentation into packages/functions.
- Return useful errors for malformed lines.
- Add unit tests for the parser.

### Concepts
`bufio`, structs, maps, error handling, package design, streaming, testing.

### Stretch
Add filtering:
```bash
loganalyzer --status 404 --from "2026-01-01" server.log
```

---

## 2. Password Strength Analyzer
**Difficulty:** Moderate

Create a CLI that evaluates passwords using configurable rules.

Rules might include:
- minimum length
- uppercase/lowercase
- digits
- symbols
- repeated sequences
- common-password detection

### Engineering requirements
- Don't put all validation in `main`.
- Create a validator interface.
- Return structured validation results instead of strings only.
- Test each rule independently.

### Stretch
Allow rules to be loaded from a JSON configuration file.

### Concepts
interfaces, dependency injection, configuration, testing.

---

## 3. In-Memory Task Manager
**Difficulty:** Moderate

Build a command-line TODO application.

Commands:
```text
add "Learn interfaces"
list
complete 3
delete 3
search interfaces
```

Each task should have:
- ID
- title
- description
- status
- created time
- completed time

### Engineering requirements
- IDs must remain predictable.
- Separate storage from business logic.
- Create a `TaskRepository` interface.

### Stretch
Implement both:
- in-memory repository
- JSON-file repository

without changing the business logic.

### Concepts
interfaces, repositories, domain models, persistence abstraction.

---

## 4. JSON Database
**Difficulty:** Moderate+

Build a tiny database that stores arbitrary JSON records.

Support:
```text
insert
get
update
delete
list
```

Example:
```json
{"id": 10, "name": "Alice", "age": 24}
```

### Engineering requirements
- Persist data to disk.
- Handle corrupt database files.
- Avoid losing existing data if the process crashes during a write.

### Stretch
Implement atomic writes using a temporary file + rename.

### Concepts
serialization, persistence, atomicity, file systems, failure handling.

---

## 5. URL Shortener
**Difficulty:** Moderate+

Build a local URL-shortening service.

Example:
```text
POST /shorten
{
  "url": "https://example.com/articles/go"
}
```

Response:
```json
{
  "short_url": "/a8F2kL"
}
```

Then:
```text
GET /a8F2kL
```
redirects to the original URL.

### Requirements
- HTTP server
- JSON API
- unique short IDs
- input validation
- persistent storage

### Stretch
Track:
- number of clicks
- last accessed time
- creation time

### Concepts
HTTP, REST, JSON, persistence, routing, validation.

---

## 6. Configuration System
**Difficulty:** Moderate+

Build a configuration package supporting:
- defaults
- JSON configuration
- environment variables
- command-line overrides

Priority:
```text
CLI > environment > config file > defaults
```

### Engineering requirement
The rest of the application should not care where configuration came from.

### Stretch
Validate configuration at startup and print all invalid fields together.

### Concepts
configuration management, precedence, validation, clean boundaries.

---

## 7. Custom Logger
**Difficulty:** Moderate+

Create your own logging package.

Support:
```go
logger.Info("server started")
logger.Error(err, "database failed")
```

Levels:
- DEBUG
- INFO
- WARN
- ERROR

### Requirements
- timestamps
- structured fields
- configurable output
- configurable minimum level

### Stretch
Support JSON output:
```json
{
  "level": "ERROR",
  "message": "database failed",
  "request_id": "abc123"
}
```

### Concepts
interfaces, structured logging, dependency injection.

---

## 8. CLI Expense Tracker
**Difficulty:** Moderate+

Build:
```text
expense add 120 food "Lunch"
expense add 500 transport "Taxi"
expense list
expense summary
expense summary --month 2026-09
```

### Requirements
- persistence
- categories
- dates
- summaries
- input validation

### Stretch
Export CSV.

### Concepts
CLI design, persistence, aggregation, formatting.

---

## 9. File Backup Utility
**Difficulty:** Moderate+

Create:
```bash
backup ./documents ./backup
```

The program should:
- recursively scan directories
- copy files
- preserve directory structure
- detect changed files
- skip unchanged files

### Stretch
Use SHA-256 hashes to detect modifications.

### Hard mode
Run file copies concurrently with a bounded worker pool.

### Concepts
file systems, hashing, concurrency, worker pools.

---

## 10. Inventory Management System
**Difficulty:** Moderate+

Build a service for:
- products
- stock levels
- purchases
- sales
- low-stock alerts

Example:
```text
POST /products
POST /products/{id}/purchase
POST /products/{id}/sell
GET /products/{id}
```

### Important rule
A product must never have negative stock.

### Stretch
Record an immutable transaction history.

### Concepts
business rules, APIs, state transitions, domain modeling.

---

# LEVEL 2 — APIs, Databases, and Testing

## 11. REST API for a Blog
**Difficulty:** Moderate+

Build a blog API with:
- users
- posts
- comments
- tags

Support CRUD operations.

### Requirements
- proper HTTP status codes
- validation
- pagination
- filtering
- centralized error responses

### Stretch
Add soft deletion.

### Concepts
REST, HTTP semantics, pagination, API design.

---

## 12. SQLite Task API
**Difficulty:** Moderate+

Take your task manager and replace JSON storage with SQLite.

Implement:
- migrations
- CRUD
- indexes
- transactions

### Engineering challenge
Do not put SQL queries directly inside HTTP handlers.

### Concepts
SQL, transactions, repositories, migrations, layered architecture.

---

## 13. API Authentication
**Difficulty:** Moderate+

Add authentication to an API.

Features:
- registration
- password hashing
- login
- token/session authentication
- protected endpoints
- logout

### Security requirements
- Never store plaintext passwords.
- Never log passwords or tokens.
- Validate authentication failures consistently.

### Stretch
Implement refresh tokens.

### Concepts
authentication, authorization, security, sessions/tokens.

---

## 14. Role-Based Access Control
**Difficulty:** Moderate+

Add roles:
- admin
- editor
- user

Example:
```text
admin  -> everything
editor -> create/update posts
user   -> read posts
```

### Requirement
Authorization must happen in a reusable middleware/service rather than being duplicated in every handler.

### Stretch
Implement permissions such as:
```text
post.read
post.create
post.update
post.delete
```

### Concepts
middleware, authorization, policy design.

---

## 15. Pagination Engine
**Difficulty:** Moderate+

Build reusable pagination for an API.

Support:
```text
?page=2&limit=20
```

Return:
```json
{
  "data": [],
  "page": 2,
  "limit": 20,
  "total": 148,
  "total_pages": 8
}
```

### Hard mode
Implement cursor-based pagination.

### Concepts
API contracts, database pagination, cursor design.

---

## 16. Request Validation Package
**Difficulty:** Moderate+

Create a reusable validation library.

Support rules like:
```text
required
min
max
email
one_of
matches
```

Example:
```go
validator.Validate(user)
```

### Engineering requirements
Return structured errors:
```text
email -> invalid email
age   -> must be >= 18
```

### Stretch
Support nested structs.

### Concepts
reflection, interfaces, reusable libraries, error modeling.

---

## 17. HTTP Middleware Framework
**Difficulty:** Moderate+

Implement middleware for:
- logging
- recovery
- authentication
- request IDs
- timing

Make middleware composable:
```go
handler := Chain(
    RequestID,
    Logger,
    Recovery,
)(finalHandler)
```

### Concepts
higher-order functions, middleware, HTTP architecture.

---

## 18. Health Check Service
**Difficulty:** Moderate+

Create:
```text
GET /health
GET /ready
```

Health checks:
- database
- cache
- filesystem
- external service

Differentiate:
```text
healthy
degraded
unavailable
```

### Stretch
Run checks concurrently with timeouts.

### Concepts
interfaces, context, concurrency, service health.

---

## 19. API Rate Limiter
**Difficulty:** Moderate+

Implement a rate limiter.

Example:
```text
100 requests/minute/IP
```

Return:
```text
429 Too Many Requests
```

### Requirements
- concurrency safe
- configurable limits
- cleanup old clients

### Stretch
Implement:
- fixed window
- sliding window
- token bucket

and compare their behavior.

### Concepts
concurrency, mutexes, algorithms, time.

---

## 20. HTTP Client Library
**Difficulty:** Moderate+

Build a reusable API client package.

Requirements:
- base URL
- authentication
- request timeout
- retries
- JSON decoding
- error handling

Example:
```go
client.Users().Get(ctx, 42)
```

### Stretch
Add exponential backoff.

### Concepts
HTTP clients, context, retries, library API design.

---

# LEVEL 3 — Concurrency and Systems Thinking

## 21. Concurrent Web Crawler
**Difficulty:** Hard

Build a crawler that starts from a URL and discovers links.

Requirements:
- maximum depth
- maximum concurrent requests
- visited URL tracking
- timeouts
- cancellation

### Critical challenge
Multiple goroutines must not process the same URL.

### Stretch
Respect robots.txt.

### Concepts
goroutines, channels, mutexes, context, concurrency control.

---

## 22. Worker Pool
**Difficulty:** Hard

Build a generic worker pool.

Example:
```go
pool.Submit(job)
```

Features:
- configurable workers
- job queue
- graceful shutdown
- job results
- error propagation

### Hard mode
Support context cancellation.

### Concepts
goroutines, channels, lifecycle management, backpressure.

---

## 23. Concurrent File Hasher
**Difficulty:** Hard

Given a directory, calculate SHA-256 for every file concurrently.

Requirements:
- bounded concurrency
- no goroutine leaks
- cancellation on error
- deterministic output

### Stretch
Compare two directories and report:
- added files
- removed files
- changed files
- unchanged files

### Concepts
worker pools, pipelines, context cancellation.

---

## 24. Producer-Consumer Pipeline
**Difficulty:** Hard

Build a pipeline:

```text
File Reader
    ↓
Parser
    ↓
Validator
    ↓
Transformer
    ↓
Writer
```

Each stage should run concurrently.

### Requirements
- backpressure
- cancellation
- graceful shutdown
- error propagation

### Concepts
channels, pipelines, backpressure, cancellation.

---

## 25. Concurrent Cache
**Difficulty:** Hard

Implement an in-memory cache:

```go
cache.Set("user:42", value, 5*time.Minute)
cache.Get("user:42")
```

Features:
- TTL
- expiration
- concurrency safety
- cleanup

### Stretch
Implement LRU eviction.

### Hard mode
Add statistics:
```text
hits
misses
evictions
size
```

### Concepts
mutexes, maps, time, caching algorithms.

---

## 26. Job Scheduler
**Difficulty:** Hard

Build a background scheduler.

Support:
```text
run every 10 seconds
run at 12:00
run once after 5 minutes
```

### Requirements
- background execution
- cancellation
- graceful shutdown
- job status
- failure tracking

### Stretch
Persist scheduled jobs.

### Concepts
time, goroutines, lifecycle management, persistence.

---

## 27. Message Queue
**Difficulty:** Hard

Implement a local message queue.

Support:
```text
Publish(topic, message)
Subscribe(topic)
```

Requirements:
- multiple consumers
- message ordering per topic
- concurrent publishers
- graceful shutdown

### Stretch
Add acknowledgements and retries.

### Hard mode
Persist messages to disk.

### Concepts
pub/sub, concurrency, delivery semantics, persistence.

---

## 28. Event Bus
**Difficulty:** Hard

Create an application event bus.

Example:
```go
bus.Publish(UserRegistered{UserID: 42})
```

Handlers:
```text
send welcome email
create audit log
update analytics
```

### Challenge
Decide whether handlers run:
- synchronously
- concurrently
- sequentially per event

Document your decision.

### Concepts
events, interfaces, decoupling, concurrency.

---

## 29. Circuit Breaker
**Difficulty:** Hard

Build a circuit breaker for unreliable external services.

States:
```text
CLOSED
OPEN
HALF_OPEN
```

Rules:
- failures increase failure count
- too many failures → OPEN
- after timeout → HALF_OPEN
- successful recovery → CLOSED

### Stretch
Add metrics and configurable thresholds.

### Concepts
state machines, concurrency, resilience.

---

## 30. Retry Engine
**Difficulty:** Hard

Create a reusable retry package.

Support:
- max attempts
- exponential backoff
- jitter
- retryable errors
- context cancellation

Example:
```go
Retry(ctx, policy, operation)
```

### Engineering challenge
Prevent retry storms.

### Concepts
resilience, backoff, context, error classification.

---

# LEVEL 4 — Production-Style Projects

## 31. URL Shortener v2
**Difficulty:** Hard

Rebuild the URL shortener as a production-style service.

Add:
- PostgreSQL/SQLite
- authentication
- rate limiting
- analytics
- caching
- structured logs
- request IDs
- graceful shutdown
- tests

### Required architecture
Separate:
```text
HTTP
Service
Repository
Domain
Infrastructure
```

### Concepts
clean architecture, observability, persistence, resilience.

---

## 32. File Upload Service
**Difficulty:** Hard

Build an HTTP service allowing users to upload files.

Requirements:
- multipart uploads
- maximum file size
- MIME validation
- generated IDs
- metadata
- download endpoint

### Security
Do not trust the client-provided filename or MIME type.

### Stretch
Store metadata in a database and files on disk.

### Concepts
HTTP, security, storage abstraction, validation.

---

## 33. Image Processing Queue
**Difficulty:** Hard

Create an API where users upload images.

Instead of processing immediately:
```text
Upload
  ↓
Create Job
  ↓
Queue
  ↓
Worker
  ↓
Resize
  ↓
Store Result
```

Support job status:
```text
pending
processing
completed
failed
```

### Stretch
Run multiple workers.

### Concepts
asynchronous processing, queues, state machines.

---

## 34. Notification Service
**Difficulty:** Hard

Build a notification service supporting:
- email
- SMS mock
- in-app notifications

Use an interface:
```go
type NotificationSender interface {
    Send(ctx context.Context, n Notification) error
}
```

### Engineering challenge
The business logic should not know which delivery provider is being used.

### Stretch
Add retry + dead-letter queue.

### Concepts
interfaces, dependency injection, queues, adapters.

---

## 35. Audit Log System
**Difficulty:** Hard

Build an audit system recording:
- actor
- action
- resource
- timestamp
- IP
- metadata

Example:
```text
USER 42 deleted POST 91
```

### Requirements
Audit logs should be append-only.

### Stretch
Create an API to search audit logs with filters and pagination.

### Concepts
immutable data, security, persistence, querying.

---

## 36. Feature Flag Service
**Difficulty:** Hard

Build a service for feature flags.

Example:
```text
new_checkout = enabled
dark_mode = disabled
```

Support:
- global flags
- user-specific flags
- percentage rollout

Example:
```text
new_checkout -> 20% of users
```

### Hard challenge
Make percentage assignment deterministic for a user.

### Concepts
hashing, configuration, distributed-system concepts.

---

## 37. Metrics Library
**Difficulty:** Hard

Build a small metrics package.

Support:
- counters
- gauges
- histograms

Example:
```go
requests.Inc()
latency.Observe(duration)
```

Expose metrics over HTTP.

### Stretch
Export in a Prometheus-compatible format.

### Concepts
instrumentation, aggregation, HTTP, concurrency.

---

## 38. Distributed Lock Simulator
**Difficulty:** Hard

Build a program simulating multiple application instances trying to acquire the same lock.

Implement:
```text
Acquire(lock)
Release(lock)
```

Features:
- expiration
- ownership
- renewal

### Important
Document why an in-memory lock is not truly distributed.

### Stretch
Use Redis or another external store.

### Concepts
distributed systems, leases, race conditions, failure modes.

---

## 39. Leader Election Simulator
**Difficulty:** Hard

Simulate 5 nodes.

Each node:
- sends heartbeats
- detects failed leader
- attempts to become leader

### Requirements
Simulate:
- node failure
- network delay
- leader recovery

### Goal
Understand the problem before implementing a real consensus algorithm.

### Concepts
distributed systems, heartbeats, timeouts, leader election.

---

## 40. Mini Reverse Proxy
**Difficulty:** Hard

Build an HTTP reverse proxy.

Example:
```text
client
  ↓
Go proxy
  ↓
backend A / backend B
```

Support:
- forwarding
- load balancing
- health checks
- request logging

### Stretch
Implement round-robin and least-connections strategies.

### Concepts
HTTP internals, networking, load balancing, interfaces.

---

# LEVEL 5 — High Difficulty / Capstone Challenges

## 41. Mini Git
**Difficulty:** Very Hard

Build a simplified version-control system.

Commands:
```text
init
add
commit
log
status
```

Use content-addressed storage.

A commit should reference:
- parent
- author
- timestamp
- tree/files

### Stretch
Implement branches.

### Concepts
hashing, graphs, content-addressed storage, file systems.

---

## 42. Mini Docker
**Difficulty:** Very Hard

Build a simplified container runner for Linux.

Start with:
```bash
mycontainer run /bin/sh
```

Investigate and implement concepts such as:
- processes
- namespaces
- filesystem isolation
- resource limits

### Important
Do this only in a safe Linux environment and understand the security implications.

### Concepts
OS processes, namespaces, isolation, systems programming.

---

## 43. Mini Database
**Difficulty:** Very Hard

Build a tiny database engine.

Start with:
```text
SET key value
GET key
DELETE key
```

Then implement:
- WAL
- persistence
- recovery
- indexing
- compaction

### Stretch
Add transactions.

### Concepts
storage engines, durability, indexing, recovery.

---

## 44. Distributed Key-Value Store Simulator
**Difficulty:** Very Hard

Create multiple Go nodes that communicate over HTTP or TCP.

Support:
```text
PUT key value
GET key
DELETE key
```

Then explore:
- replication
- leader
- node failure
- stale reads

### Goal
You don't need perfect distributed consistency. The goal is to experience the problems.

### Concepts
networking, replication, consistency, failures.

---

## 45. Chat Server
**Difficulty:** Very Hard

Build a TCP or WebSocket chat server.

Features:
- users
- rooms
- join/leave
- broadcast
- private messages

### Requirements
- concurrent clients
- graceful disconnects
- slow-client handling

### Stretch
Persistent message history.

### Concepts
networking, concurrency, protocols, backpressure.

---

## 46. Task Queue Platform
**Difficulty:** Very Hard

Build a mini background-job platform.

API:
```text
POST /jobs
GET /jobs/{id}
POST /jobs/{id}/cancel
```

Workers execute jobs.

Support:
- retries
- delayed jobs
- priorities
- concurrency limits
- dead-letter jobs
- job history

### Stretch
Multiple worker processes.

### Concepts
distributed jobs, scheduling, retries, queues, persistence.

---

## 47. API Gateway
**Difficulty:** Very Hard

Build an API gateway that routes requests to different services.

Example:
```text
/api/users/* → user-service
/api/orders/* → order-service
/api/products/* → product-service
```

Add:
- authentication
- rate limiting
- request IDs
- retries
- circuit breakers
- metrics

### Stretch
Dynamic configuration reload.

### Concepts
microservices, resilience, routing, observability.

---

## 48. Mini Kubernetes Scheduler
**Difficulty:** Very Hard

Simulate a cluster with nodes.

Each node has:
```text
CPU
Memory
```

Jobs request:
```text
CPU: 2
Memory: 512MB
```

Your scheduler decides where to place them.

Implement:
- scheduling
- node failure
- rescheduling
- resource accounting

### Stretch
Add scheduling strategies:
- first fit
- least loaded
- best fit

### Concepts
scheduling algorithms, resource management, failure recovery.

---

## 49. Production-Grade Web Service
**Difficulty:** Very Hard

Build a complete service of your own choosing.

Examples:
- URL shortener
- expense tracker
- e-commerce backend
- issue tracker
- booking system
- learning platform

### It must include

**Application**
- REST API
- authentication
- authorization
- validation
- persistence

**Engineering**
- unit tests
- integration tests
- configuration
- migrations
- logging
- metrics
- health checks

**Reliability**
- timeouts
- graceful shutdown
- retries
- rate limiting
- error handling

**Documentation**
- README
- API documentation
- architecture diagram
- local setup instructions
- design decisions

### Goal
This should look like something you could discuss in a junior/mid-level backend interview.

---

## 50. Final Capstone — Build a Distributed Job Platform
**Difficulty:** Extreme

Build a small distributed job-processing platform.

### Architecture

```text
             ┌──────────────┐
             │    Client    │
             └──────┬───────┘
                    │
                    ▼
             ┌──────────────┐
             │  API Server  │
             └──────┬───────┘
                    │
                    ▼
             ┌──────────────┐
             │ Job Storage  │
             └──────┬───────┘
                    │
                    ▼
             ┌──────────────┐
             │ Job Queue    │
             └──────┬───────┘
                    │
          ┌─────────┼─────────┐
          ▼         ▼         ▼
       Worker 1  Worker 2  Worker 3
```

### Required features

#### API
```text
POST /jobs
GET /jobs
GET /jobs/{id}
POST /jobs/{id}/cancel
```

#### Job lifecycle
```text
QUEUED
  ↓
RUNNING
  ↓
COMPLETED

RUNNING
  ↓
FAILED
  ↓
RETRYING
  ↓
QUEUED
```

#### Reliability
- retries
- exponential backoff
- timeouts
- cancellation
- graceful shutdown
- dead-letter queue

#### Concurrency
- multiple workers
- bounded concurrency
- safe job claiming
- race-free state transitions

#### Persistence
Store:
- jobs
- attempts
- status changes
- timestamps

#### Observability
Expose:
- jobs processed
- jobs failed
- jobs retried
- queue depth
- worker count
- execution latency

#### Testing
Write:
- unit tests
- integration tests
- concurrency tests
- failure tests

Run:
```bash
go test ./...
go test -race ./...
```

### Final challenge

Simulate:
1. 10,000 jobs
2. 5 workers
3. random worker failures
4. random job failures
5. slow jobs
6. cancelled jobs
7. temporary database failures

Then answer:

> Can your system recover without losing jobs or executing the same job unexpectedly many times?

You will discover why real distributed job systems are difficult.

---

# Suggested Progression

Don't attempt all 50 randomly.

Use this path:

```text
1  → 2  → 3  → 4  → 5
                ↓
6  → 7  → 8  → 9  → 10
                ↓
11 → 12 → 13 → 14 → 15
                ↓
17 → 18 → 19 → 20
                ↓
21 → 22 → 23 → 24 → 25
                ↓
26 → 27 → 28 → 29 → 30
                ↓
31 → 32 → 33 → 34 → 35
                ↓
36 → 37 → 38 → 39 → 40
                ↓
41 → 43 → 45 → 46 → 47
                ↓
49 → 50
```

# Engineering Checklist

For every project, progressively ask yourself:

### Code quality
- Are responsibilities separated?
- Are functions small enough to understand?
- Are names clear?
- Are dependencies explicit?
- Can I test the business logic without starting the whole application?

### Errors
- What can fail?
- Can the caller recover?
- Should this error be wrapped?
- Am I accidentally hiding useful context?

### Concurrency
- What data is shared?
- Who owns it?
- Do I need a mutex or a channel?
- Can two goroutines perform the same operation?
- Can a goroutine leak?
- What happens during shutdown?

### APIs
- Are HTTP methods appropriate?
- Are status codes meaningful?
- Are requests validated?
- Is the API consistent?
- Are errors machine-readable?

### Persistence
- What if the process crashes during a write?
- What if the database is unavailable?
- What if the same operation happens twice?
- What needs a transaction?

### Testing
Aim for:
```text
Unit tests
Integration tests
Edge-case tests
Failure tests
Concurrency tests
```

### Production thinking
Ask:
```text
How do I know this is broken?
How do I stop it safely?
How do I debug it?
How does it behave under load?
What happens when a dependency fails?
What happens when two requests race?
```

# Recommended Go Tooling

As you progress, become comfortable with:

```bash
go fmt ./...
go vet ./...
go test ./...
go test -race ./...
go test -cover ./...
go mod tidy
go run .
go build .
```

Also learn to read Go documentation rather than relying exclusively on tutorials.

# Personal Rule

For the first 20 problems, don't optimize prematurely.

For problems 21–40, start thinking about:
- correctness
- concurrency
- failure modes
- interfaces
- observability

For problems 41–50, start thinking like a software engineer rather than someone completing a coding exercise.

The objective isn't to finish 50 projects.

The objective is to reach the point where, when someone gives you a vague requirement, you can turn it into a design, identify the failure modes, implement it in Go, test it, and explain your decisions.
