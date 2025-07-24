# 🚀 Starter Golang

> **Modern Go starter template** dengan clean architecture, siap untuk production. Menggunakan teknologi terdepan seperti Gin framework, MongoDB, dependency injection dengan Wire, dan JWT authentication.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/laudryfadian/starter-golang)

---

## ✨ Features

- 🌐 **Gin Framework** - Web framework cepat dan ringan untuk REST API
- 🍃 **MongoDB** - NoSQL database dengan performa tinggi
- 🧩 **Google Wire** - Dependency injection yang type-safe dan compile-time
- 🏗️ **Clean Architecture** - Struktur kode yang maintainable dan testable
- 🔐 **JWT Authentication** - Sistem autentikasi yang aman dan modern

---

## 📋 Requirements

- **Go** 1.21 atau lebih baru
- **MongoDB** 4.4 atau lebih baru
- **Wire** untuk dependency injection

---

## 📁 Project Structure

```
starter-golang/
├── cmd/                    # Application entrypoints
│   └── main.go            # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── domain/           # Business domain interfaces
│   ├── entity/           # Domain entities/models
│   ├── handler/          # HTTP handlers (controllers)
│   ├── helpers/          # Utility functions
│   ├── injector/         # Wire dependency injection
│   ├── middleware/       # HTTP middlewares
│   ├── repository/       # Data access layer
│   ├── routes/           # Route definitions
│   └── usecase/          # Business logic layer
├── .env.example          # Environment variables template
├── .gitignore           # Git ignore rules
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # Project documentation
```

---

## 🚀 Quick Start

### 1. Clone Repository
```bash
git clone https://github.com/laudryfadian/starter-golang.git
cd starter-golang
```

### 2. Setup Environment
```bash
# Copy environment template
cp .env.example .env

# Edit environment variables
nano .env
```

### 3. Environment Variables
```env
# Server Configuration
PORT=8080
MONGODB_URI=mongodb://localhost:27017
DB_NAME=starter_golang_db
.........
```

### 4. Install Dependencies
```bash
# Install Go dependencies
go mod tidy

# Install Wire (if not installed)
go install github.com/google/wire/cmd/wire@latest

# Install Air for hot reload (optional)
go install github.com/cosmtrek/air@latest
```

### 5. Generate Wire Dependencies
```bash
go generate ./internal/injector
```

### 6. Run Application
```bash
# Development mode with hot reload
air

# Or run directly
go run ./cmd/main.go
```

---

## 🏗️ Architecture

Proyek ini menggunakan **Clean Architecture** dengan pembagian layer sebagai berikut:

1. **Handler Layer** - Menangani HTTP requests dan responses
2. **Usecase Layer** - Berisi business logic aplikasi
3. **Repository Layer** - Mengakses data dari database
4. **Entity Layer** - Mendefinisikan struktur data domain


---

## 🛠️ Development

### Adding New Feature
1. Buat entity di `internal/entity/`
2. Definisikan interface di `internal/domain/`
3. Implementasi repository di `internal/repository/`
4. Buat usecase di `internal/usecase/`
5. Tambahkan handler di `internal/handler/`
6. Daftarkan route di `internal/routes/`
7. Update Wire injector di `internal/injector/`

### Code Style
- Gunakan `gofmt` untuk formatting
- Ikuti Go naming conventions
- Tambahkan comment untuk exported functions
- Gunakan meaningful variable names

---

## 👨‍💻 Author

**Laudry Fadian**
- GitHub: [@laudryfadian](https://github.com/laudryfadian)

---

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver)
- [Google Wire](https://github.com/google/wire)
- [JWT Go](https://github.com/golang-jwt/jwt)

---

## Make With Love ❤️
