# Build an Orchestrator in Go (From Scratch)

[Currently this project is in development, trying to build an hobby of reverse engineering things, why not !?]

This project is inspired by Timothy Wang’s tutorial, **"Build an Orchestrator in Go (From Scratch)"**. It provides a simple yet powerful starting point for creating an orchestrator in Go, designed to manage and schedule tasks in a distributed environment.

## Features

- **Task Scheduling:** Efficiently schedule and manage tasks using Go’s concurrency primitives.
- **Concurrent Execution:** Utilize goroutines and channels to execute tasks concurrently.
- **Error Handling:** Robust error management with built-in logging for easy debugging.
- **Modular Design:** Easily extend and integrate additional functionality as needed.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Git](https://git-scm.com/downloads)
- [Optional] Docker, if you wish to run the orchestrator within a containerized environment.

## Getting Started

### Clone the Repository

Clone this repository to your local machine:

```bash
git clone <repository-url>
cd go-orchestrator