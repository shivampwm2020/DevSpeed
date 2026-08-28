# DevSpeed

⚡ DevSpeed is a developer-machine benchmarking CLI that measures how fast your computer actually is for software development.

The experience is designed to feel like **Speedtest.net / Geekbench for developers**, not like another generic system benchmark.

## Getting Started

### Installation

```bash
# Binary download (all platforms)
curl -fsSL https://reqbeam.dev/install.sh | sh

# Or using wget
curl -fsSL https://reqbeam.dev/install.sh -o install.sh && sh install.sh

# From source
git clone https://github.com/shivampwm2020/DevSpeed.git
cd DevSpeed
cd cli
./build.sh build
# Binary will be available at ../devspeed

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

### Building from Source

```bash
# Clone the repository
git clone https://github.com/shivampwm2020/DevSpeed.git
cd DevSpeed

# Install Go dependencies
# DevSpeed requires Go 1.22+
go mod tidy

# Build the CLI for your platform
./scripts/release.sh

# The binaries will be in the releases/ directory
ls releases/

# For local development, you can build just for your platform
cd cli
./build.sh build
# Binary will be at ../devspeed
```

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) first.

## License

MIT
