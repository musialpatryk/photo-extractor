# Project Name

## Purpose
This project is a Go-based application designed to extract photos from any source with weird directory/file convention and copy them into clean date based structure.

## Prerequisites
* Docker and Docker Compose
* Make

## Getting Started

### Initialization
Start the containerized environment and initialize the Go module:
```
make up
make init
```

### Running the Application
To run the extractor inside the container:
```
make run
```

### Managing Dependencies
To clean up and sync dependencies:
```
make tidy
```

## Development & Building

### Cross-Compilation
The project supports building for multiple platforms via the container:

- Build for Linux:   make build-linux
- Build for Windows: make build-windows
- Build all:         make build-all

Binaries will be located in the /bin directory.

### Workspace Access
To open a shell inside the running container:
make shell

### Stopping the Environment
To stop and remove the containers:
make down