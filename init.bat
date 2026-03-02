@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo === AI Writer Development Environment Setup ===
echo.

rem Check Node.js version
where node >nul 2>&1
if %errorlevel% equ 0 (
    for /f "delims=" %%i in ('node --version') do set NODE_VERSION=%%i
    echo ✓ Node.js found: %NODE_VERSION%
) else (
    echo ✗ Node.js not found. Please install Node.js 16+ from https://nodejs.org/
    exit /b 1
)

rem Check Go version
where go >nul 2>&1
if %errorlevel% equ 0 (
    for /f "delims=" %%i in ('go version') do set GO_VERSION=%%i
    echo ✓ Go found: !GO_VERSION!
) else (
    echo ✗ Go not found. Please install Go 1.20+ from https://go.dev/dl/
    exit /b 1
)

rem Check for npm or pnpm
where pnpm >nul 2>&1
if %errorlevel% equ 0 (
    echo ✓ pnpm found
    set PM=pnpm
) else (
    where npm >nul 2>&1
    if %errorlevel% equ 0 (
        echo ✓ npm found
        set PM=npm
    ) else (
        echo ✗ No package manager found. Please install pnpm or npm.
        exit /b 1
    )
)

rem Setup frontend
echo.
echo === Setting up Frontend ===
cd frontend

if exist node_modules (
    echo Dependencies already installed
) else (
    echo Installing frontend dependencies...
    if "!PM!"=="pnpm" (
        call pnpm install
    ) else (
        call npm install
    )
)

echo.
echo Frontend setup complete!
echo Run 'npm run dev' or 'pnpm dev' in the frontend directory to start the development server.
cd ..

rem Setup backend
echo.
echo === Setting up Backend ===
cd backend

if exist go.mod (
    echo Downloading Go dependencies...
    call go mod download
) else (
    echo Initializing Go module...
    call go mod init aiwriter
)

echo.
echo Backend setup complete!
echo Run 'go run main.go' in the backend directory to start the server.
cd ..

rem Print helpful information
echo.
echo === Development Server Access ===
echo Frontend: http://localhost:5173 ^(Vite default^)
echo Backend:  http://localhost:8080 ^(Gin default^)
echo.
echo === Project Structure ===
echo frontend/ - React + TypeScript + Vite application
echo backend/  - Go + Gin application
echo.
echo === Database Setup ===
echo Ensure MySQL/MySQL is running and configure database connection in backend/.env
echo.
echo === Next Steps ===
echo 1. Configure database connection in backend/.env
echo 2. Run database migrations
echo 3. Start backend: cd backend ^&^& go run main.go
echo 4. Start frontend: cd frontend ^&^& npm run dev

endlocal
