# gRPC-Mini User Service (Go)

This repository is a small, production-style example of a **gRPC-based internal service written in Go**, designed to demonstrate **contract-first API design**, **safe Protobuf evolution**, and **reliable service-to-service communication**.

The goal of this project is not feature richness, but **correctness, clarity, and architectural discipline**.

---

## Why gRPC?

This service is intended to represent an **internal service boundary**, where:

- All clients are controlled
- Strong contracts are desirable
- APIs evolve over time
- Latency and predictability matter

gRPC + Protobuf were chosen to provide:

- **Strongly typed APIs** with compile-time guarantees
- **Explicit contracts** via `.proto` files
- **Safe backward-compatible evolution**
- **Clear error semantics** using gRPC status codes
- **First-class support for deadlines and cancellation**

For browser- or public-facing APIs, a REST or GraphQL gateway would typically sit in front of this service.

---

## Service Responsibilities

The User Service is responsible for:

- Creating users
- Fetching users by ID
- Updating user profile data

It explicitly does **not** handle:

- Authentication or authorization logic
- Presentation-layer concerns
- API gateway responsibilities

Those concerns are expected to live at the edge or in shared infrastructure.

---

## API Design Principles

This service follows a few core design rules:

### 1. Contract-first design
The `.proto` file is the source of truth.  
All clients and servers are generated from it.

### 2. Explicit request / response messages
Input messages are never reused as output messages, and domain models are not used directly as RPC inputs.

This prevents accidental coupling and allows safe API evolution.

### 3. Backward compatibility by default
- Fields are only added, never repurposed
- Field numbers are never reused
- Optional fields are used for PATCH-like semantics

### 4. Errors are not data
Operation outcomes are expressed via **gRPC status codes**, not booleans or sentinel values.

---

## Project Structure

```text
grpc-user-service/
├── proto/                 # Protobuf contracts (versioned)
├── internal/
│   ├── service/           # Domain logic (no gRPC dependencies)
│   ├── store/             # Persistence layer (in-memory example)
│   └── server/            # gRPC adapters & interceptors
├── cmd/
│   └── server/            # Application entry point
└── client/
    └── demo/              # Simple demo client
```

**Key ideas:**

- Business logic is transport-agnostic
- gRPC code is kept thin and focused
- Cross-cutting concerns live in interceptors

---

## Error Handling

The service uses gRPC status codes to clearly distinguish failure types:

- `InvalidArgument` — malformed or invalid input
- `NotFound` — requested user does not exist
- `AlreadyExists` — conflicting create operations
- `Internal` — unexpected server errors

Clients are expected to branch on status codes, not error strings.

---

## Deadlines & Cancellation

All gRPC calls are expected to include deadlines.

**Deadlines:**
- Prevent resource leaks
- Enable cancellation propagation
- Avoid cascading failures under partial outages

Server handlers respect `context.Context` and stop work when requests are canceled.

---

## Interceptors

The service demonstrates server-side interceptors for:

- Logging
- Cross-cutting concerns (auth, metrics, tracing would be added here)

This keeps business logic clean and predictable.

---

## Running the Service

### 1. Generate Protobuf code
```bash
protoc \
  --go_out=. \
  --go-grpc_out=. \
  proto/user/v1/user.proto
```

### 2. Start the server
```bash
go run ./cmd/server
```

The server listens on `:50051`.

### 3. Run the demo client
```bash
go run ./client/demo
```

---

## What This Project Intentionally Omits

To keep the example focused, the following are intentionally left out:

- Persistence beyond an in-memory store
- Authentication / authorization
- Docker / Kubernetes configuration
- CI/CD pipelines

These are orthogonal concerns and would be added based on system requirements.

---

## Intended Audience

This project is intended for:

- Engineers learning gRPC beyond basic tutorials
- Interview preparation for backend or distributed systems roles
- Demonstrating API design and system boundaries in Go

---

## Key Takeaways

- gRPC excels at internal, contract-driven APIs
- Protobuf enforces discipline around evolution
- Deadlines and structured errors are essential for reliability
- Clean boundaries matter more than feature count

This repository is meant to reflect how a small, well-designed internal service might look in a larger system.