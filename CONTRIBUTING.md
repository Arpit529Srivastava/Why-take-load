# Contributing to Simple Load Balancer Project

## Welcome Contributors!

We appreciate your interest in improving our Simple Load Balancer project. This document provides guidelines to help you contribute effectively and smoothly.

## 1. Project Overview

This project implements a load balancer in Go using a Round Robin algorithm, with multiple RESTful APIs as backend services. Before contributing, familiarize yourself with the project's structure and goals.

## 2. Getting Started

### Prerequisites
- Go 1.16 or higher
- Git
- MongoDB (for running backend APIs)

### Setup
1. Fork the repository on GitHub
2. Clone your forked repository
   ```sh
   git clone https://github.com/Arpit529Srivastava/Why-take-load.git
   cd Why-Take-Load
   ```
3. Initialize Go modules
   ```sh
   go mod tidy
   ```

## 3. Development Workflow

### Branch Strategy
- Create a new branch for each feature or bugfix
- Use a descriptive branch name
  ```sh
  git checkout -b feature/add-health-check-improvements
  git checkout -b bugfix/resolve-proxy-timeout-issue
  ```

### Commit Guidelines
- Write clear, concise commit messages
- Use present tense and imperative mood
  ```sh
  # Good commit message
  git commit -m "Add advanced health check mechanism for backend servers"
  
  # Avoid
  git commit -m "Added some changes to health checking"
  ```

## 4. Code Quality

### Coding Standards
- Follow Go standard formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments to explain complex logic
- Keep functions focused and modular

### Code Review Checklist
- [ ] Code follows Go best practices
- [ ] Functions have appropriate error handling
- [ ] New code includes relevant comments
- [ ] Tests are added/updated for changes

## 5. Testing

### Running Tests
```sh
go test ./...
```

### Test Coverage
- Add unit tests for new functionality
- Ensure existing tests pass
- Aim to maintain or improve test coverage

## 6. Performance Considerations

When contributing performance-related improvements:
- Benchmark new implementations
- Compare with existing code
- Provide performance metrics in your PR description

## 7. Documentation

- Update README.md if project setup or usage changes
- Document new features or significant modifications
- Keep inline comments clear and informative

## 8. Pull Request Process

1. Ensure your code passes all tests
2. Update documentation if needed
3. Squash commits for a clean history
4. Open a pull request with:
   - Clear title
   - Detailed description of changes
   - Reference any related issues

## 9. Communication

- Be respectful and constructive
- Ask questions if something is unclear
- Provide context for proposed changes

## 10. Reporting Issues

When reporting issues:
- Use GitHub Issues
- Provide a clear, reproducible description
- Include:
  - Steps to reproduce
  - Expected behavior
  - Actual behavior
  - Go version
  - Operating system

## Code of Conduct

We are committed to providing a friendly, safe, and welcoming environment. Harassment and discrimination are not tolerated.

## Questions?

If you have questions about contributing, please open an issue or contact the maintainers.

**Happy Contributing!** 🚀
