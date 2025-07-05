### Testando localmente 

```go run ./cmd/main.go --url=http://localhost:8080 --requests=11 --concurrency=2```

### Build Dockerfile

```docker build -t stress-tester .```

### Docker run test 

```docker run --rm stress-tester --url=http://host.docker.internal:8080 --requests=10 --concurrency=2```