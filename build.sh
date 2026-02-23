#!/usr/bin/env bash
# Build script for cross-platform binaries

set -e

VERSION=${1:-$(git describe --tags --always --dirty 2>/dev/null || echo "dev")}
OUTPUT_DIR="dist"

echo "Building Kamadhenu v${VERSION} for all platforms..."

# Create dist directory
mkdir -p "$OUTPUT_DIR"

# Build matrix
platforms=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
)

for platform in "${platforms[@]}"; do
    IFS='/' read -r GOOS GOARCH <<< "$platform"
    
    output_name="kamadhenu-${GOOS}-${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output_name="${output_name}.exe"
    fi
    
    echo "Building for ${GOOS}/${GOARCH}..."
    GOOS=$GOOS GOARCH=$GOARCH go build \
        -ldflags="-s -w -X main.version=${VERSION}" \
        -o "${OUTPUT_DIR}/${output_name}"
    
    # Compress (optional)
    # gzip "${OUTPUT_DIR}/${output_name}"
done

echo "✅ Build complete! Binaries in ${OUTPUT_DIR}/"
ls -lh "$OUTPUT_DIR"
