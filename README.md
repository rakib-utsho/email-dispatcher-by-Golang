# Email Dispatcher - Go Application

A concurrent email campaign dispatcher built in Go using producer-consumer architecture with goroutines and channels.

## Overview

This application implements a resource-efficient, concurrent email sending system that processes email recipients from a CSV file and dispatches them through an SMTP service. It demonstrates Go's concurrency patterns using goroutines and channels.

## Architecture

### Design Pattern
**Producer-Consumer Architecture**

- **Producer**: Reads email data from `emails.csv` and sends recipients through a channel
- **Consumer**: Multiple worker goroutines that process recipients concurrently
- **Channel**: Acts as a thread-safe communication mechanism between producer and consumers

### System Flow

```
emails.csv → Producer (loadRecipients) → Channel → Workers (eamilWorker) → SMTP Service → Email Recipients
```

## Project Structure

```
email-dispatcher/
├── main.go              # Application entry point with goroutine coordination
├── producer.go          # CSV file reader that produces recipient data
├── consumer.go          # Worker function that consumes recipient data
├── emails.csv           # Email list data source
├── go.mod              # Go module configuration
├── systemDesign.drawio # System architecture diagram
└── README.md           # This file
```

## File Details

### main.go
Orchestrates the application flow:
- Defines the `Recipient` struct with `Name` and `Email` fields
- Creates an unbuffered channel for recipient communication
- Spawns a producer goroutine to load recipients from CSV
- Creates 5 worker goroutines to process recipients concurrently
- Uses `sync.WaitGroup` to synchronize goroutine completion

**Key Configuration:**
```go
workerCount := 5  // Number of concurrent workers
```

### producer.go
Handles CSV file processing:
- `loadRecipients(filePath string, ch chan Recipient) error`
- Reads all records from the CSV file
- Skips header row (`records[1:]`)
- Sends each recipient record to the channel
- Closes the channel when finished (important for worker termination)

**Error Handling:**
- Returns error if file cannot be opened
- Returns error if CSV cannot be parsed

### consumer.go
Implements worker processing logic:
- `eamilWorker(id int, ch chan Recipient, wg *sync.WaitGroup)`
- Each worker has a unique ID for identification
- Consumes recipients from the channel using range-based iteration
- Automatically signals completion via `wg.Done()`
- Workers terminate when channel is closed

### emails.csv
Sample data file containing:
- Header row: `Name,Email`
- 11 email records
- Two fields per record: recipient name and email address

**Sample Data:**
```
User 1,user1@example.com
User 2,user2@example.com
...
utsho,raibul104.abslive@gmail.com
```

## Key Features

### 🚀 Concurrent Processing
- Multiple worker goroutines process emails in parallel
- Configurable number of workers for performance tuning
- Channel-based communication ensures thread-safe data passing

### 📊 Resource Efficiency
- Unbuffered channel for balanced producer-consumer rate
- Goroutines are lightweight compared to OS threads
- Automatic cleanup with `sync.WaitGroup`

### 🔄 Synchronization
- `sync.WaitGroup` ensures all workers complete before program exits
- Channel closure signals workers to terminate gracefully
- No busy-waiting or polling required

## How It Works

1. **Initialization**: Main creates a channel and starts the producer goroutine
2. **Production**: Producer reads CSV file and sends recipients through channel
3. **Consumption**: 5 worker goroutines consume from the channel and process recipients
4. **Termination**: Producer closes channel → workers drain queue → workers exit → program ends
5. **Synchronization**: `wg.Wait()` ensures main doesn't exit until all workers finish

## Go Version

- **Go**: 1.25.1

## Module Information

```
Module: github.com/rakib-utsho/email-dispatcher
```

## Goals

✅ Use resources efficiently  
✅ Send emails concurrently  
✅ Demonstrate Go concurrency patterns  

## Current Implementation Status

- ✅ Producer goroutine reading CSV files
- ✅ Worker goroutines with graceful shutdown
- ✅ Channel-based communication
- ✅ Concurrent recipient processing
- ⚠️ SMTP integration (ready for implementation)

## Note on Function Name

There is a minor typo in the code: the function is named `eamilWorker` instead of `emailWorker`. Consider renaming for clarity:

```go
// Current: eamilWorker
// Suggested: emailWorker
```

## Future Enhancements

- Integrate actual SMTP service for email sending
- Add error handling and retry logic
- Implement rate limiting
- Add logging and metrics collection
- Support configuration file for worker count and batch size
- Add unit tests
- Implement graceful shutdown handling

## Usage

```bash
# Run the application
go run main.go producer.go consumer.go

# Build executable
go build -o email-dispatcher

# Run executable
./email-dispatcher
```

## Dependencies

- Only standard library packages:
  - `fmt` - Console output
  - `sync` - Synchronization primitives
  - `encoding/csv` - CSV parsing
  - `os` - File operations

## Author

Created by: Rakib Utsho (rakib-utsho)

---

**Last Updated**: February 2026
