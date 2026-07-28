# Contributing to Cloud Cost Optimizer

Thank you for contributing to the **Cloud Cost Optimizer** engine!

## Development Workflow

1. Fork & clone the repository.
2. Build binary: `go build -o optimizer-server main.go`
3. Create a feature branch: `git checkout -b feat/my-optimization-module`
4. Write tests for new engine capabilities in `tests/` or package `*_test.go`.
5. Run test suite: `go test -v ./...`
6. Submit a Pull Request.
