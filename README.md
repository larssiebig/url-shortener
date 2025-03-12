# URL Shortener in Go 🚀

A simple, fast, and efficient **URL shortener** built with **Go**, **Gin**, **PostgreSQL**, and **Redis**.

## 📌 Features

✅ Shorten long URLs into short, unique links
✅ Redirect users from short links to original URLs
✅ Store URLs in **PostgreSQL** with caching in **Redis** for faster lookups
✅ Built using **Go Modules**, following **Go's official project structure**

---

## ⚙️ Setup Guide

### 1️⃣ Clone the repository

```shell
git clone https://github.com/larssiebig/url-shortener.git
cd url-shortener
```

### 2️⃣ Install dependencies

```shell
go mod tidy
```

### 3️⃣ Start PostgreSQL & Redis (using Docker)

Make sure you have **Docker** installed, then run:

```shell
docker run --name postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=shortener -p 5432:5432 -d postgres
docker run --name redis -p 6379:6379 -d redis
```

### 4️⃣ Run the server

```shell
go run cmd/server/main.go
```

## 🚀 How to Use

### 🔗 Shorten a URL

```shell
curl -X POST http://localhost:8080/shorten \
    -d '{"long_url": "https://golang.org"}' \
    -H "Content-Type: application/json"
span
```

**Response:**

```shell
curl -L http://localhost:8080/abc123
```
