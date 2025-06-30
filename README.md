# gotodo

This is a toy project I created to get hands-on experience with Go and the Gin Gonic web framework, which I will be using for my capstone project at The Marcy Lab School. The goal is to build a simple full-stack TODO application that uses Go on the backend.

## Tech Stack

- **Backend:** Go, Gin Gonic, GORM, PostgreSQL
- **Frontend:** React, TypeScript, Vite

## Getting Started

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/gotodo.git
cd gotodo
```

### 2. Set up environment variables

Copy the example environment file and fill in the required values:

```sh
cp .env.example .env
```

Edit `.env` and provide values for `POSTGRES_DSN`, `SESSION_SECRET`, etc.

### 3. Start the backend server

Make sure you have Go and PostgreSQL installed and running.

```sh
cd cmd/gotodo
go run main.go
```

The backend server will start on [http://localhost:3000](http://localhost:3000).

### 4. Start the frontend

In a new terminal window:

```sh
cd frontend
npm install
npm run dev
```

The frontend will start on [http://localhost:5173](http://localhost:5173) (default Vite port).

### 5. Access the app

- Open [http://localhost:5173](http://localhost:5173) in your browser.
- Sign up for a new account and start adding todos!
