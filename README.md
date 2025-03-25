# go-event-processor

A proof-of-concept event processor written in Go.

This project is intended to:
- Explore Go
- Evaluate the feasibility of using Go for future event-driven services

## 🛠 Build

### 🔹 Native Build

Clone the repository and navigate to the root:

```bash
go build -o event_processor cmd/event-processor/main.go
```

Run the compiled executable:
```
./event_processor
```

### 🔹 Docker Build
A Dockerfile is provided for containerized builds.

To build the Docker image:
```
docker build -t go-event-processor .
```

To run the container:
```
docker run --rm -p 3000:3000 go-event-processor
```