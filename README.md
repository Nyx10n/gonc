# gonc (Go Netcat) - Micro-Tool 1

This is a Micro Tool 1 that I made in order to understand the concept of networking and how to use Go's standard libraries to build a bidirectional communication tool from scratch.

## Usage

**Start a listener (Server Mode):**

`./gonc -l <PORT>`

**Connect to a target (Client Mode):**

`./gonc <IP> <PORT>`

## What I Learned

While building this Micro Tool, I understood several backend and networking mechanics:

- Handling raw socket creation and connection states.
- Using GoRoutines to manage simultaneous, non-blocking read/write loops, allowing traffic to flow in both directions without freezing the terminal
- Replacing standard read/write arrays with `io.Copy` to stream data efficiently and prevent the RAM exhaustion.
- Implementing top-level OS signal trapping to catch `Ctrl+C` and execute graceful shutdowns.
