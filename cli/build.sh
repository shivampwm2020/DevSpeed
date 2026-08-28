#!/usr/bin/env bash
# Build script for ReqBeam CLI

set -e

CLI_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$CLI_DIR")"

build() {
    echo "Building ReqBeam CLI..."
    cd "$CLI_DIR"
    go build -o "$ROOT_DIR/reqbeam" ./cmd/devspeed
    echo "Build complete: $ROOT_DIR/reqbeam"
}

run() {
    build
    "$ROOT_DIR/reqbeam" "$@"
}

clean() {
    echo "Cleaning build artifacts..."
    rm -f "$ROOT_DIR/reqbeam"
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
