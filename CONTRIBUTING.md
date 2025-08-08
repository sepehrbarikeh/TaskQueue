# Contributing to TaskQueue Engine

Thank you for your interest in contributing to TaskQueue Engine! This document provides guidelines and information for contributors.

## 🤝 How to Contribute

### Reporting Issues

Before creating an issue, please:

1. **Search existing issues** to avoid duplicates
2. **Use the issue template** and provide all requested information
3. **Include reproduction steps** for bugs
4. **Add relevant logs** and error messages

### Feature Requests

When requesting features:

1. **Describe the use case** clearly
2. **Explain the benefits** of the feature
3. **Consider implementation complexity**
4. **Check if it aligns** with project goals

### Pull Requests

#### Before Submitting

1. **Fork the repository**
2. **Create a feature branch** from `main`
3. **Make your changes** following the guidelines below
4. **Test thoroughly** before submitting
5. **Update documentation** if needed

#### Development Guidelines

##### Code Style

- Follow **Go coding standards**
- Use **meaningful variable names**
- Add **comments** for complex logic
- Keep **functions small** and focused
- Use **consistent formatting**

##### Testing

- Write **unit tests** for new features
- Ensure **test coverage** is maintained
- Test **error conditions**
- Add **integration tests** for API changes

##### Documentation

- Update **README.md** for user-facing changes
- Add **code comments** for complex logic
- Update **API documentation** if needed
- Include **usage examples**

#### Commit Messages

Use **conventional commit messages**:

```
type(scope): description

feat(api): add job priority support
fix(worker): resolve race condition in job processing
docs(readme): update installation instructions
test(dispatcher): add unit tests for retry logic
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Test additions/changes
- `chore`: Maintenance tasks

#### Pull Request Process

1. **Update the changelog** if applicable
2. **Ensure tests pass** locally
3. **Check code coverage** is maintained
4. **Update documentation** as needed
5. **Request review** from maintainers

## 🛠️ Development Setup

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Git

### Local Development

1. **Fork and clone** the repository
2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up development environment**:
   ```bash
   docker-compose -f docker-compose.dev.yml up -d
   ```

4. **Run tests**:
   ```bash
   go test ./...
   ```

5. **Build the application**:
   ```bash
   go build -o taskqueue .
   ```

### Testing

#### Unit Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./pkg/worker -v
```

#### Integration Tests
```bash
# Start services
docker-compose -f docker-compose.dev.yml up -d

# Run API tests
./test_api.sh
```

#### Performance Tests
```bash
# Benchmark specific functions
go test -bench=. ./pkg/worker
```

### Code Quality

#### Linting
```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run linter (if using golangci-lint)
golangci-lint run
```

#### Pre-commit Hooks

Consider setting up pre-commit hooks:

```bash
# Install pre-commit
pip install pre-commit

# Install hooks
pre-commit install
```

## 📋 Issue Templates

### Bug Report

```markdown
## Bug Description
Brief description of the bug

## Steps to Reproduce
1. Step 1
2. Step 2
3. Step 3

## Expected Behavior
What should happen

## Actual Behavior
What actually happens

## Environment
- OS: [e.g., Ubuntu 20.04]
- Go Version: [e.g., 1.24.2]
- Docker Version: [e.g., 20.10.0]

## Additional Information
- Logs
- Screenshots
- Related issues
```

### Feature Request

```markdown
## Feature Description
Brief description of the feature

## Use Case
How this feature would be used

## Proposed Implementation
Suggested approach (optional)

## Alternatives Considered
Other approaches (optional)

## Additional Information
- Related issues
- References
```

## 🏷️ Labels

We use the following labels:

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Improvements to documentation
- `good first issue`: Good for newcomers
- `help wanted`: Extra attention is needed
- `priority: high`: High priority issues
- `priority: low`: Low priority issues
- `priority: medium`: Medium priority issues

## 📞 Getting Help

- **Issues**: [GitHub Issues](https://github.com/your-repo/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-repo/discussions)
- **Email**: support@taskqueue.com

## 🙏 Recognition

Contributors will be recognized in:

- **README.md** contributors section
- **Release notes**
- **Project documentation**

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.
