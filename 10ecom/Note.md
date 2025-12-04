# Pointer vs Value Semantics

## Pointer (`&application`)

- Allocated on **heap**
- Multiple refs point to **same memory**
- Changes are **shared** across all references
- **8 bytes** copy cost (just pointer)
- Use for: Large structs, mutable state, shared resources

```go
api := &application{config: cfg}  // Type: *application
```

## Value (`application`)

- Allocated on **stack**
- Passing creates a **full copy**
- Changes to copy **don't affect original**
- Copy cost = **entire struct size**
- Use for: Small structs, immutable data, isolation

```go
api := application{config: cfg}  // Type: application
```

## Method Receivers

```go
// Pointer receiver - modifies original
func (a *application) update() {
    a.config.addr = ":9090"  // ✅ Changes persist
}

// Value receiver - modifies copy only
func (a application) update() {
    a.config.addr = ":9090"  // ❌ Changes lost
}
```

## Config Field: Value vs Pointer

```go
type application struct {
    config config   // Each instance gets own copy
}

type application struct {
    config *config  // All instances share same config
}
```

| `config config`    | `config *config` |
| ------------------ | ---------------- |
| Copied into struct | Shared reference |
| Isolated changes   | Shared changes   |
| No nil checks      | Can check nil    |

## Best Practice

```go
// ✅ Recommended pattern
type application struct {
    config config      // Value: immutable, isolated
    db     *sql.DB     // Pointer: shared resource
    logger *log.Logger // Pointer: shared resource
}

app := &application{config: cfg}  // Pointer to app itself
```

**Why?**

- `&application`: Shared across handlers, needs mutation
- `config` (value): Read-only, small, prevents accidental changes
- `db`, `logger` (pointers): Shared resources across app

---

# Chi Router Middleware

## What is r.Use(middleware.X)?

**r.Use()** = Register middleware that runs **before** every route handler

**Middleware** = Function that wraps HTTP handlers to add behavior (logging, auth, recovery, etc.)

## Execution Order

Middleware executes in **chain order** (top to bottom):
```
Request → RequestID → RealIP → Logger → Recoverer → Timeout → Handler → Response
```

---

## middleware.RequestID

**Purpose**: Generates unique ID for each request

**How it works**:
```go
// Adds ID to request context
ctx := context.WithValue(r.Context(), middleware.RequestIDKey, "abc123")
```

**Usage**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    reqID := middleware.GetReqID(r.Context())
    log.Printf("[%s] Processing request", reqID)
}
```

**Why**: Track requests across logs/services in distributed systems

---

## middleware.RealIP

**Purpose**: Extract **real client IP** from proxies/load balancers

**Problem**: 
- Direct connection: r.RemoteAddr = "203.0.113.5:54321" ✅
- Behind proxy: r.RemoteAddr = "10.0.0.1:12345" ❌ (proxy IP, not client)

**Solution**: Reads headers:
- X-Forwarded-For: 203.0.113.5, 192.168.1.1
- X-Real-IP: 203.0.113.5

**Sets**: r.RemoteAddr = "203.0.113.5" (actual client IP)

**Why**: Rate limiting, geolocation, security

---

## middleware.Logger

**Purpose**: Log every HTTP request/response

**Logs**:
- HTTP method (GET, POST, etc.)
- Request path (/api/users)
- Status code (200, 404, 500)
- Response time (125ms)
- Request ID

**Output example**:
```
2025/12/04 10:30:45 "GET /api/users HTTP/1.1" 200 125ms
2025/12/04 10:30:46 "POST /api/login HTTP/1.1" 401 23ms
```

**Why**: Debugging, monitoring, auditing

---

## middleware.Recoverer

**Purpose**: Catch **panics** in handlers to prevent server crash

**Without Recoverer**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    panic("something broke")  // ❌ Server crashes
}
```

**With Recoverer**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    panic("something broke")  // ✅ Returns 500, logs error, server stays up
}
```

**What it does**:
1. Catches panic with defer recover()
2. Logs stack trace
3. Returns 500 Internal Server Error
4. Server continues running

**Why**: Production resilience - one bad request shouldn't kill the server

---

## middleware.Timeout(60 * time.Second)

**Purpose**: Cancel long-running requests automatically

**How it works**:
```go
ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
defer cancel()
```

**After 60 seconds**:
- ctx.Done() channel closes
- Returns 503 Service Unavailable
- Handler should check ctx.Done() and stop work

**Example**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    select {
    case <-time.After(120 * time.Second):  // Slow operation
        w.Write([]byte("done"))
    case <-r.Context().Done():  // Timeout triggered at 60s
        return  // Stop processing
    }
}
```

**Why**: Prevent resource exhaustion from slow clients/operations

---

## Middleware Chain Concept

```go
r.Use(middleware.RequestID)   // 1st: Assign ID
r.Use(middleware.RealIP)      // 2nd: Fix IP
r.Use(middleware.Logger)      // 3rd: Start logging
r.Use(middleware.Recoverer)   // 4th: Catch panics
r.Use(middleware.Timeout(60)) // 5th: Set deadline

r.Get("/", handler)  // 6th: Your handler runs last
```

**Request flow**:
```
Client Request
    ↓
[RequestID] - Generates "req-abc123"
    ↓
[RealIP] - Sets RemoteAddr to actual client IP
    ↓
[Logger] - Starts timer, logs method/path
    ↓
[Recoverer] - Sets up panic recovery
    ↓
[Timeout] - Creates 60s context deadline
    ↓
[Handler] - Your business logic runs
    ↓
[Logger] - Logs status + response time
    ↓
[RequestID] - (cleanup if needed)
    ↓
Response to Client
```

---

## Key Concepts

**Middleware = Wrapper Function**:
```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)  // Call next middleware/handler
        log.Printf("took %v", time.Since(start))
    })
}
```

**Context carries data down the chain**:
```go
// RequestID adds to context
ctx := context.WithValue(r.Context(), "reqID", "abc")

// Logger reads from context
reqID := r.Context().Value("reqID")
```

**Order matters**:
- Logger should be early to measure full request time
- Recoverer should wrap handlers to catch panics
- Timeout should be after logging setup but before heavy work
