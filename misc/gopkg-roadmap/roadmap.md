# Go Standard Library Learning Roadmap

A prioritized roadmap to learn Go's standard library effectively. The goal is to become familiar with at least half of it, starting from the most essential packages.

---

## 🪡 Tier 1: Core Programming (Start Here)
Foundational packages every Go developer uses:

- [`fmt`](https://pkg.go.dev/fmt) – Formatting I/O
- [`errors`](https://pkg.go.dev/errors) – Error handling
- [`strings`](https://pkg.go.dev/strings) – String manipulation
- [`strconv`](https://pkg.go.dev/strconv) – String<->number conversion
- [`bytes`](https://pkg.go.dev/bytes) – Byte slice operations
- [`math`](https://pkg.go.dev/math) – Basic math functions
- [`sort`](https://pkg.go.dev/sort) – Sorting functionality
- [`time`](https://pkg.go.dev/time) – Time and duration utilities

---

## 🧩 Tier 2: Working with Data
For file, data, and encoding operations:

- [`io`](https://pkg.go.dev/io) – Core I/O interfaces
- [`bufio`](https://pkg.go.dev/bufio) – Buffered I/O
- [`encoding/json`](https://pkg.go.dev/encoding/json) – JSON handling
- [`encoding/csv`](https://pkg.go.dev/encoding/csv) – CSV support
- [`encoding/base64`](https://pkg.go.dev/encoding/base64) – Base64 encoding
- [`encoding/hex`](https://pkg.go.dev/encoding/hex) – Hex encoding
- [`os`](https://pkg.go.dev/os) – File system, environment
- [`path/filepath`](https://pkg.go.dev/path/filepath) – File path tools

---

## 🌐 Tier 3: Networking and HTTP
For building web servers, APIs, and networking:

- [`net`](https://pkg.go.dev/net) – TCP/UDP networking
- [`net/http`](https://pkg.go.dev/net/http) – HTTP servers/clients
- [`net/url`](https://pkg.go.dev/net/url) – URL parsing
- [`context`](https://pkg.go.dev/context) – Context and cancellation

---

## 🔒 Tier 4: Concurrency & Synchronization
Crucial for writing idiomatic Go:

- [`sync`](https://pkg.go.dev/sync) – Sync primitives like Mutex
- [`sync/atomic`](https://pkg.go.dev/sync/atomic) – Low-level atomic ops
- [`runtime`](https://pkg.go.dev/runtime) – Go runtime access
- [`runtime/debug`](https://pkg.go.dev/runtime/debug) – Debugging tools

---

## 🧪 Tier 5: Testing and Tooling
For writing robust code and testing:

- [`testing`](https://pkg.go.dev/testing) – Unit testing framework
- [`testing/quick`](https://pkg.go.dev/testing/quick) – Property-based testing
- [`flag`](https://pkg.go.dev/flag) – Command-line argument parsing
- [`log`](https://pkg.go.dev/log) – Logging

---

## ⚙️ Tier 6: Useful Utilities and Internals
Good to know as your experience grows:

- [`reflect`](https://pkg.go.dev/reflect) – Reflection utilities
- [`regexp`](https://pkg.go.dev/regexp) – Regex engine
- [`hash`](https://pkg.go.dev/hash) – Hashing interfaces
- [`crypto/*`](https://pkg.go.dev/crypto) – Secure cryptographic tools
- [`math/rand`](https://pkg.go.dev/math/rand) – Random number generation
- [`math/big`](https://pkg.go.dev/math/big) – Arbitrary-precision math

---

## 🚧 Tier 7: Advanced / Niche
Specialized packages, learn on demand:

- [`go/parser`, `go/ast`, `go/token`](https://pkg.go.dev/go/parser) – Go code analysis
- [`plugin`](https://pkg.go.dev/plugin) – Runtime plugin loading
- [`text/template`, `html/template`](https://pkg.go.dev/text/template) – Templating engines
- [`embed`](https://pkg.go.dev/embed) – Embed static files
- [`syscall`](https://pkg.go.dev/syscall) – Low-level OS interaction (mostly replaced by `x/sys`)

---

## 📓 Learning Strategy

- Focus first on Tiers 1–4 — they cover most practical use cases.
- Use [pkg.go.dev/std](https://pkg.go.dev/std) as a reference.
- Build sample projects:
  - CLI tools → `fmt`, `os`, `flag`, `log`
  - HTTP servers → `net/http`, `encoding/json`, `context`
  - File processors → `bufio`, `encoding/csv`
  - Concurrent workers → `sync`, `channels`, `go`

---

Happy coding with Go! ⚡

