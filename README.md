# Restart SQL Server

This project is a Go application designed to monitor and restart a SQL Server instance if it becomes unresponsive. It uses the `cron` library to schedule periodic health checks and executes Windows commands to manage SQL Server services.

## Features

- **Health Monitoring**: Periodically checks if the SQL Server service is running.
- **Automatic Restart**: Restarts the SQL Server (`MSSQLSERVER`) and SQL Server Agent (`SQLSERVERAGENT`) services if they are not running.
- **Scheduling**: Uses a cron job to perform health checks daily at 10:00 AM.
- **Logging**: Provides detailed logs for service status, restart attempts, and errors.

## Prerequisites

- **Go**: Version 1.23 or later.
- **Windows**: The application uses Windows commands to manage services.
- **SQL Server**: A running SQL Server instance.

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

3. (Optional) Create a `.env` file in the root directory if database credentials are required in the future.

## Usage

1. Build the application:
   ```sh
   go build -o service.exe main.go
   ```

2. Run the application:
   ```sh
   ./service.exe
   ```

The application will run continuously, checking the SQL Server service status and restarting it if necessary.

## How It Works

1. The application uses the `sc query` command to check the status of the `MSSQLSERVER` service.
2. If the service is not running, it executes the following commands to restart it:
   - Stops the `SQLSERVERAGENT` service.
   - Stops the `MSSQLSERVER` service.
   - Starts the `MSSQLSERVER` service.
   - Starts the `SQLSERVERAGENT` service.
3. A cron job schedules the health check to run daily at 10:00 AM.
4. Logs are generated for each action, including errors and successful restarts.

## Project Structure

```
.gitignore            # Ignored files and directories
go.mod                # Go module dependencies
go.sum                # Go module checksums
main.go               # Main application logic
README.md             # Project documentation
service.exe           # Compiled executable (ignored by Git)
```

## Dependencies

- [github.com/robfig/cron](https://github.com/robfig/cron): Library for scheduling tasks.

## Logging

The application logs important events, including:
- Service status checks.
- Restart attempts and their outcomes.
- Errors encountered during execution.

Logs are printed to the console and can be redirected to a file if needed.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Author

Developed by [lkphuong](https://github.com/lkphuong).