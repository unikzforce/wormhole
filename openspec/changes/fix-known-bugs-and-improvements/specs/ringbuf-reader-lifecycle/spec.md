## ADDED Requirements

### Requirement: Ring buffer reader SHALL be created once per goroutine lifetime
A `ringbuf.NewReader` instance MUST be created exactly once before entering any read loop and MUST be closed via `defer` at the goroutine scope (not inside the loop body). Creating a new reader on every iteration accumulates open file descriptors because `defer` is scoped to the enclosing function, not the loop body.

#### Scenario: MAC discovery goroutine starts up
- **WHEN** `handleNewDiscoveredEntriesRingBuffer` is called
- **THEN** exactly one `ringbuf.NewReader` is created, its `Close` is deferred at function scope, and the goroutine enters a `for` loop that calls `rd.Read()` repeatedly on the same reader

#### Scenario: Ring buffer read error
- **WHEN** `rd.Read()` returns a non-nil error inside the read loop
- **THEN** the goroutine logs the error and exits the loop (allowing the deferred `rd.Close()` to fire), rather than silently continuing

#### Scenario: File descriptor count stays constant during operation
- **WHEN** the switch agent processes 10,000 MAC discovery events
- **THEN** the number of open file descriptors held by the process does not grow relative to startup
