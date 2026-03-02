# AGENTS.md - AI Writer Development Guide

## Project Overview

AI Writer is an AI-assisted writing system with:
- **Frontend**: React 18 + TypeScript + Vite
- **Backend**: Go 1.20+ with Gin framework
- **Database**: MySQL/MySQL with GORM

---

## Build/Lint/Test Commands

### Frontend (in `frontend/` directory)

```bash
# Install dependencies
npm install

# Development server (port 5173, proxies /api to localhost:8080)
npm run dev

# Production build
npm run build

# Lint (ESLint)
npm run lint

# Preview production build
npm run preview
```

### Backend (in `backend/` directory)

```bash
# Install dependencies
go mod download

# Run development server
go run main.go

# Build binary
go build -o aiwriter

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a single test
go test -v -run TestFunctionName ./path/to/package
```

---

## Code Style Guidelines

### TypeScript (Frontend)

#### Imports
- Use path aliases: `@/` maps to `src/`
- Order imports: external libs → internal modules → relative imports
- Example:
```typescript
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useUserStore } from '@/store'
import api from '@/api'
import './Login.less'
```

#### Naming Conventions
- **Components**: PascalCase (e.g., `WorksList.tsx`, `Login.tsx`)
- **Hooks**: camelCase starting with `use` (e.g., `useUserStore`)
- **Interfaces/Types**: PascalCase (e.g., `User`, `ApiResponse<T>`)
- **Variables/Functions**: camelCase
- **Files**: kebab-case for non-component files

#### TypeScript Rules
- Strict mode enabled (`strict: true` in tsconfig)
- Use explicit types for function parameters and return values
- Use `unknown` for catch clause errors, then narrow type
- Avoid `any`; use `unknown` or proper generics

#### React Patterns
- Use functional components with hooks
- Use Zustand for global state (see `src/store/userSlice.ts`)
- Extract page components to `pages/` directory
- Extract reusable components to `src/components/`
- Use `.less` files for component styling

#### Error Handling
```typescript
try {
  const response = await api.post('/auth/login', formData)
  if (response.code === 0) {
    // success
  } else {
    setError(response.message || '登录失败')
  }
} catch (err: unknown) {
  setError('用户名或密码错误')
}
```

### Go (Backend)

#### Package Structure
```
backend/
├── main.go
├── api/v1/          # Route handlers
├── internal/
│   ├── config/      # Configuration
│   ├── dto/         # Data transfer objects
│   ├── handler/     # HTTP handlers
│   ├── middleware/  # Custom middleware
│   ├── model/       # Database models
│   ├── repository/  # Data access layer
│   └── service/     # Business logic
├── pkg/             # Reusable packages
│   ├── errors/      # Error definitions
│   └── utils/       # Utility functions
```

#### Naming Conventions
- **Packages**: lowercase, short (e.g., `handler`, `service`)
- **Files**: snake_case (e.g., `user_handler.go`)
- **Functions/Variables**: camelCase
- **Structs/Interfaces**: PascalCase
- **Constants**: PascalCase or SCREAMING_SNAKE_CASE

#### Error Handling
- Define custom errors in `pkg/errors/`
- Return structured responses using `pkg/utils/response.go`
- Use middleware for centralized error handling

#### Database
- Use GORM for ORM
- Define models in `internal/model/models.go`
- Use migrations for schema changes

---

## Configuration

### Frontend Environment Variables
Create `frontend/.env`:
```env
VITE_API_BASE_URL=/api/v1
```

### Backend Environment Variables
Create `backend/.env`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=123456
DB_NAME=aiwriter
JWT_SECRET=your-secret-key
```

---

## API Development Notes

- Frontend mock mode enabled by default (`MOCK_MODE = true` in `src/api/index.ts`)
- API proxy configured in `vite.config.ts`: `/api` → `http://localhost:8080`
- Use standard RESTful patterns: GET/POST/PUT/DELETE

---

## Common Tasks

### Adding a New Page
1. Create component in `src/pages/PageName/`
2. Add route in `src/App.tsx`
3. Add sidebar navigation if needed in `src/components/layout/Sidebar/`

### Adding a New API Endpoint
1. Add handler in `backend/api/v1/`
2. Register route in `backend/api/v1/router.go`
3. Update frontend API in `src/api/index.ts`

### Running Single Test
```bash
# Backend - run test matching name pattern
go test -v -run "^TestUser$" ./internal/service/

# Frontend - configure Vitest or Jest (not currently set up)
```
