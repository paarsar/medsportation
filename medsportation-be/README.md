# Medsportation Backend (Go Gin + SQLite)

This is a standalone web server that handles quote requests and stores them in a persistent SQLite database.

## Architecture
- **Framework:** Gin (Golang)
- **Database:** SQLite (Persistent)
- **ORM:** GORM
- **Deployment:** Fly.io (Containerized)

## Operations via Makefile

A `Makefile` is provided to simplify common tasks.

### Local Development
- **Run the server:** `make run`
- **Build the binary:** `make build`

### Fly.io Deployment
- **Initial Setup:**
  1. `fly launch --name medsportation-be --region ewr --no-deploy`
  2. `make volume` (Creates the persistent SQLite volume)
- **Deploy Updates:** `make deploy`
- **Monitor Logs:** `make logs`
- **Check Status:** `make status`

## Production Database Management (Fly.io)

### Database Path
In production, the database is stored on a persistent volume.
- **Path:** `/data/medsportation.db`
- **Configuration:** Set via `DATABASE_PATH` secret or `fly.toml`.

### Accessing the Production Database
To inspect or manually modify the database on the live server:
1.  **SSH into the instance:**
    ```bash
    fly ssh console
    ```
2.  **Open the database with SQLite:**
    ```bash
    sqlite3 /data/medsportation.db
    ```
3.  **Example: List all admin users:**
    ```sql
    SELECT id, username, created_at FROM users;
    ```

### Adding Admin Users Manually
Since passwords must be hashed with **Bcrypt**, you cannot simply insert a plain-text password via SQL. 

**Option 1: Automated Initial Admin**
The server automatically creates `admin` with password `admin123` if the `users` table is empty.

**Option 2: Creating a new user via SQL**
If you need to add a user manually, you must first generate a Bcrypt hash. From the `fly ssh console`, you can use the application binary (if configured with a flag) or a simple SQL insert if you have a pre-computed hash.

## Database Schema
The server automatically creates and migrates the SQLite database file at the path specified by `DATABASE_PATH`.
