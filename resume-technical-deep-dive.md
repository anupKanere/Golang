# Resume Technical Deep-Dive — Interview Reference

Precise, step-by-step technical explanations for every major technology and architecture pattern on [resume.html](file:///Users/prajakta/Desktop/Github/Golang/resume.html). Designed for direct, accurate answers in technical interviews.

---

## 1. Apache Kafka & Change Data Capture (CDC)

### How Kafka Works (Step-by-Step)
1. **Topics & Partitions:** A topic is a logical stream of events split into **partitions** (ordered, immutable, append-only commit logs). Ordering is guaranteed **only within a single partition**, not across the whole topic.
2. **Producers & Partitioning:** Producers write records to a topic. If a message key (e.g., `partner_id`) is provided, Kafka hashes the key (`murmur2(key) % num_partitions`) so all events for the same entity land in the same partition in order. If no key is provided, it uses sticky partitioning for batching efficiency.
3. **Brokers & Replication:** Each partition has one **Leader** broker and $N-1$ **Follower** brokers. All reads and writes go to the Leader. Followers replicate the log; followers caught up within `replica.lag.time.max.ms` form the **In-Sync Replicas (ISR)**.
4. **Durability (`acks`):**
   - `acks=0`: Fire and forget (fastest, possible data loss).
   - `acks=1`: Leader writes to its local log and acks (data lost if leader crashes before replication).
   - `acks=all` (`-1`): Leader waits for all ISRs to persist the record (paired with `min.insync.replicas >= 2` for zero data loss).
5. **Consumers & Offsets:** Consumers pull messages sequentially using a monotonic **offset** per partition. A **Consumer Group** distributes partitions among its instances: **1 partition is consumed by at most 1 consumer** in the group at a time (scaling max parallelism = number of partitions).
6. **Offset Commits & Delivery Semantics:**
   - **At-most-once:** Commit offset *before* processing. Crash after commit = lost message.
   - **At-least-once:** Process message *first*, then commit offset. Crash before commit = reprocessed on restart (requires **idempotent consumers**).
   - **Dead-Letter Queue (DLQ):** Non-retryable/malformed messages (poison pills) are routed to a DLQ topic after $N$ retries so they do not block the partition offset.

### How CDC (Change Data Capture) Works
1. **Database Transaction Log:** Instead of polling DB tables (`SELECT ... WHERE updated_at > ?`), a CDC connector (e.g., Debezium) tails the database's binary/write-ahead log (MySQL **binlog** or PostgreSQL **WAL**).
2. **Event Emission:** Every `INSERT`, `UPDATE`, or `DELETE` is captured as a structured event (`before`, `after`, `op`, `ts_ms`) and published to a Kafka topic partitioned by primary key.
3. **Why CDC over Polling:** Zero extra query load on the production DB, captures hard `DELETE`s, and preserves exact commit order.

---

## 2. Transactional Outbox Pattern & Idempotency

### The Problem (Dual-Write Failure)
If a service updates its database AND publishes to Kafka (or calls an external API like SuperLeap) in the same request:
* **Case A:** DB commits, Kafka publish fails → downstream is out of sync.
* **Case B:** Kafka publishes, DB rolls back → downstream has phantom state.

### Step-by-Step Solution (Transactional Outbox)
1. **Single Local ACID Transaction:** Inside the same DB transaction, write the business state change (e.g., `merchant_access_map`) **and** insert an event row into an `outbox` table (`id`, `aggregate_id`, `payload`, `status`, `created_at`). Either both commit or both roll back.
2. **Relay / Poller:** A separate worker (or CDC on the `outbox` table) reads unpublished rows from `outbox` and sends them to Kafka or the downstream API.
3. **Mark Processed:** Once acknowledged, the relay marks the row as `processed` (or deletes it).
4. **Idempotent Consumer:** Because the relay can crash *after* sending but *before* marking `processed`, the Outbox guarantees **at-least-once** delivery. The downstream consumer must be **idempotent**:
   - Store processed `event_id`s in a deduplication table/Redis (`SETNX`), OR
   - Perform idempotent upserts (`INSERT ... ON CONFLICT DO UPDATE` / conditional update with version/timestamp).

---

## 3. Go (Golang) Core Internals & Concurrency

### GMP Scheduler Model
* **G (Goroutine):** Lightweight user-space execution unit (~2 KB initial stack, grows dynamically).
* **M (Machine):** OS kernel thread managed by the OS scheduler.
* **P (Processor):** Logical context holding a local run queue of Gs (count = `GOMAXPROCS`, typically CPU cores).
* **How it works:** An M must acquire a P to execute Gs. If a G makes a blocking syscall, the M detaches from P; P attaches to another M (from idle pool or newly created) so other Gs keep running. If a P's local queue is empty, it **steals** half the Gs from another P's queue (work stealing).

### Channels & `select`
* **Unbuffered (`make(chan T)`):** Synchronous handoff—sender blocks until receiver is ready.
* **Buffered (`make(chan T, n)`):** Asynchronous up to capacity `n`—sender blocks only when full; receiver blocks only when empty.
* **Rules:**
  - Only the **sender** should close a channel (`close(ch)`).
  - Sending to a closed channel → **panic**.
  - Receiving from a closed channel → returns zero value immediately (`v, ok := <-ch` has `ok == false`).
  - Sending/receiving on a `nil` channel → **blocks forever** (useful in `select` to dynamically disable a case).

### `sync.Mutex` vs `sync.RWMutex` vs `sync/atomic`
* `sync.Mutex`: Exclusive lock for both reads and writes.
* `sync.RWMutex`: Multiple concurrent readers (`RLock`), exclusive writer (`Lock`). Use only when reads vastly outnumber writes and critical sections are non-trivial.
* `sync/atomic`: Lock-free CPU-level atomic instructions for simple counters, flags, or pointers (`atomic.Int64`, `atomic.Bool`).

### Bounded Concurrency (Worker Pool / `errgroup`)
Never spawn unbounded goroutines per request/item (causes OOM or file descriptor exhaustion).
1. **Semaphore pattern:** Buffered channel `sem := make(chan struct{}, maxWorkers)` or `golang.org/x/sync/semaphore`.
2. **`errgroup.Group` with `SetLimit(n)`:** Spawns at most `n` concurrent goroutines, propagates the first error, and cancels the shared `context.Context`.

### Slices & Interfaces
* **Slice Header:** 3 fields (24 bytes on 64-bit): `Data` pointer, `Len`, `Cap`. Passing a slice copies the header (cheap), so mutating elements mutates the backing array; however, `append` exceeding `Cap` allocates a new backing array.
* **Nil Interface Gotcha:** An interface holds `(type, value)`. An interface is `nil` **only** if *both* `type` and `value` are `nil`. Returning a typed `nil` pointer (`var p *MyErr = nil; return p`) as `error` produces a **non-nil** interface (`err != nil` is `true`).

### Garbage Collector
* Concurrent **tri-color mark-and-sweep** collector with write barriers (low pause times, typically sub-millisecond STW).
* **Reducing GC pressure:** Avoid unnecessary heap allocations (check with `go build -gcflags="-m"` escape analysis), pre-allocate slices (`make([]T, 0, n)`), and reuse temporary objects with `sync.Pool`.

---

## 4. REST APIs, Twirp RPC & Microservices Decomposition

### REST vs Twirp RPC vs gRPC
| Feature | REST (JSON/HTTP) | Twirp RPC | gRPC |
| :--- | :--- | :--- | :--- |
| **Contract** | OpenAPI/Swagger (optional) | Protobuf `.proto` (strict) | Protobuf `.proto` (strict) |
| **Transport** | HTTP/1.1 or HTTP/2 | HTTP/1.1 or HTTP/2 | Requires HTTP/2 |
| **Payload** | JSON (text) | Binary Protobuf **or** JSON | Binary Protobuf |
| **Streaming** | No (needs WS/SSE) | Unary request/response only | Unary + Bidirectional streaming |
| **Debuggability** | Easy (`curl`) | Easy (`curl` with JSON) | Harder (needs `grpcurl`) |

* **Why Twirp:** Gives strict Protobuf code generation (type safety across Go/PHP/Python) without needing HTTP/2 everywhere or complex gRPC load balancing.

### Monolith Decomposition (Strangler Fig Pattern)
1. **Identify Boundary:** Isolate the domain (e.g., Partner Reporting) and its data ownership.
2. **Build New Service:** Implement the new Go endpoint behind the API gateway/router.
3. **Shadow / Dual-Read Verification:** Optionally shadow traffic from the PHP monolith to the Go service and diff responses to guarantee 100% parity.
4. **Progressive Cutover:** Shift live traffic via feature flags/gateway routing (1% → 10% → 50% → 100%).
5. **Decommission:** Remove dead PHP routes and decouple shared database access.

### Downstream Client Resilience
* **Timeouts:** Always set `context.WithTimeout` + HTTP client `Timeout` (never rely on default infinite timeouts).
* **Retries:** Retry only **idempotent** operations on transient errors (`502`, `503`, `429`, network timeouts) using **exponential backoff + jitter** (prevents synchronized retry storms).
* **Circuit Breaker:** Trips to `Open` state when error rate exceeds a threshold, failing fast for a cooldown period before testing recovery in `Half-Open` state.

---

## 5. MySQL / Aurora MySQL Performance & Pagination

### Indexing & Query Optimization
1. **B+Tree Structure:** InnoDB stores table rows inside the **Clustered Index** (Primary Key). Secondary indexes store `(indexed_col, primary_key)`—if a query needs columns not in the secondary index, InnoDB performs a **double lookup** back to the clustered index.
2. **Composite Index Leftmost-Prefix Rule:** An index on `(a, b, c)` accelerates queries filtering on `(a)`, `(a, b)`, or `(a, b, c)`, but **not** `(b)` or `(c)` alone. Put high-selectivity equality columns first, range columns (`>`, `<`, `BETWEEN`) last.
3. **Covering Index:** Including all `SELECT`ed columns in the index avoids touching the clustered index (`Using index` in `EXPLAIN`).
4. **Common CPU Spike Causes:**
   - Full table scans (`type: ALL` in `EXPLAIN`) due to missing indexes or functions on indexed columns (e.g., `DATE(created_at) = ...` disables the index; use range `created_at >= ... AND created_at < ...`).
   - N+1 queries inside loops (fix with `WHERE id IN (...)` batching or `JOIN`s).
   - Large `OFFSET` scans (`LIMIT 10 OFFSET 100000` scans and discards 100,000 rows).

### Cursor / Keyset Pagination (Fixing Cross-Window Data Loss)
* **Why `OFFSET` or Timestamp-only Cursors Fail:**
  - `OFFSET` shifts rows if new records are inserted/deleted between pages.
  - Filtering by `created_at > last_ts` drops rows if multiple records share the exact same `last_ts` at the page boundary.
* **Keyset Pagination Fix:** Order by a deterministic unique tuple `(updated_at, id)` and query:
  ```sql
  WHERE (updated_at > :last_ts)
     OR (updated_at = :last_ts AND id > :last_id)
  ORDER BY updated_at ASC, id ASC
  LIMIT 500;
  ```
  *(Or in MySQL tuple syntax: `WHERE (updated_at, id) > (:last_ts, :last_id)`.)*

---

## 6. PostgreSQL vs DuckDB & Apache Parquet (OLTP vs OLAP)

### Why Row Storage (PostgreSQL) Fails on 200GB Audit Logs
* PostgreSQL stores data **row-by-row** (heap pages) optimized for point lookups and transactional updates (OLTP).
* Scanning 1–2 columns across 200GB forces reading *all* columns from disk, plus heavy B-tree index bloat and `VACUUM` overhead.

### How Apache Parquet + DuckDB Solves It (200GB → 10–15GB)
1. **Columnar Layout:** Values of the same column are stored contiguously. Queries reading 3 out of 50 columns skip 94% of the data (**column pruning**).
2. **High Compression (~15x reduction):** Similar data types sit together, enabling:
   - **Dictionary Encoding:** Repeated strings (e.g., event types, domain names) are stored once and replaced with small integers.
   - **Run-Length Encoding (RLE) & Bit-Packing:** Repeated/low-cardinality values compress drastically, followed by **ZSTD/Snappy** block compression.
3. **Predicate Pushdown (Zone Maps):** Parquet row-group headers store `min`/`max` statistics per column. DuckDB skips entire row groups if the `WHERE` filter falls outside `[min, max]`.
4. **Vectorized Execution:** DuckDB processes batches (vectors) of ~2048 values in CPU L1/L2 cache using SIMD instructions, running embedded in-process with zero network serialization overhead.
5. **Handling Immutability:** Parquet files are immutable—append new partitions (e.g., `tenant_id=.../date=YYYY-MM-DD/*.parquet`) and run periodic background compaction to merge small files.

---

## 7. Redis Caching Patterns & Pitfalls

### Caching Strategies
* **Cache-Aside (Lazy Loading):** App checks Redis → on miss, reads DB → writes to Redis with TTL → returns. (Most common; only caches requested data).
* **Write-Through:** App writes to Redis and DB synchronously (consistent, higher write latency).

### Three Classic Cache Failure Modes & Fixes
1. **Cache Stampede (Thundering Herd):** Hot key expires → 1,000 concurrent requests miss and hit the DB simultaneously.
   - **Fix:** Request coalescing (`golang.org/x/sync/singleflight`) so only 1 goroutine queries the DB while others wait for its result; or per-key distributed lock.
2. **Cache Penetration:** Requests for non-existent IDs always miss Redis and hit the DB.
   - **Fix:** Cache empty/null markers with a short TTL, or use a **Bloom Filter** in front of the DB.
3. **Cache Avalanche:** Many keys expire at the exact same second (e.g., loaded together with identical TTL).
   - **Fix:** Add random **jitter** to TTLs (`base_ttl + rand(0, 60s)`).

---

## 8. Messaging Systems: Kafka vs NATS vs RabbitMQ

| Dimension | Apache Kafka | NATS (Core / JetStream) | RabbitMQ (AMQP) |
| :--- | :--- | :--- | :--- |
| **Model** | Distributed append-only log (Pull) | Subject-based pub/sub (Push; JetStream adds persisted streams) | Smart broker with Exchanges & Queues (Push) |
| **Delivery** | At-least-once / Exactly-once (within Kafka) | Core: At-most-once (fire & forget)<br>JetStream: At-least-once | At-least-once (with `ack`/`nack`) |
| **Replayability** | Yes (retained by time/size; reset offsets anytime) | Core: No<br>JetStream: Yes | No (deleted once acked) |
| **Routing** | Topic + Partition key | Wildcard subjects (`orders.*.created`) | Complex routing (Direct, Topic, Fanout, Headers) |
| **Best For** | High-throughput event streaming, CDC, audit logs, replay | Ultra-low latency microservice RPC/events, edge/IoT | Complex task queues, per-message priority/routing, delayed jobs |

---

## 9. Docker (Go + Alpine), Kubernetes & Prometheus

### Docker Optimization for Go
1. **Multi-Stage Build:** Compile in `golang:1.25-alpine` (large SDK image), copy only the compiled binary into `alpine:3.22` (or `scratch`/`distroless`) for a ~15–25 MB final image with minimal CVE surface.
2. **`CGO_ENABLED=0`:** Produces a statically linked binary with no dependency on C libraries (`glibc` vs Alpine's `musl libc`), avoiding runtime `not found` errors.
3. **Gotchas:** Minimal images lack `ca-certificates` (breaks HTTPS calls) and `tzdata` (breaks `time.LoadLocation`)—explicitly copy or install them.

### Kubernetes Core Primitives
* **Pod / Deployment:** Manages replica sets and rolling updates (`maxSurge`, `maxUnavailable`).
* **Probes:**
  - `readinessProbe`: Determines if a pod receives traffic from the Service (e.g., DB/Kafka connections ready).
  - `livenessProbe`: Restarts the container if the process is deadlocked.
* **HPA (Horizontal Pod Autoscaler):** Scales pods based on CPU/memory or custom Prometheus metrics.

### Prometheus Monitoring
* **Pull Model:** Prometheus scrapes `/metrics` HTTP endpoints at a fixed interval.
* **Metric Types:**
  - **Counter:** Monotonically increasing (e.g., `http_requests_total`; query with `rate()`).
  - **Gauge:** Point-in-time value that goes up/down (e.g., active goroutines, memory usage).
  - **Histogram:** Buckets observations (e.g., request latency) to compute quantiles (`histogram_quantile(0.99, ...)`).

---

## 10. Security, Vulnerability Scanning & Compliance

### OpenSCAP (OSCAP) on Linux
* **XCCDF:** XML standard defining security checklists/benchmarks (e.g., CIS Benchmarks, PCI DSS).
* **OVAL:** Declarative XML definitions describing *how* to inspect system state (package versions, file permissions, kernel sysctls).
* **Workflow:** Go orchestrates `oscap xccdf eval` via `os/exec` with context timeouts, parses the resulting XML report into structured findings, and maps failures to compliance controls.

### Attack Surface Mapping (Subdomain & DNS Discovery)
1. **Passive Enumeration:** Query certificate transparency logs (`crt.sh`) and public DNS datasets without touching target infrastructure.
2. **Active Enumeration:** Concurrent DNS resolution (`A`, `AAAA`, `CNAME`, `MX`, `TXT`) using bounded worker pools and rate limiters to avoid DNS resolver throttling.
3. **Wildcard DNS Filtering:** Detect wildcard records (`*.example.com` resolving random non-existent subdomains to the same IP) so false positives are filtered out.

---

## 11. Integration Testing (SLIT) & Progressive Rollout (A/B)

### Service-Level Integration Testing (SLIT)
* **Unit Tests:** Mock external dependencies; fast, tests business logic in isolation.
* **SLIT (Integration Tests):** Boots the real service binary against real/containerized dependencies (MySQL, Kafka, Redis) and stubbed external HTTP/Twirp APIs (e.g., WireMock) to verify DB queries, migrations, serialization, and Kafka consumer/producer wiring end-to-end.
* **CI Quality Gates:** Run deterministically on every PR using isolated schemas/topics per test suite so parallel runs never collide.

### Progressive Rollout (1% → 100%)
1. **Gate by Deterministic Hash:** Route traffic based on `hash(merchant_id) % 100 < rollout_percentage` so a given merchant sees consistent behavior.
2. **Compare & Monitor:** Track p95/p99 latency, 5xx error rates, and DB CPU at each ramp stage (1% → 10% → 50% → 100%).
3. **Cleanup:** Remove the feature flag and legacy code path immediately after 100% stabilization to prevent dead-code drift.
