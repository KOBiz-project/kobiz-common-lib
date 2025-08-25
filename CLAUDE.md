# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a Go common library for the KoBiz project, containing shared data models, schemas, DTOs, and constants used across the KoBiz ecosystem. The library uses GORM for database ORM and provides standardized structures for telecommunications business operations.

## Project Structure

- `domain/` - Core domain models and DTOs
  - `models.go` - Common domain models like RequestParameter, FileStructure, UserToken
  - `dto.go` - Data transfer objects including CommonRequest, CommonResponse, ConnectionInfos
- `schemas/` - Database schema definitions using GORM
  - `model.go` - Device model schema with telecom carrier codes (SKT, KT, LGU)
  - `*.go` files for various business entities (members, orders, products, telecom plans, etc.)
- `cconst/` - Application constants and error codes
  - `consts.go` - Standardized error codes and response constants used across the KoBiz system

## Key Architecture Patterns

### Database Schema Design
- Uses GORM v1.25.12 for ORM
- All schemas embed `gorm.Model` for standard ID, CreatedAt, UpdatedAt, DeletedAt fields
- Custom table names defined via `TableName() string` methods
- Korean comments in GORM tags for business context
- Consistent naming: snake_case for database columns, camelCase for JSON

### Error Handling
- Centralized error constants in `cconst/consts.go`
- Standardized error codes like `OK`, `SUCCESS`, `DATA_NOT_FOUND`, etc.
- Korean business context in error constant comments

### Common Patterns
- `CommonResponse` struct for standardized API responses with result_code, error_desc, and data fields
- `CommonRequest` struct for standardized API requests with parameters, HTTP request, and JWT token
- Telecom carrier-specific fields (SKT, KT, LGU) throughout schemas

## Development Commands

This is a Go module library without build scripts. Standard Go commands apply:

```bash
# Install dependencies
go mod tidy

# Run tests (if any exist)
go test ./...

# Check for syntax errors
go vet ./...

# Format code
go fmt ./...
```

## Dependencies

- `gorm.io/gorm v1.25.12` - Primary ORM for database operations
- Standard Go libraries for HTTP, time, and multipart file handling

## Business Context

This library serves a Korean telecommunications business platform (KoBiz) with:
- Device model management across three major Korean carriers (SKT, KT, LGU+)
- Member management with admin levels and authentication
- Order processing for both wireless and internet services
- Telecom plan management per carrier
- File upload and AWS S3 integration support