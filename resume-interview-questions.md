# Resume Interview Preparation — Question Bank

Based on [resume.html](file:///Users/prajakta/Desktop/Github/Golang/resume.html). Use this document to draft and refine your detailed technical, system design, and behavioral (STAR) answers.

---

## 1. Header, Logistics, & Career Story (Non-Technical / HR)

### Location & Visa Status (`Haneda, Tokyo, Japan (Dependent Visa)` + `+91 7720888632`)
- [ ] **Work Authorization:** You are on a Dependent Visa in Japan—do you currently have the 28-hour/week work permit (*shikakugai katsudo kyoka*), or do you require company sponsorship to switch to a full-time Engineer/Specialist in Humanities/International Services visa?
- [ ] **Contact & Availability:** Your phone number is an Indian `+91` number—are you currently residing in Tokyo full-time, and what is your notice period / availability?
- [ ] **Language & Culture:** What is your Japanese language level (JLPT)? How comfortable are you working in a bilingual or Japanese-speaking engineering team?

### Career Transition & Timeline Gaps
- [ ] **Mechanical to Software (2021–2023):** You have a Diploma and B.E. in Mechanical Engineering, and your Awards section mentions **JSW Steel Ltd (2022–2023)**, which isn't listed under Experience. What was your role at JSW Steel, why did you leave core mechanical engineering, and why did you choose PG-DAC at CDAC?
- [ ] **Gap (June 2025 – September 2025):** What were you doing between leaving Connect Secure in June 2025 and joining Razorpay in September 2025?
- [ ] **Vendor/Contract Role (`Razorpay via InfoBahn Software World`):** How did this engagement work? Were you embedded directly in a Razorpay product team, attending their standups/on-call, or working in an outsourced vendor pod?
- [ ] **Leaving Razorpay (June 2026):** You were at Razorpay for ~10 months. Why did you leave in June 2026?

---

## 2. Razorpay Experience — Point-by-Point Questions

### Bullet 1: Salesforce to SuperLeap CRM Migration
> *Kafka CDC consumer for `partner_type` change events, Outbox-pattern handler for `merchant_access_map` sync, SuperLeap API client, `GET /v1/partner/sales_poc` endpoint.*

- [ ] **CDC & Kafka:** How did the Kafka CDC pipeline work (e.g., Debezium / Maxwell / internal CDC)? What happens if a CDC event arrives out of order or gets duplicated?
- [ ] **Kafka Reliability:** How did you handle consumer offsets, retries, dead-letter queues (DLQ), and poison-pill messages?
- [ ] **Outbox Pattern (Why):** Why did you use the Transactional Outbox pattern for `merchant_access_map` instead of publishing directly to Kafka or calling SuperLeap inside the DB transaction?
- [ ] **Outbox Pattern (How):** How did the outbox poller/relay work, and how did you guarantee at-least-once delivery and idempotency on the SuperLeap side?
- [ ] **API Client Design:** How did you design the SuperLeap API client (timeouts, circuit breakers, exponential backoff, rate-limit handling `429`)?
- [ ] **Endpoint & Cutover:** What does `GET /v1/partner/sales_poc` do under the hood—does it read from a local synced store or proxy to SuperLeap? How did you handle data consistency during the cutover before Salesforce was decommissioned?
- [ ] **Behavioral (Deadline & Parity):** The Salesforce decommission had a hard org-wide deadline. How did you validate data parity between Salesforce and SuperLeap before switching traffic?

### Bullet 2: `GET Pricing` API Optimization & Splitz A/B Rollout
> *Eliminated redundant DB and entity-origin API calls, reduced latency, Splitz A/B experiment 1% → 100%, cleaned up gating code.*

- [ ] **Profiling & Root Cause:** How did you identify the redundant DB and entity-origin API calls (distributed tracing, pprof, DB query logs, flamegraphs)?
- [ ] **Nature of Redundancy:** Why were those calls redundant—were they N+1 queries, repeated calls within the same request lifecycle, or stale/unnecessary hydration?
- [ ] **Impact Metrics:** By how much did latency drop (p50, p95, p99), and what was the QPS of `GET Pricing`?
- [ ] **Experimentation:** How did Splitz work (feature flag / traffic splitting by merchant ID or request ID)?
- [ ] **Safe Rollout:** What metrics did you monitor at 1%, 10%, 50%, and 100% to confirm "zero regressions" (latency, error rate, response payload diffing / shadow comparison)?
- [ ] **Tech Debt Cleanup:** Why is cleaning up experiment gating code important, and what risks happen if you leave dead flags in production?

### Bullet 3: Aurora MySQL CPU Spikes, Batch Comparator & Outbox Window
> *Fixed 3 production DB queries causing CPU spikes across 3 incidents, optimized comparator batch job cursor management to prevent cross-window data loss, reduced `api_outbox` fetch window from 7 days to 1 day.*

- [ ] **Incident Walkthrough:** Walk me through one of the 3 incidents: how did you trace the CPU spike to a specific query (`EXPLAIN` / `EXPLAIN ANALYZE`, slow query log, Performance Insights)?
- [ ] **Query Fixes:** Were the fixes adding composite indexes, rewriting `OR`/`LIKE`/subqueries, avoiding full table scans, or changing pagination?
- [ ] **Aurora Architecture:** Why does Aurora MySQL shared cluster CPU matter (reader vs writer instance contention, connection pool saturation)?
- [ ] **Cursor Management Bug:** What was the bug in the comparator batch job's cursor management (e.g., timestamp-based pagination missing rows with identical timestamps at window boundaries, `OFFSET` shifting, or inclusive/exclusive `>` vs `>=` bounds)?
- [ ] **Cursor Fix:** How did you fix it (e.g., keyset pagination using `(updated_at, id)`, overlapping lookback window + idempotent processing)?
- [ ] **Outbox Window (7 days → 1 day):** Why was scanning 7 days of `api_outbox` hurting CPU (missing index on `created_at`/`status`, large table scan, InnoDB buffer pool thrashing)?
- [ ] **Outage Edge Case:** What happens if an outage lasts longer than 1 day—did you have a separate backfill/DLQ mechanism so events older than 1 day aren't silently dropped?

### Bullet 4: SLIT Integration Testing & Go 1.25 / Alpine 3.22 Upgrade
> *27 integration test scenarios across `partner`, `commission`, `merchant`, `invoice` packages; CI quality gates; Go 1.25 golden images with Alpine 3.22.*

- [ ] **Testing Strategy:** What is the SLIT framework (service-level integration testing), and how does it differ from unit tests with mocks vs end-to-end tests?
- [ ] **Test Isolation:** How did you manage test data isolation and external dependencies (DB, Kafka, downstream Twirp/HTTP services—Docker Compose, Testcontainers, or WireMock)?
- [ ] **CI Reliability:** How did you prevent flaky tests in CI quality gates?
- [ ] **Go & Container Upgrade:** Why upgrade to Go 1.25 and Alpine 3.22? What breaking changes, runtime improvements, or CVE/security fixes did you have to watch out for (e.g., `CGO_ENABLED=0`, musl libc vs glibc differences in Alpine, tzdata/ca-certificates)?

### Bullet 5: Reporting Visibility Decomposition (PHP Monolith → Go)
> *Migrated partner reporting from PHP API monolith into Go Partnerships & Reporting services; decomposed 2 legacy APIs.*

- [ ] **Migration Strategy:** How did you decompose the 2 legacy PHP APIs without breaking existing clients (Strangler Fig pattern, API gateway routing, dual-writes or shadow reads)?
- [ ] **Language Translation:** What were the biggest differences/challenges translating legacy PHP logic (dynamic typing, implicit DB queries, coupled state) into strongly-typed Go microservices?
- [ ] **Service Boundaries:** How did you define service ownership boundaries between Partnerships and Reporting services?

### Bullet 6: On-Call Engineering & Incident Resolution
> *Resolved 2.5K stuck commission transactions, settlement discrepancies, merchant data retrieval issues via debugging and RCA.*

- [ ] **Stuck Transactions Incident:** Walk me through the **2.5K stuck commission transactions** incident: how were they detected, why did they get stuck (state machine deadlock, failed Kafka consumer, downstream timeout), and how did you safely replay/reconcile them without double-crediting?
- [ ] **Financial Accuracy:** How do you handle settlement discrepancies in fintech where financial accuracy is critical (idempotency keys, ledger reconciliation, audit trails)?
- [ ] **RCA & Prevention:** Tell me about a post-mortem / RCA you wrote: what was the root cause, and what permanent preventive action did you ship?

---

## 3. Connect Secure Experience — Point-by-Point Questions

### Bullets 1 & 2: Vulnerability & Compliance Scanning (Python, Go, OSCAP)
> *Python scanning/remediation for Windows & macOS (40% improvement); Go scanning for Linux using OSCAP; GDPR, HIPAA, ISO 27001, SOC 2, PCI DSS.*

- [ ] **Metric Definition:** You mention improving security risk management by **40%**. How was that 40% defined and measured? *(Prepare exact before/after baselines for 40%, 60%, 50%, and 70%.)*
- [ ] **Windows/macOS Remediation:** How did your Python agent execute remediation safely on Windows/macOS (registry/WMI/PowerShell vs bash/launchd, privilege escalation, rollback if a patch breaks something)?
- [ ] **Linux & OSCAP:** How did Go interface with OpenSCAP (OSCAP) on Linux (`os/exec` subprocesses, parsing XCCDF/OVAL XML reports, memory usage on large scans)?
- [ ] **Compliance Mapping:** How did you translate abstract compliance controls (like SOC 2 or PCI DSS) into deterministic technical checks?

### Bullets 3 & 5: Rust & Go APIs + Attack Surface Mapper
> *REST APIs in Rust and Go with goroutines; Attack Surface Mapper using Rust APIs and Python process management for subdomain enumeration and DNS discovery.*

- [ ] **Go vs Rust vs Python:** Why did you use Rust for some APIs and Go for others? What were the trade-offs in development speed, memory safety, and concurrency (Tokio/async in Rust vs goroutines/channels in Go)?
- [ ] **Cross-Language Integration:** How did Rust APIs communicate with Python process management in the Attack Surface Mapper (HTTP/gRPC, message queue, subprocesses, FFI)?
- [ ] **Bounded Concurrency in Go:** How did you bound goroutines during concurrent scanning/API processing (worker pools, `errgroup`, semaphores, `context.Context` cancellation) so you didn't exhaust file descriptors or memory?
- [ ] **DNS & Subdomain Discovery:** How does subdomain enumeration and DNS record discovery work (passive OSINT vs active brute-forcing, DNS rate limiting, handling wildcard DNS)?

### Bullet 4: DuckDB & Apache Parquet Optimization
> *Reduced 200GB PostgreSQL tables to 10–15GB Parquet files with 60% faster query performance for AD audit data.*

- [ ] **PostgreSQL Bottleneck:** Why was PostgreSQL struggling with 200GB of Active Directory (AD) audit data (row-oriented storage, index bloat, vacuum overhead)?
- [ ] **Parquet + DuckDB Internals:** Why Parquet + DuckDB? Explain columnar compression, predicate pushdown, and dictionary encoding that enabled a ~15x size reduction (200GB → 10–15GB).
- [ ] **Mutations & Partitioning:** How did you handle writes/updates/deletes to Parquet (immutable files, partitioning strategy e.g. by tenant/date, compaction of small files)?
- [ ] **Serving Queries:** How were queries served—was DuckDB embedded in-process inside Go/Python, and how did you handle concurrent analytical queries?

### Bullets 6, 7, 8, 9: Redis, NATS, Report Generation, DSA, Monitoring
- [ ] **NATS vs Kafka vs RabbitMQ:** You used NATS at Connect Secure and Kafka at Razorpay. When would you choose NATS over Kafka (persistence/replay, consumer groups, at-most-once vs at-least-once, JetStream)?
- [ ] **Redis Caching (50% improvement):** What caching pattern did you use (cache-aside, write-through), how did you handle cache invalidation/TTL, and how did you prevent cache stampedes?
- [ ] **Report Generation (70% less manual effort):** How did you generate large compliance PDFs/reports without OOMing or blocking API threads (async worker + NATS + object storage)?
- [ ] **Data Structures & Algorithms:** Bullet 8 says you optimized scanning workflows using efficient data structures and algorithms. Give a specific example (e.g., trie/radix tree for IP/CIDR/domain matching, hash maps, topological sort for dependency checks, or sliding windows).
- [ ] **On-Prem Monitoring:** How did the on-premises memory and SSL expiry monitoring service work securely behind client firewalls (push vs pull, TLS cert chain inspection)?

### International Client Assignment: Dubai Government Project
- [ ] **Scale & Scope:** Tell me about your two stints in Dubai (Oct–Nov 2024, May–Jun 2025). What was the scale of the data migration and security assessment?
- [ ] **Client & Environment Challenges:** What unique challenges did you face working on-site with a government client (air-gapped environments, strict data residency/compliance, stakeholder communication)?

---

## 4. Technical Skills Cross-Check (Listed in Skills, Not in Bullets)

- [ ] **Java:** Listed under Programming Languages—where did you use Java (CDAC or projects), and what is your current proficiency?
- [ ] **Twirp RPC:** Listed under Backend Technologies. What is Twirp (Protobuf-based RPC over HTTP/1.1 or HTTP/2, JSON + binary support), and how does it compare to gRPC?
- [ ] **TiDB:** Listed under Databases. What is TiDB (distributed MySQL-compatible HTAP database, Raft consensus, TiKV), and how did you use it?
- [ ] **RabbitMQ:** Listed under Caching & Messaging. Where did you use RabbitMQ vs NATS vs Kafka, and how do their delivery guarantees/routing (exchanges/queues) differ?
- [ ] **Kubernetes & Prometheus:** Listed under Cloud & DevOps. What is your hands-on experience with K8s (deployments, readiness/liveness probes, HPA, debugging pods) and Prometheus (PromQL, histograms vs counters, alerting)?

---

## 5. General Go & System Design Questions Likely to Follow

### Go Internals
- [ ] **Scheduler:** Explain the Go GMP scheduler model (Goroutines, Machines, Processors) and work stealing.
- [ ] **Channels:** Buffered vs unbuffered channels, closing rules, and what happens when reading/writing to closed or `nil` channels.
- [ ] **Synchronization:** `sync.Mutex` vs `sync.RWMutex` vs `sync/atomic`. When do you use channels vs mutexes?
- [ ] **Slices:** Slice internals (pointer, length, capacity), `append` reallocation, and common gotchas when passing slices to functions or slicing sub-slices.
- [ ] **Garbage Collector:** How does Go's tri-color mark-and-sweep GC work, and how do you reduce GC pressure (e.g., `sync.Pool`, escape analysis, stack vs heap)?
- [ ] **Interfaces:** Interface internals (`iface`/`eface`) and the typed-`nil` interface gotcha.
- [ ] **Context:** How does `context.Context` propagation work for timeouts, deadlines, and cancellation across goroutines and HTTP/RPC calls?

### Tailored System Design Prompts
- [ ] **Prompt 1:** Design an Outbox-based event sync system between two services that handles 10k events/sec with ordering guarantees.
- [ ] **Prompt 2:** Design a multi-tenant vulnerability scanner and compliance reporting pipeline that ingests 200GB/day of audit logs.
- [ ] **Prompt 3:** Design a distributed rate-limited API client for a third-party CRM with circuit breaking and retries.

---

## 6. Resume Polish Checklist (Before Sending Out)

- [ ] **JSW Steel Ltd (2022–2023):** Clarify or add context since it appears in *Awards* but not *Experience*.
- [ ] **Typo in Skills (line 334):** Fix spacing in `Go (Golang), Python ,Java , Rust`.
- [ ] **Phone vs Location:** Consider adding a Japanese `+81` phone number (if available) alongside or instead of `+91` to match `Haneda, Tokyo, Japan`.
- [ ] **Metrics Backup:** Document concrete before/after numbers for **40%**, **60%**, **50%**, and **70%** at Connect Secure.
