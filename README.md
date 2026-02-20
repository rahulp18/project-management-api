# Project Management API

A Go-based REST API with clean architecture.

## Goals

We'll build:

- **Register** - User registration system
- **Login** - User authentication
- **JWT Auth** - Token-based security
- **Projects** - Project management
- **Tasks** - Task tracking
- **Role-based Access** - Permission system
- **Pagination** - Efficient data loading
- **Middleware** - Request/response handling
- **Clean Architecture** - Maintainable code structure

## Features

- User Registration & Login
- JWT Authentication
- Projects & Tasks Management
- Role-based Access Control
- Pagination & Middleware

---

## Phase 1 � Project Setup

- [ ] Folder structure
- [ ] Go modules
- [ ] Environment config
- [ ] Docker + PostgreSQL
- [ ] Migration tool

## Phase 2 � Auth System

- [ ] Users table
- [ ] Password hashing (bcrypt)
- [ ] JWT tokens
- [ ] Auth middleware

## Phase 3 � Projects + Tasks

- [ ] Relational schema
- [ ] Authorization rules

## Phase 4 � Production Polish

- [ ] Logging
- [ ] Error response format
- [ ] Pagination
- [ ] Testing

---

## Architecture Principles & Guidelines

### 🏗️ Clean Architecture

- **No business logic in handlers** - Keep handlers thin, delegate to services
- **No DB queries in handlers** - Repository pattern for data access
- **Layer separation** - Handler → Service → Repository
- **Dependency inversion** - Depend on interfaces, not concrete implementations

### 🔒 Security & Auth

- **No JWT parsing in handlers** - Use middleware for token validation
- **Context-based auth** - Pass user context through request lifecycle
- **Secure defaults** - Fail closed, validate inputs

### 💾 Data Management

- **No direct SQL in services** - Use repository abstractions
- **Transaction management** - Handle database transactions properly
- **Query optimization** - Use proper indexes and efficient queries

### 🔧 Code Quality

- **Everything injected via constructor** - Explicit dependencies
- **No global variables** - Pass dependencies explicitly
- **Proper error wrapping** - Use `fmt.Errorf` with `%w` verb for error chains
- **Context passed everywhere** - Enable cancellation and timeouts
- **No panic** - Return errors, don't crash the application

### 📡 API Standards

- **Consistent response format** - Standardized JSON structure
- **HTTP status codes** - Use appropriate codes (200, 201, 400, 401, 404, 500)
- **Input validation** - Validate all incoming data
- **Structured logging** - Use structured logs with appropriate levels

### 🧪 Testing & Observability

- **Unit tests** - Test business logic in isolation
- **Integration tests** - Test API endpoints end-to-end
- **Mocking** - Use interfaces for testable code
- **Health checks** - Implement readiness and liveness probes
