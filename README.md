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
