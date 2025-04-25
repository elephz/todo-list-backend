# 🚀 React + Go Fiber + MySQL CRUD Project

A simple fullstack CRUD application for managing users with:

Frontend: React

Backend: Go Fiber

Database: MySQL


# ⚙️ Backend Setup (Go Fiber)

1. Prerequisites

Go installed

MySQL running

2. Clone and Setup

cd backend
go mod tidy

3. Set up .env

Create a .env file in the backend/ directory:

DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=fiber_crud

4. Run MySQL and Create DB

CREATE DATABASE fiber_crud;

The app will auto-create the users table via GORM's auto-migration.

5. Run the Server

go run main.go

Your backend will be running on http://localhost:3000.

# ✅ Sample Payload

When creating a user via POST /users, use JSON like:

{
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "2000-01-01T07:00:00+07:00",
  "address": "123 Main St"
}

# 📦 Tech Stack

Frontend: React, Axios, Moment.js

Backend: Go, Fiber, GORM

Database: MySQL

Middleware: CORS