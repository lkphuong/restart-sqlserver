# Restart SQL Server

This project is a Go application that periodically pings a SQL Server database and executes a batch script (`system.bat`) every 1 hours. It uses the `cron` library for scheduling tasks and connects to the database using the `github.com/microsoft/go-mssqldb` driver.

## Features

- Connects to a SQL Server database using credentials from a `.env` file.
- Pings the database to ensure connectivity.
- Executes a batch script (`system.bat`) every 1 hours.
- Logs the output of the batch script and database connection status.

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