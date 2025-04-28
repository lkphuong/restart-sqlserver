# Restart SQL Server

This project is a Go application that periodically pings a SQL Server database# Restart SQL Server

This project is a Go application designed to periodically monitor and restart a SQL Server instance if connectivity issues are detected. It uses the `cron` library for scheduling tasks and connects to the database using the `github.com/microsoft/go-mssqldb` driver.

## Features

- **Database Connectivity**: Connects to a SQL Server database using credentials from a `.env` file.
- **Health Check**: Pings the database every day at 10:00 AM to ensure connectivity.
- **Automatic Restart**: Restarts the SQL Server and SQL Server Agent services if the database is unreachable.

## Prerequisites

- **Go**: Version 1.23 or later.
- **SQL Server**: A running SQL Server instance.
- **Environment Variables**: A `.env` file with the following variables:
  ```env
  DATABASE_HOST=your-database-host
  DATABASE_PORT=your-database-port
  DATABASE_NAME=your-database-name
  DATABASE_USERNAME=your-database-username
  DATABASE_PASSWORD=your-database-password
  ```

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/lkphuong/restart-sqlserver.git
   cd restart-sqlserver
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Create a `.env` file in the root directory and configure your database credentials.

## Usage

1. Build the application:
   ```sh
   go build -o restart-sqlserver.exe main.go
   ```

2. Run the application:
   ```sh
   ./restart-sqlserver.exe
   ```

The application will run continuously, checking the database connection and restarting services if necessary.

## Project Structure

```
.env                  # Environment variables for database configuration
.gitignore            # Ignored files and directories
go.mod                # Go module dependencies
go.sum                # Go module checksums
main.go               # Main application logic
configs/
  hard_code.go        # Placeholder for additional configurations
  databases/
    database.go       # Database connection logic
```

## How It Works

1. The application loads database credentials from the `.env` file.
2. It establishes a connection to the SQL Server using the `github.com/microsoft/go-mssqldb` driver.
3. A cron job is scheduled to ping the database every day at 10:00 AM.
4. If the database is unreachable, the application executes Windows commands to restart the SQL Server and SQL Server Agent services.

## Dependencies

- [github.com/microsoft/go-mssqldb](https://github.com/microsoft/go-mssqldb): SQL Server driver for Go.
- [github.com/robfig/cron](https://github.com/robfig/cron): Library for scheduling tasks.
- [github.com/joho/godotenv](https://github.com/joho/godotenv): Loads environment variables from a `.env` file.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Author

Developed by [lkphuong](https://github.com/lkphuong). every 1 hours. It uses the `cron` library for scheduling tasks and connects to the database using the `github.com/microsoft/go-mssqldb` driver.

## Features

- Connects to a SQL Server database using credentials from a `.env` file.
- Pings the database to ensure connectivity.

## Prerequisites

- Go 1.23 or later
- A SQL Server instance
- A `.env` file with the following variables:
  ```env
  DATABASE_HOST=your-database-host
  DATABASE_PORT=your-database-port
  DATABASE_NAME=your-database-name
  DATABASE_USERNAME=your-database-username
  DATABASE_PASSWORD=your-database-password