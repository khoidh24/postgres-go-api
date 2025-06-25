# Postgres Go API 🌿

This is a basic CRUD API for managing notes using Go Fiber and Supabase Postgres.

## Features ✨

- Create a note
- Get all notes
- Get a note by ID
- Update a note
- Delete a note
- Delete multiple notes
- Toggle active status

## API Endpoints 🔗

- POST /api/v1/notes
- GET /api/v1/notes
- GET /api/v1/notes/:id
- PUT /api/v1/notes/:id
- DELETE /api/v1/notes/:id
- DELETE /api/v1/notes

## Authentication 🔑

- Supabase Auth

### How to setup 🧩

1. Copy the example.env file to .env
2. Fill in the values for the environment variables
3. Run the following command to install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the following command to start the server:
   ```bash
   go run cmd/main.go
   ```
