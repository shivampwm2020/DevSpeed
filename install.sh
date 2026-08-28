#!/bin/bash

set -e

# DevSpeed Installer
# Downloads and installs the DevSpeed CLI

ARCH=""
OS=""
BINARY_URL=""
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="devspeed"

# Detect architecture
get_arch() {
    local kernel_arch=$(uname -m)
    case "${kernel_arch}" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            echo "Unsupported architecture: ${kernel_arch}" >&2
            exit 1
            ;;
    esac
}

# Detect OS
get_os() {
    local kernel_name=$(uname -s)
    case "${kernel_name}" in
        Linux*)
            OS="linux"
            ;;
        Darwin*)
            OS="darwin"
            ;;
        *)
            echo "Unsupported OS: ${kernel_name}" >&2
            exit 1
            ;;
    esac
}

# Determine binary URL based on OS and architecture
get_binary_url() {
    # For now, we'll use a placeholder URL
    # In a real implementation, this would point to the actual release
    BINARY_URL="https://github.com/shivampwm2020/DevSpeed/releases/latest/download/${BINARY_NAME}-${OS}-${ARCH}"
}

print_usage() {
    echo "Usage: $0 [FLAGS]"
    echo ""
    echo "Flags:"
    echo "  -h, --help     Show this help message"
    echo "  -d, --dir DIR  Install to a specific directory (default: ${INSTALL_DIR})"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            print_usage
            exit 0
            ;;
        -d|--dir)
            INSTALL_DIR="$2"
            shift 2
            ;;
        *)
            echo "Unknown argument: $1" >&2
            print_usage
            exit 1
            ;;
    esac
done

# Create temporary directory
TMP_DIR="$(mktemp -d)"
trap 'rm -rf ${TMP_DIR}' EXIT

# Ensure install directory exists
if [ ! -d "${INSTALL_DIR}" ]; then
    echo "Creating install directory: ${INSTALL_DIR}" >&2
    sudo mkdir -p "${INSTALL_DIR}"
fi

# Detect system
get_arch
get_os
get_binary_url

# Download binary
echo "Downloading DevSpeed for ${OS}/${ARCH}..." >&2
if command -v curl >/dev/null 2>&1; then
    curl -# -L -o "${TMP_DIR}/${BINARY_NAME}" "${BINARY_URL}"
elif command -v wget >/dev/null 2>&1; then
    wget -q --show-progress -O "${TMP_DIR}/${BINARY_NAME}" "${BINARY_URL}"
else
    echo "Neither curl nor wget found" >&2
    exit 1
fi

# Make binary executable
chmod +x "${TMP_DIR}/${BINARY_NAME}"

# Check if we need sudo for installation
if [ -w "${INSTALL_DIR}" ]; then
    INSTALL_CMD="cp"
else
    INSTALL_CMD="sudo cp"
fi

# Install binary
${INSTALL_CMD} "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"

# Verify installation
if command -v ${BINARY_NAME} >/dev/null 2>&1; then
    echo "DevSpeed installed successfully to ${INSTALL_DIR}/${BINARY_NAME}" >&2
    ${BINARY_NAME} version
else
    echo "Installation failed" >&2
    exit 1
fi
