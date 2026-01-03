# GEDIS

> A Redis-inspired in-memory data store built from scratch in Go.

---

## What is Redis?

**Redis (Remote Dictionary Server)** is an open-source, in-memory data store  
(system that keeps its data **directly in RAM**) often used as a:

- Database
- Cache
- Message broker

It is known for its **speed, simplicity, and support for multiple data structures**.

---

## What makes Redis special?

- **Atomic operations**  
  Every Redis command is atomic — while one command is executing, another command does not partially execute on the same data.

- **In-memory storage with persistence**  
  Data is stored in memory for speed, but Redis can also save data to disk (**persistence**).

- **TTL (Time-to-Live)**  
  Keys can expire automatically after a specified time.

- **Pub/Sub**  
  Enables real-time communication between services.

---

## Concurrent Programming (High-level Idea)

**Concurrent programming** means doing **more than one thing at the same time**.

In servers like Redis, this includes:

- Handling multiple clients
- Processing multiple requests
- Managing I/O efficiently

> Note: Redis is mostly single-threaded for command execution, but uses concurrency for I/O and background tasks.
