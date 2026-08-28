#!/bin/bash

set -e

# ReqBeam Release Script
# Builds the CLI for multiple platforms

# Change to cli directory
cd "$(dirname "$0")/../cli"

# Define supported platforms
PLATFORMS="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64"

# Create releases directory
RELEASES_DIR="../releases"
rm -rf ${RELEASES_DIR}
mkdir -p ${RELEASES_DIR}

# Build for each platform
for platform in ${PLATFORMS}; do
    OS=$(echo ${platform} | cut -d'/' -f1)
    ARCH=$(echo ${platform} | cut -d'/' -f2)
    
    echo "Building for ${OS}/${ARCH}..."
    
    # Set GOOS and GOARCH for cross-compilation
    CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build \
        -ldflags "-s -w" \
        -o ${RELEASES_DIR}/reqbeam-${OS}-${ARCH} \
        ./cmd/devspeed/main.go
    
    # Create compressed version
    if command -v gzip >/dev/null; then
        gzip -9 -c ${RELEASES_DIR}/reqbeam-${OS}-${ARCH} > ${RELEASES_DIR}/reqbeam-${OS}-${ARCH}.gz
    fi
    
    # Create zip version
    if command -v zip >/dev/null; then
        zip -9 ${RELEASES_DIR}/reqbeam-${OS}-${ARCH}.zip ${RELEASES_DIR}/reqbeam-${OS}-${ARCH}
    fi
done

# Generate checksums
if command -v sha256sum >/dev/null; then
    sha256sum ${RELEASES_DIR}/* > ${RELEASES_DIR}/checksums.txt
fi

echo "Releases created in ${RELEASES_DIR}/"
echo "Available binaries:"
ls -1 ${RELEASES_DIR}/reqbeam-*
