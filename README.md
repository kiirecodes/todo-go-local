# React + Go (Gin) To-Do App (MongoDB)

## Overview
This repository contains a full-stack To-Do application built with **React** (frontend) and **Go (Gin)** (backend) using **MongoDB** as the database. Styling uses **raw CSS** with automatic dark/light theme support via `prefers-color-scheme`.

## Quick start (local)

### Backend (Go)
Requires Go 1.20+ and a running MongoDB.

```bash
cd backend
go mod download
# set environment variables or use defaults in .env.example
export MONGO_URI='mongodb://localhost:27017'
export MONGO_DB='todo_db'
go run main.go
```

API will be available at `http://localhost:8080/api`.

### Frontend (React)
```bash
cd frontend
npm install
npm start
```

Open `http://localhost:3000`.

## Deployment
- Backend: Render or any service that supports Go; set `MONGO_URI`, `MONGO_DB`, and `FRONTEND_ORIGIN` env vars.
- Frontend: Vercel or Netlify; set `REACT_APP_API_URL` environment variable (e.g. `https://your-backend.onrender.com/api`).

## Notes
- The backend uses the official MongoDB Go driver.
- This starter uses simple patterns to keep code readable for job submissions. Replace MongoDB URI with managed MongoDB for production (Atlas).
- Tests: a basic handler health test is included for quick CI checks.

Generated: 2025-10-22T08:41:30.565489 UTC
