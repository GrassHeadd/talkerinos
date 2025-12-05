# Talkerinos

A tiny place on the internet for me to yap, rant, and brain-dump.  
Frontend is a separate app – this repo is the Go backend that powers it.

---

## What is this?

Talkerinos is a minimal blog engine, I write posts (probably overshare).

---

## Features (v1)

- To be determined lol

---

## Tech Stack

### Language / Runtime
- **Go** – latest stable (1.22+)

### Web Framework
- **Gin** – fast, minimal, great for JSON APIs

### Database
- **PostgreSQL** – even for a solo blog; good practice for real systems

### DB Access
- **sqlc** – write raw SQL, get type-safe Go code generated
- **golang-migrate** – schema migrations (`CREATE TABLE`, `ALTER TABLE`, etc.)

### Auth
- **JWT** (access token only, for now)
- **bcrypt** password hashing (`golang.org/x/crypto/bcrypt`)

### Config
- Env-based config → `Config` struct:
  - DB URL
  - HTTP port
  - JWT secret
- Just `os.Getenv` + tiny helper functions (no heavy config lib)

### Logging
- **log/slog** (std lib) for now
- Can upgrade to **uber-go/zap** later for structured JSON logs

### Infra / Dev
- **Docker** – containerize the API
- **docker-compose** – run API + Postgres locally

---

## Why this exists

I talk a lot.  
Now my Go backend can listen.
