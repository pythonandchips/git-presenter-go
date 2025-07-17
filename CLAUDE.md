# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building
```bash
go build -o bin/git-presenter ./cmd/git-presenter
```

### Running Tests
```bash
go test ./internal
```

### Testing Single Package
```bash
go test ./internal  # All tests in internal package
```

### Installation
```bash
go build -o $HOME/bin/git-presenter ./cmd/git-presenter
```

## Architecture Overview

Git Presenter follows clean architecture principles with clear separation of concerns:

**Entry Point**: `cmd/git-presenter/main.go`
- Handles command-line argument parsing
- Creates GitPresenter instance and delegates to it
- Supports both interactive and command modes (`-c` flag)

**Core Components** (all in `internal/` package):

1. **GitPresenter** (`internal/presenter.go`)
   - Main orchestrator that coordinates all components
   - Handles command routing and interactive mode
   - Manages the presentation lifecycle

2. **Controller** (`internal/controller.go`)
   - Manages presentation initialization and configuration
   - Interfaces with git repository using go-git library
   - Creates `.presentation` YAML file with slide metadata
   - Handles git operations for discovering commits

3. **Presentation** (`internal/presentation.go`)
   - Core presentation logic and state management
   - Handles slide navigation (next, back, start, end, list)
   - Manages current slide tracking and status display
   - Provides interactive command execution

4. **Slide** (`internal/slide.go`)
   - Represents individual slides (git commits)
   - Handles git checkout operations to move between commits
   - Stores commit hash and message metadata

## Key Dependencies

- **go-git/go-git/v5**: Git operations and repository management
- **gopkg.in/yaml.v2**: YAML parsing for presentation configuration
- **Go standard library**: Core functionality

## Presentation Workflow

1. **Initialize**: `git-presenter init` creates `.presentation` file from git history
2. **Start**: `git-presenter start` begins presentation (interactive or command mode)
3. **Navigate**: Use commands like `next`, `back`, `start`, `end`, `list`
4. **Update**: `git-presenter update` refreshes presentation configuration

## Development Methodology

The codebase follows Test-Driven Development (TDD) with comprehensive test coverage for all components. Each package has corresponding `*_test.go` files that should be maintained when making changes.

## Important Files

- `.presentation`: YAML configuration file created by `init` command
- `go.mod`: Module definition with go-git dependency
- `internal/`: All application logic following clean architecture
- `cmd/git-presenter/main.go`: Application entry point

## Testing Strategy

All tests are in the `internal/` package with a flattened structure. Tests focus on unit testing individual components. When adding new features, ensure corresponding tests are added to maintain coverage.