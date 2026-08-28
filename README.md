# DevSpeed

⚡ DevSpeed is a developer-machine benchmarking CLI that measures how fast your computer actually is for software development.

The experience is designed to feel like **Speedtest.net / Geekbench for developers**, not like another generic system benchmark.

## Getting Started

### Installation

```bash
# Binary download (all platforms)
curl -fsSL https://devspeed.dev/install.sh | sh

# macOS (with Homebrew)
brew install devspeed

# Windows (with winget)
winget install DevSpeed
```

### Usage

```bash
# Run the default benchmark
devspeed

# Get version information
devspeed version

# Show system information
devspeed system

# Run benchmark in verbose mode
devspeed run --verbose
```

## Features

- **Simple**: Just run `devspeed` to get started
- **No login required**: Anonymous benchmarking and sharing
- **Local-first**: All benchmarks run locally, no sensitive data leaves your machine
- **Reproducible**: Versioned benchmark suites ensure consistent results
- **Fast**: Default benchmark completes in 1-3 minutes on most machines

## Benchmarks

The current benchmark suite (v0.1.0) includes:

- **Filesystem Small Files**: Measures create, read, and delete operations with many small files
- **Git**: Performance of common Git operations like status, diff, and checkout
- **Node.js**: Dependency installation performance
- **TypeScript**: Compilation performance
- **Docker**: Build and bind-mount filesystem performance

## Development

### Project Structure

```text
devspeed/
├── cli/                  # Go CLI implementation
├── web/                  # Next.js web interface
├── api/                  # Backend API (Node.js/Fastify)
├── packages/             # Shared packages
├── docs/                 # Documentation
├── fixtures/             # Test fixtures for benchmarks
└── README.md
```

### Prerequisites

- Go 1.22+
- Node.js 18+
- PostgreSQL

### Building the CLI

```bash
cd cli
./build.sh build
```

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) first.

## License

MIT
