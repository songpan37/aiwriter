@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo ============================================
echo    AI Writer Development Environment
echo ============================================
echo.

rem Check Node.js version
where node >nul 2>&1
if %errorlevel% equ 0 (
    for /f "delims=" %%i in ('node --version') do set NODE_VERSION=%%i
    echo [OK] Node.js: !NODE_VERSION!
) else (
    echo [ERROR] Node.js not found. Please install Node.js 16+ from https://nodejs.org/
    exit /b 1
)

rem Check Go version
where go >nul 2>&1
if %errorlevel% equ 0 (
    for /f "delims=" %%i in ('go version') do set GO_VERSION=%%i
    echo [OK] Go: !GO_VERSION!
) else (
    echo [ERROR] Go not found. Please install Go 1.20+ from https://go.dev/dl/
    exit /b 1
)

rem Check for npm or pnpm
set PM=npm
where npm >nul 2>&1
if %errorlevel% neq 0 (
    where pnpm >nul 2>&1
    if %errorlevel% equ 0 (
        set PM=pnpm
    ) else (
        echo [ERROR] No package manager found. Please install pnpm or npm.
        exit /b 1
    )
)
if "%PM%"=="pnpm" (
    echo [OK] pnpm found
) else (
    echo [OK] npm found
)

echo.
echo ============================================
echo Installing Dependencies
echo ============================================

rem Setup frontend
echo.
echo [1/2] Setting up Frontend...
cd frontend

if exist node_modules (
    echo       Dependencies already installed
) else (
    echo       Installing dependencies...
    if "!PM!"=="pnpm" (
        call pnpm install
    ) else (
        call npm install
    )
)

cd ..

rem Setup backend
echo.
echo [2/2] Setting up Backend...
cd backend

if exist go.mod (
    echo       Checking Go dependencies...
    call go mod download
) else (
    echo       Initializing Go module...
    call go mod init aiwriter
)

cd ..

echo.
echo ============================================
echo Starting Servers
echo ============================================
echo.
echo Starting Backend server on port 8080...
start "AIWriter Backend" cmd /c "cd /d "%~dp0backend" && go run main.go"

echo Starting Frontend server on port 5173...
start "AIWriter Frontend" cmd /c "cd /d "%~dp0frontend" && npm run dev"

echo.
echo ============================================
echo Server Status
echo ============================================
echo.
echo Backend:  http://localhost:8080
echo Frontend: http://localhost:5173
echo.
echo Servers started! Close this window to stop all servers.
