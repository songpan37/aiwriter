#!/bin/bash

# AI Writer - Development Environment Setup Script
# This script helps set up and run the development environment

set -e

echo "=== AI Writer Development Environment Setup ==="
echo ""

# Check Node.js version
check_node() {
    if command -v node &> /dev/null; then
        NODE_VERSION=$(node --version)
        echo "✓ Node.js found: $NODE_VERSION"
    else
        echo "✗ Node.js not found. Please install Node.js 16+ from https://nodejs.org/"
        exit 1
    fi
}

# Check Go version
check_go() {
    if command -v go &> /dev/null; then
        GO_VERSION=$(go version | grep -oP '\d+\.\d+')
        echo "✓ Go found: $(go version)"
    else
        echo "✗ Go not found. Please install Go 1.20+ from https://go.dev/dl/"
        exit 1
    fi
}

# Check for npm or pnpm
check_package_manager() {
    if command -v pnpm &> /dev/null; then
        echo "✓ pnpm found"
        PM="pnpm"
    elif command -v npm &> /dev/null; then
        echo "✓ npm found"
        PM="npm"
    else
        echo "✗ No package manager found. Please install pnpm or npm."
        exit 1
    fi
}

# Setup frontend
setup_frontend() {
    echo ""
    echo "=== Setting up Frontend ==="
    cd frontend
    
    if [ -d "node_modules" ]; then
        echo "Dependencies already installed"
    else
        echo "Installing frontend dependencies..."
        if [ "$PM" = "pnpm" ]; then
            pnpm install
        else
            npm install
        fi
    fi
    
    echo ""
    echo "Frontend setup complete!"
    echo "Run 'npm run dev' or 'pnpm dev' in the frontend directory to start the development server."
    cd ..
}

# Setup backend
setup_backend() {
    echo ""
    echo "=== Setting up Backend ==="
    cd backend
    
    if [ -f "go.mod" ]; then
        echo "Downloading Go dependencies..."
        go mod download
    else
        echo "Initializing Go module..."
        go mod init aiwriter
    fi
    
    echo ""
    echo "Backend setup complete!"
    echo "Run 'go run main.go' in the backend directory to start the server."
    cd ..
}

# Print helpful information
print_info() {
    echo ""
    echo "=== Development Server Access ==="
    echo "Frontend: http://localhost:5173 (Vite default)"
    echo "Backend:  http://localhost:8080 (Gin default)"
    echo ""
    echo "=== Project Structure ==="
    echo "frontend/ - React + TypeScript + Vite application"
    echo "backend/  - Go + Gin application"
    echo ""
    echo "=== Database Setup ==="
    echo "Ensure MySQL is running and configure database connection in backend/.env"
    echo ""
    echo "=== Next Steps ==="
    echo "1. Configure database connection in backend/.env"
    echo "2. Run database migrations"
    echo "3. Start backend: cd backend && go run main.go"
    echo "4. Start frontend: cd frontend && npm run dev"
}

# Main execution
main() {
    check_node
    check_go
    check_package_manager
    
    # Create project directories if they don't exist
    mkdir -p frontend backend
    
    # Setup components
    setup_frontend
    setup_backend
    
    print_info
}

main "$@"
