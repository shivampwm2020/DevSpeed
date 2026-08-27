#!/usr/bin/env bash
# Build script for DevSpeed CLI

set -e

CLI_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$CLI_DIR")"

build() {
    echo "Building DevSpeed CLI..."
    cd "$CLI_DIR"
    go build -o "$ROOT_DIR/devspeed" ./cmd/devspeed
    echo "Build complete: $ROOT_DIR/devspeed"
}

run() {
    build
    "$ROOT_DIR/devspeed" "$@"
}

clean() {
    echo "Cleaning build artifacts..."
    rm -f "$ROOT_DIR/devspeed"
}

# Default action
if [[ $# -eq 0 ]]; then
    run
else
    case "$1" in
        build)
            build
            ;;
        run)
            shift
            run "$@"
            ;;
        clean)
            clean
            ;;
        *)
            echo "Unknown command: $1"
            echo "Usage: $0 [build|run|clean]"
            exit 1
            ;;
    esac
fi
