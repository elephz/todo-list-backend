🚀 React + Go Fiber + MySQL CRUD Project

A simple fullstack CRUD application for managing users with:

Frontend: React

Backend: Go Fiber

Database: MySQL

Supports adding users with fields:

first_name

last_name

birth_date

address

created_at, updated_at, and deleted_at

📂 Project Structure

.
├── backend/        # Go Fiber + GORM backend API
│   ├── main.go
│   └── .env
└── frontend/       # React frontend
    └── src/

⚙️ Backend Setup (Go Fiber)

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

🎨 Frontend Setup (React)

1. Prerequisites

Node.js + npm

2. Setup

cd frontend
npm install

3. Run React App

npm start

React app will run on http://localhost:5173 (or port specified by Vite or CRA).

🌐 API Endpoints

Method

Endpoint

Description

GET

/users

List all users

GET

/users/:id

Get user by ID

POST

/users

Create a user

PUT

/users/:id

Update a user

DELETE

/users/:id

Delete a user

🔐 CORS Setup (Go Fiber)

To allow frontend (e.g. http://localhost:5173) to call the backend:

import "github.com/gofiber/fiber/v2/middleware/cors"

app := fiber.New()

// Enable CORS
app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:5173",
    AllowMethods: "GET,POST,PUT,DELETE",
}))

✅ Sample Payload

When creating a user via POST /users, use JSON like:

{
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "2000-01-01T07:00:00+07:00",
  "address": "123 Main St"
}

🧠 Bonus: Calculate Age (React + Moment.js)

const age = moment().diff(moment(birthDate), 'years');

Format to Thai:

moment.locale('th');
moment(birthDate).format('DD/MM/YYYY');

📦 Tech Stack

Frontend: React, Axios, Moment.js

Backend: Go, Fiber, GORM

Database: MySQL

Middleware: CORS